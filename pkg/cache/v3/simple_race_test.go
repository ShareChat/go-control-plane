// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package cache_test

import (
	"context"
	"sync"
	"testing"
	"time"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// Every proxy in a zone connects with the SAME node ID, so all of their streams
// land on one status entry. This is the shape that produced the production
// incident these tests exist for.
const sharedNodeID = "shared-node"

// TestConcurrentCreateDeltaWatchNoOrphans is the regression test for orphaned
// delta watches.
//
// snapshotCache.CreateDeltaWatch used to read, then create, then write a status
// slot with no synchronisation (the cache-wide mutex was commented out). When
// several streams for one node ID opened their first watch at the same moment —
// exactly what a control-plane restart causes, since every proxy reconnects at
// once — each goroutine saw a nil statusInfo and installed its own. Only the
// last write survived in the cache; watches registered into the others were
// unreachable from respondDeltaWatches for the life of the process.
//
// The loss was permanent, not transient. An orphaned stream is blocked waiting
// on a response that can no longer be sent to it, so it never ACKs, so the
// server never calls CreateDeltaWatch again to re-register it. Only a brand-new
// stream recovered.
//
// Measured in production 2026-07-31: 45 of 52 proxies stuck without two
// clusters that were present in the control plane's own snapshot, indefinitely.
func TestConcurrentCreateDeltaWatchNoOrphans(t *testing.T) {
	const streams = 50

	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	node := &core.Node{Id: sharedNodeID}

	// No snapshot exists yet — a freshly started control plane — so every
	// CreateDeltaWatch takes the delayed-response path and registers a watch
	// rather than replying inline.
	chans := make([]chan cache.DeltaResponse, streams)
	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < streams; i++ {
		chans[i] = make(chan cache.DeltaResponse, 1)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-start // maximise overlap on the nil check
			c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
				Node:    node,
				TypeUrl: rsrc.ClusterType,
			}, stream.NewStreamState(true, nil), chans[i])
		}(i)
	}
	close(start)
	wg.Wait()

	// Every one of those registrations must be reachable from the single
	// statusInfo the cache kept.
	info := c.GetStatusInfo(sharedNodeID)
	if info == nil {
		t.Fatal("no status info registered for the shared node ID")
	}
	if got := info.GetNumDeltaWatches(); got != streams {
		t.Errorf("status holds %d delta watches, want %d: %d registrations were "+
			"orphaned in statusInfo objects the cache overwrote and can no longer reach",
			got, streams, streams-got)
	}

	// And a subsequent update must actually reach all of them.
	err := c.UpsertResources(context.Background(), sharedNodeID, rsrc.ClusterType,
		map[string]*types.ResourceWithTTL{
			clusterName: {Resource: testCluster, Version: "v1"},
		})
	if err != nil {
		t.Fatalf("UpsertResources: %v", err)
	}

	time.Sleep(500 * time.Millisecond)
	delivered := 0
	for i := 0; i < streams; i++ {
		select {
		case <-chans[i]:
			delivered++
		default:
		}
	}
	if delivered != streams {
		t.Errorf("%d/%d streams received the update; %d are permanently starved",
			delivered, streams, streams-delivered)
	}
}

// TestConcurrentCreateWatchNoOrphans is the SOTW counterpart. CreateWatch had
// the identical unsynchronised read-then-write of the status slot.
func TestConcurrentCreateWatchNoOrphans(t *testing.T) {
	const streams = 50

	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	node := &core.Node{Id: sharedNodeID}

	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < streams; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			c.CreateWatch(&discovery.DiscoveryRequest{
				Node:    node,
				TypeUrl: rsrc.ClusterType,
			}, stream.NewStreamState(true, nil), make(chan cache.Response, 1))
		}()
	}
	close(start)
	wg.Wait()

	info := c.GetStatusInfo(sharedNodeID)
	if info == nil {
		t.Fatal("no status info registered for the shared node ID")
	}
	if got := info.GetNumWatches(); got != streams {
		t.Errorf("status holds %d watches, want %d: %d registrations were orphaned",
			got, streams, streams-got)
	}
}

// TestDistinctNodesDoNotShareState guards the other half of the old layout.
// Snapshots and status used to live in flat 16384-entry slices indexed by
// hash.CacheIndexFromKey(nodeID), with no record of which node owned a slot, so
// two node IDs whose hashes collided silently shared one snapshot — node A
// being served node B's config. collidingHash forces that collision.
func TestDistinctNodesDoNotShareState(t *testing.T) {
	c := cache.NewSnapshotCache(false, collidingHash{}, logger{t: t})

	if err := c.SetSnapshot(context.Background(), "node-a", fixture.snapshot()); err != nil {
		t.Fatal(err)
	}

	if _, err := c.GetSnapshot("node-b"); err == nil {
		t.Error("node-b was served node-a's snapshot: distinct node IDs are sharing a slot")
	}

	snap, err := c.GetSnapshot("node-a")
	if err != nil {
		t.Fatalf("node-a lost its own snapshot: %v", err)
	}
	if snap == nil {
		t.Error("node-a snapshot is nil")
	}
}

// collidingHash maps every node ID to the same index, the degenerate case of
// the shipped IDHash default (which returns len(key), so any two node IDs of
// equal length collide).
type collidingHash struct{}

func (collidingHash) ID(node *core.Node) string {
	if node != nil {
		return node.GetId()
	}
	return key
}
func (collidingHash) CacheIndex(*core.Node) int    { return 0 }
func (collidingHash) CacheIndexFromKey(string) int { return 0 }
