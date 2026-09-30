package cache_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// DeleteResources was a no-op (body commented out in 3ca8faf7). A delete must drop the
// resource and tell a parked wildcard CDS watch via removed_resources.
func TestDeleteResources_RemovesAndNotifiesDeltaWatch(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSnapshotCache(true, group{}, nil)
	const node = "n1"
	require.NoError(t, c.UpsertResources(ctx, node, rsrc.ClusterType, map[string]*types.ResourceWithTTL{
		"a": {Resource: &clusterv3.Cluster{Name: "a"}, Version: "1"},
		"b": {Resource: &clusterv3.Cluster{Name: "b"}, Version: "1"},
	}))

	req := &cache.DeltaRequest{TypeUrl: rsrc.ClusterType, Node: &core.Node{Id: node}}
	// First watch answers immediately with a+b; adopt its versions as the client's state.
	first := make(chan cache.DeltaResponse, 1)
	_, delayed := c.CreateDeltaWatch(req, stream.NewStreamState(true, nil), first)
	require.False(t, delayed)
	state := stream.NewStreamState(true, nil)
	state.SetResourceVersions((<-first).GetNextVersionMap()) // also clears "first"

	// Second watch parks: the client is up to date.
	second := make(chan cache.DeltaResponse, 1)
	_, delayed = c.CreateDeltaWatch(req, state, second)
	require.True(t, delayed)

	require.NoError(t, c.DeleteResources(ctx, node, rsrc.ClusterType, []string{"b"}))

	select {
	case r := <-second:
		dr, err := r.GetDeltaDiscoveryResponse()
		require.NoError(t, err)
		assert.Equal(t, []string{"b"}, dr.GetRemovedResources())
	default:
		t.Fatal("parked watch was not told about the removal")
	}
	snap, err := c.GetSnapshot(node)
	require.NoError(t, err)
	_, stillThere := snap.GetResourcesAndTTL(rsrc.ClusterType)["b"]
	assert.False(t, stillThere)
}

// A named (EDS-style) subscription to a CLA that gets deleted must be told via
// removed_resources.
func TestDeleteResources_NamedSubscriptionNotified(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSnapshotCache(true, group{}, nil)
	const node = "n1"
	require.NoError(t, c.UpsertResources(ctx, node, rsrc.EndpointType, map[string]*types.ResourceWithTTL{
		"a": {Resource: &endpointv3.ClusterLoadAssignment{ClusterName: "a"}, Version: "1"},
		"b": {Resource: &endpointv3.ClusterLoadAssignment{ClusterName: "b"}, Version: "1"},
	}))

	req := &cache.DeltaRequest{TypeUrl: rsrc.EndpointType, Node: &core.Node{Id: node}}
	sub := map[string]struct{}{"b": {}}

	first := stream.NewStreamState(false, nil)
	first.SetSubscribedResourceNames(sub)
	ch1 := make(chan cache.DeltaResponse, 1)
	_, delayed := c.CreateDeltaWatch(req, first, ch1)
	require.False(t, delayed)

	state := stream.NewStreamState(false, nil)
	state.SetSubscribedResourceNames(sub)
	state.SetResourceVersions((<-ch1).GetNextVersionMap())
	ch2 := make(chan cache.DeltaResponse, 1)
	_, delayed = c.CreateDeltaWatch(req, state, ch2)
	require.True(t, delayed)

	require.NoError(t, c.DeleteResources(ctx, node, rsrc.EndpointType, []string{"b"}))

	select {
	case r := <-ch2:
		dr, err := r.GetDeltaDiscoveryResponse()
		require.NoError(t, err)
		assert.Equal(t, []string{"b"}, dr.GetRemovedResources())
	default:
		t.Fatal("named watch was not told about the removal")
	}
}

// Unknown names and nodes without a snapshot are silent no-ops: no error, no version
// bump, no watch response.
func TestDeleteResources_NoopCases(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSnapshotCache(true, group{}, nil)
	require.NoError(t, c.DeleteResources(ctx, "missing-node", rsrc.ClusterType, []string{"x"}))

	require.NoError(t, c.UpsertResources(ctx, "n1", rsrc.ClusterType, map[string]*types.ResourceWithTTL{
		"a": {Resource: &clusterv3.Cluster{Name: "a"}, Version: "1"},
	}))
	before, err := c.GetSnapshot("n1")
	require.NoError(t, err)
	v := before.GetVersion(rsrc.ClusterType)
	require.NoError(t, c.DeleteResources(ctx, "n1", rsrc.ClusterType, []string{"not-there"}))
	after, err := c.GetSnapshot("n1")
	require.NoError(t, err)
	assert.Equal(t, v, after.GetVersion(rsrc.ClusterType))
}
