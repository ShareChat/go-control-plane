package cache_test

import (
	"context"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// hookLogger runs onRequested the first time CreateWatch logs its "requested ... known ... Diff"
// debug line. That line is emitted after CreateWatch has read the snapshot and before it
// registers the watch, so it is a deterministic point inside that window.
type hookLogger struct {
	armed       int32
	fired       int32
	onRequested func()
}

func (l *hookLogger) Debugf(format string, _ ...interface{}) {
	if atomic.LoadInt32(&l.armed) == 1 && strings.HasPrefix(format, "nodeID %q requested") && atomic.CompareAndSwapInt32(&l.fired, 0, 1) {
		l.onRequested()
	}
}
func (l *hookLogger) Infof(string, ...interface{})  {}
func (l *hookLogger) Warnf(string, ...interface{})  {}
func (l *hookLogger) Errorf(string, ...interface{}) {}

func upsertClusters(t *testing.T, c cache.SnapshotCache, node string, names ...string) string {
	t.Helper()
	in := map[string]*types.ResourceWithTTL{}
	for _, n := range names {
		in[n] = &types.ResourceWithTTL{Resource: &clusterv3.Cluster{Name: n}, Version: "1"}
	}
	require.NoError(t, c.UpsertResources(context.Background(), node, rsrc.ClusterType, in))
	snap, err := c.GetSnapshot(node)
	require.NoError(t, err)
	return snap.GetVersion(rsrc.ClusterType)
}

func parkSOTWWatch(t *testing.T, c cache.SnapshotCache, node, version string) (chan cache.Response, func()) {
	t.Helper()
	req := &cache.Request{TypeUrl: rsrc.ClusterType, Node: &core.Node{Id: node}, VersionInfo: version}
	ch := make(chan cache.Response, 1)
	cancel := c.CreateWatch(req, stream.NewStreamState(true, nil), ch)
	require.NotNil(t, cancel)
	require.Empty(t, ch, "up-to-date SOTW request must park")
	return ch, cancel
}

func clusterNames(t *testing.T, r cache.Response) []string {
	t.Helper()
	raw, ok := r.(*cache.RawResponse)
	require.True(t, ok)
	names := make([]string, 0, len(raw.Resources))
	for _, res := range raw.Resources {
		names = append(names, res.Name)
	}
	sort.Strings(names)
	return names
}

// A DeleteResources that runs between CreateWatch deciding "client is up to date" and registering
// its SOTW watch must still reach that client: otherwise it keeps the deleted cluster.
func TestCreateWatch_SOTWConcurrentDeleteNotMissed(t *testing.T) {
	ctx := context.Background()
	const node = "n1"
	logger := &hookLogger{}
	c := cache.NewSnapshotCache(true, group{}, logger)
	oldVersion := upsertClusters(t, c, node, "a", "b")

	deleteDone := make(chan struct{})
	logger.onRequested = func() {
		go func() {
			defer close(deleteDone)
			_ = c.DeleteResources(ctx, node, rsrc.ClusterType, []string{"b"})
		}()
		select {
		case <-deleteDone:
		case <-time.After(300 * time.Millisecond):
		}
	}
	atomic.StoreInt32(&logger.armed, 1)

	req := &cache.Request{TypeUrl: rsrc.ClusterType, Node: &core.Node{Id: node}, VersionInfo: oldVersion}
	ch := make(chan cache.Response, 1)
	cancel := c.CreateWatch(req, stream.NewStreamState(true, nil), ch)
	require.NotNil(t, cancel)
	defer cancel()
	require.Equal(t, int32(1), atomic.LoadInt32(&logger.fired), "hook never fired")
	<-deleteDone

	select {
	case r := <-ch:
		assert.Equal(t, []string{"a"}, clusterNames(t, r))
	default:
		t.Fatalf("SOTW watch parked at version %q was never notified of the delete: the client keeps deleted cluster b", oldVersion)
	}
}

// An ordered-ADS SOTW stream closes its shared response channel on shutdown before its per-type
// watches are cancelled (watches.close() runs in map order). An update that walks those watches in
// between must drop them, not panic sending on the closed channel.
func TestDeleteResources_SOTWWatchOnShutdownADSStreamDoesNotPanic(t *testing.T) {
	ctx := context.Background()
	const node = "n1"
	c := cache.NewSnapshotCache(true, group{}, nil)
	version := upsertClusters(t, c, node, "a", "b")
	ch, cancel := parkSOTWWatch(t, c, node, version)
	defer cancel()
	close(ch)

	require.NotPanics(t, func() {
		require.NoError(t, c.DeleteResources(ctx, node, rsrc.ClusterType, []string{"b"}))
		// The dead watch was dropped, so the next update does not reach it either.
		require.NoError(t, c.DeleteResources(ctx, node, rsrc.ClusterType, []string{"a"}))
	})
}

// UpsertResources and BatchUpsertResources mutate the snapshot in place and bump the type version,
// so they must wake parked SOTW watches as DeleteResources and SetSnapshot do.
func TestUpsertResources_NotifiesParkedSOTWWatch(t *testing.T) {
	ctx := context.Background()
	const node = "n1"
	c := cache.NewSnapshotCache(true, group{}, nil)
	version := upsertClusters(t, c, node, "a")

	ch, cancel := parkSOTWWatch(t, c, node, version)
	defer cancel()
	version2 := upsertClusters(t, c, node, "b")
	select {
	case r := <-ch:
		assert.Equal(t, []string{"a", "b"}, clusterNames(t, r))
	default:
		t.Fatal("parked SOTW watch was not notified of the cluster upsert")
	}

	ch, cancel2 := parkSOTWWatch(t, c, node, version2)
	defer cancel2()
	require.NoError(t, c.BatchUpsertResources(ctx, rsrc.ClusterType, map[string]map[string]*types.ResourceWithTTL{
		node: {"c": {Resource: &clusterv3.Cluster{Name: "c"}, Version: "1"}},
	}))
	select {
	case r := <-ch:
		assert.Equal(t, []string{"a", "b", "c"}, clusterNames(t, r))
	default:
		t.Fatal("parked SOTW watch was not notified of the batch upsert")
	}
}
