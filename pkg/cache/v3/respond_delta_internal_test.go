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
