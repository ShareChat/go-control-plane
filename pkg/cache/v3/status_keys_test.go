package cache_test

import (
	"testing"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	cache "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// suffixHash derives a key from more than node.Id, the way a per-pod or
// per-shard scheme does. IDHash returns node.Id verbatim, which is what hid
// this: with IDHash the cache key and the raw node ID are the same string.
type suffixHash struct{}

func (suffixHash) ID(node *core.Node) string {
	if node == nil {
		return ""
	}
	return node.GetId() + "~pod-1"
}
func (h suffixHash) CacheIndex(node *core.Node) int { return h.CacheIndexFromKey(h.ID(node)) }
func (suffixHash) CacheIndexFromKey(key string) int { return len(key) % 8 }

// GetStatusKeys must return keys that work with GetSnapshot. The status map is
// keyed by hash.ID(node) (getOrCreateStatus), so returning the raw node ID off
// the stored proto instead makes every lookup built from this list miss as soon
// as the hash derives a key from anything but node.Id - silently, because a
// missing snapshot is not an error on the write paths.
func TestGetStatusKeysAreUsableAsCacheKeys(t *testing.T) {
	c := cache.NewSnapshotCache(false, suffixHash{}, nil)

	node := &core.Node{Id: "node-a"}
	responses := make(chan cache.DeltaResponse, 1)
	cancel, _ := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
		Node:    node,
		TypeUrl: testTypes[0],
	}, stream.NewStreamState(true, nil), responses)
	defer cancel()

	keys := c.GetStatusKeys()
	if len(keys) != 1 {
		t.Fatalf("GetStatusKeys returned %d keys, want 1: %v", len(keys), keys)
	}

	want := suffixHash{}.ID(node)
	if keys[0] != want {
		t.Fatalf("GetStatusKeys returned %q, want the cache key %q. The status map is keyed by hash.ID(node); returning the raw node ID makes every GetSnapshot built from this list miss.", keys[0], want)
	}

	// The contract that matters: the key round-trips through the cache.
	if info := c.GetStatusInfo(keys[0]); info == nil {
		t.Fatalf("GetStatusInfo(%q) found nothing - the key does not address the status map", keys[0])
	}
}
