package cache

import (
	"context"
	"errors"
	"testing"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// Sending on a closed channel panics. The recover used to leave the return
// values at their zero values, so a closed channel reported (nil, nil) - which
// respondDeltaWatches reads as "no state change, keep the watch" and
// CreateDeltaWatch reads as "delayedResponse, register the watch". Both then
// hold a watch against a stream nobody will ever read from, and it looks
// exactly like a healthy idle watch.
func TestRespondDeltaReportsAClosedChannel(t *testing.T) {
	snapshot, err := NewSnapshotWithTTLs("v1", map[resource.Type][]types.ResourceWithTTL{
		resource.ClusterType: {{Resource: &cluster.Cluster{Name: "c1"}, Version: "v1"}},
	})
	if err != nil {
		t.Fatalf("NewSnapshot: %v", err)
	}
	if err := snapshot.ConstructVersionMap(); err != nil {
		t.Fatalf("ConstructVersionMap: %v", err)
	}

	c := newSnapshotCache(false, IDHash{}, nil)

	closed := make(chan DeltaResponse, 1)
	close(closed)

	req := &DeltaRequest{Node: &core.Node{Id: "node-1"}, TypeUrl: resource.ClusterType}
	state := stream.NewStreamState(true, nil)

	out, err := c.respondDelta(context.Background(), snapshot, req, closed, state)

	if !errors.Is(err, ErrResponseChannelClosed) {
		t.Errorf("err = %v, want ErrResponseChannelClosed", err)
	}
	if out != nil {
		t.Errorf("response = %v, want nil", out)
	}
}

// The other nil-response path must stay distinguishable: a full channel is
// retryable and deliberately reports (nil, nil) so the caller keeps the watch.
func TestRespondDeltaFullChannelStillReportsNoError(t *testing.T) {
	snapshot, err := NewSnapshotWithTTLs("v1", map[resource.Type][]types.ResourceWithTTL{
		resource.ClusterType: {{Resource: &cluster.Cluster{Name: "c1"}, Version: "v1"}},
	})
	if err != nil {
		t.Fatalf("NewSnapshot: %v", err)
	}
	if err := snapshot.ConstructVersionMap(); err != nil {
		t.Fatalf("ConstructVersionMap: %v", err)
	}

	c := newSnapshotCache(false, IDHash{}, nil)

	full := make(chan DeltaResponse) // unbuffered, nobody receiving
	req := &DeltaRequest{Node: &core.Node{Id: "node-1"}, TypeUrl: resource.ClusterType}
	state := stream.NewStreamState(true, nil)

	out, err := c.respondDelta(context.Background(), snapshot, req, full, state)

	if err != nil {
		t.Errorf("err = %v, want nil for a full channel", err)
	}
	if out != nil {
		t.Errorf("response = %v, want nil for a full channel", out)
	}
}

// Non-ADS: a watch whose stream is gone must not stop delivery to the node's other
// watches. respondDeltaWatches drops the dead watch and keeps going; before, it
// returned on ErrResponseChannelClosed, so whichever watches map iteration had not
// reached yet missed this update.
func TestRespondDeltaWatchesNonADSSkipsAClosedChannel(t *testing.T) {
	snapshot, err := NewSnapshotWithTTLs("v1", map[resource.Type][]types.ResourceWithTTL{
		resource.ClusterType: {{Resource: &cluster.Cluster{Name: "c1"}, Version: "v1"}},
	})
	if err != nil {
		t.Fatalf("NewSnapshot: %v", err)
	}

	c := newSnapshotCache(false, IDHash{}, nil)
	node := &core.Node{Id: "node-1"}
	info := newStatusInfo(node)

	closed := make(chan DeltaResponse, 1)
	close(closed)
	// Several live watches, so map order cannot put all of them before the dead one.
	live := map[int64]chan DeltaResponse{}
	info.deltaWatches[0] = DeltaResponseWatch{
		Request:     &DeltaRequest{Node: node, TypeUrl: resource.ClusterType},
		Response:    closed,
		StreamState: stream.NewStreamState(true, nil),
	}
	for id := int64(1); id <= 8; id++ {
		live[id] = make(chan DeltaResponse, 1)
		info.deltaWatches[id] = DeltaResponseWatch{
			Request:     &DeltaRequest{Node: node, TypeUrl: resource.ClusterType},
			Response:    live[id],
			StreamState: stream.NewStreamState(true, nil),
		}
	}

	if err := c.respondDeltaWatches(context.Background(), info, snapshot); err != nil {
		t.Fatalf("respondDeltaWatches: %v, want nil", err)
	}
	if len(info.deltaWatches) != 0 {
		t.Errorf("%d watches left, want 0 (dead one dropped, live ones answered)", len(info.deltaWatches))
	}
	for id, ch := range live {
		if len(ch) != 1 {
			t.Errorf("live watch %d got no response", id)
		}
	}
}
