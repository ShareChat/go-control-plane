package cache_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
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

// Deleting the LAST resource of a type while no watch is parked: a client that then
// re-subscribes with its old versions must be answered immediately with the removal.
func TestDeleteResources_LastResourceRemovalOnResubscribe(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSnapshotCache(true, group{}, nil)
	const node = "n1"
	require.NoError(t, c.UpsertResources(ctx, node, rsrc.ClusterType, map[string]*types.ResourceWithTTL{
		"b": {Resource: &clusterv3.Cluster{Name: "b"}, Version: "1"},
	}))
	req := &cache.DeltaRequest{TypeUrl: rsrc.ClusterType, Node: &core.Node{Id: node}}
	first := make(chan cache.DeltaResponse, 1)
	_, delayed := c.CreateDeltaWatch(req, stream.NewStreamState(true, nil), first)
	require.False(t, delayed)
	state := stream.NewStreamState(true, nil)
	state.SetResourceVersions((<-first).GetNextVersionMap())

	require.NoError(t, c.DeleteResources(ctx, node, rsrc.ClusterType, []string{"b"}))

	ch := make(chan cache.DeltaResponse, 1)
	_, delayed = c.CreateDeltaWatch(req, state, ch)
	require.False(t, delayed, "removal of the last resource must answer immediately")
	dr, err := (<-ch).GetDeltaDiscoveryResponse()
	require.NoError(t, err)
	assert.Equal(t, []string{"b"}, dr.GetRemovedResources())

	// A fresh stream on the now-empty type must still park, not get an empty response.
	fresh := make(chan cache.DeltaResponse, 1)
	_, delayed = c.CreateDeltaWatch(req, stream.NewStreamState(true, nil), fresh)
	assert.True(t, delayed)
}

// Unknown names, unknown type URLs and nodes without a snapshot are silent no-ops: no
// error, no version bump, no watch response.
func TestDeleteResources_NoopCases(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSnapshotCache(true, group{}, nil)
	require.NoError(t, c.DeleteResources(ctx, "missing-node", rsrc.ClusterType, []string{"x"}))

	require.NoError(t, c.UpsertResources(ctx, "n1", rsrc.ClusterType, map[string]*types.ResourceWithTTL{
		"a": {Resource: &clusterv3.Cluster{Name: "a"}, Version: "1"},
	}))
	req := &cache.DeltaRequest{TypeUrl: rsrc.ClusterType, Node: &core.Node{Id: "n1"}}
	first := make(chan cache.DeltaResponse, 1)
	_, delayed := c.CreateDeltaWatch(req, stream.NewStreamState(true, nil), first)
	require.False(t, delayed)
	state := stream.NewStreamState(true, nil)
	state.SetResourceVersions((<-first).GetNextVersionMap())
	parked := make(chan cache.DeltaResponse, 1)
	_, delayed = c.CreateDeltaWatch(req, state, parked)
	require.True(t, delayed)

	before, err := c.GetSnapshot("n1")
	require.NoError(t, err)
	v := before.GetVersion(rsrc.ClusterType)
	require.NoError(t, c.DeleteResources(ctx, "n1", rsrc.ClusterType, []string{"not-there"}))
	require.NoError(t, c.DeleteResources(ctx, "n1", "type.googleapis.com/unknown.Type", []string{"a"}))
	after, err := c.GetSnapshot("n1")
	require.NoError(t, err)
	assert.Equal(t, v, after.GetVersion(rsrc.ClusterType))
	assert.Empty(t, parked, "no-op delete must not answer the parked watch")
}

// A restarted control plane builds types one at a time. A client reconnecting with
// versions for a type the snapshot has never populated must be parked, not told to
// remove what it holds, even when another type is upserted meanwhile.
func TestCreateDeltaWatch_NeverPopulatedTypeParksOnResubscribe(t *testing.T) {
	testNeverPopulatedTypeParksOnResubscribe(t, true)
}

func TestCreateDeltaWatch_NeverPopulatedTypeParksOnResubscribeNonADS(t *testing.T) {
	testNeverPopulatedTypeParksOnResubscribe(t, false)
}

func testNeverPopulatedTypeParksOnResubscribe(t *testing.T, ads bool) {
	ctx := context.Background()
	c := cache.NewSnapshotCache(ads, group{}, nil)
	const node = "n1"
	require.NoError(t, c.UpsertResources(ctx, node, rsrc.ClusterType, map[string]*types.ResourceWithTTL{
		"c1": {Resource: &clusterv3.Cluster{Name: "c1"}, Version: "1"},
	}))

	req := &cache.DeltaRequest{TypeUrl: rsrc.ListenerType, Node: &core.Node{Id: node}, InitialResourceVersions: map[string]string{"L": "v1"}}
	ch := make(chan cache.DeltaResponse, 1)
	_, delayed := c.CreateDeltaWatch(req, stream.NewStreamState(true, map[string]string{"L": "v1"}), ch)
	require.True(t, delayed, "never-populated type must park")
	require.Empty(t, ch)

	require.NoError(t, c.UpsertResources(ctx, node, rsrc.ClusterType, map[string]*types.ResourceWithTTL{
		"c2": {Resource: &clusterv3.Cluster{Name: "c2"}, Version: "1"},
	}))
	require.Empty(t, ch, "upsert of another type must not answer with a removal")

	require.NoError(t, c.UpsertResources(ctx, node, rsrc.ListenerType, map[string]*types.ResourceWithTTL{
		"L": {Resource: &listenerv3.Listener{Name: "L"}, Version: "v2"},
	}))
	select {
	case r := <-ch:
		dr, err := r.GetDeltaDiscoveryResponse()
		require.NoError(t, err)
		assert.Empty(t, dr.GetRemovedResources())
		require.Len(t, dr.GetResources(), 1)
		assert.Equal(t, "L", dr.GetResources()[0].GetName())
	default:
		t.Fatal("parked watch was not answered once the type was set")
	}
}

// A type that was populated and emptied by deleting its last resource is a real
// delete: a reconnecting client holding it must be told to remove it.
func TestCreateDeltaWatch_EmptiedTypeRemovesOnResubscribe(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSnapshotCache(true, group{}, nil)
	const node = "n1"
	require.NoError(t, c.UpsertResources(ctx, node, rsrc.ListenerType, map[string]*types.ResourceWithTTL{
		"L": {Resource: &listenerv3.Listener{Name: "L"}, Version: "v1"},
	}))
	require.NoError(t, c.DeleteResources(ctx, node, rsrc.ListenerType, []string{"L"}))

	req := &cache.DeltaRequest{TypeUrl: rsrc.ListenerType, Node: &core.Node{Id: node}, InitialResourceVersions: map[string]string{"L": "v1"}}
	ch := make(chan cache.DeltaResponse, 1)
	_, delayed := c.CreateDeltaWatch(req, stream.NewStreamState(true, map[string]string{"L": "v1"}), ch)
	require.False(t, delayed)
	dr, err := (<-ch).GetDeltaDiscoveryResponse()
	require.NoError(t, err)
	assert.Equal(t, []string{"L"}, dr.GetRemovedResources())
}
