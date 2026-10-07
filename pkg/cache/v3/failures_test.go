package cache_test

import (
	"errors"
	"sync"
	"testing"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

var errMarshal = errors.New("synthetic marshal failure")

// unmarshalableResource is a real proto.Message whose MarshalVTStrict always
// fails, which is the case the cache used to swallow with a fmt.Printf.
type unmarshalableResource struct {
	*cluster.Cluster
}

func (unmarshalableResource) MarshalVTStrict() ([]byte, error) { return nil, errMarshal }

// A resource that cannot be marshalled is dropped from the snapshot: it is
// never pushed, and a proxy already holding it sees it removed. Before this
// change the only trace was an unstructured line on stdout, which is
// unreadable in an environment logging 55k lines/min against a one-minute
// kubelet buffer. Nothing could count it.
func TestDroppedResourceIsReported(t *testing.T) {
	var (
		mu     sync.Mutex
		calls  int
		gotErr error
	)
	original := cache.OnResourceMarshalError
	t.Cleanup(func() { cache.OnResourceMarshalError = original })
	cache.OnResourceMarshalError = func(_, _ string, err error) {
		mu.Lock()
		defer mu.Unlock()
		calls++
		gotErr = err
	}

	items := []types.ResourceWithTTL{
		{Resource: unmarshalableResource{Cluster: &cluster.Cluster{Name: "bad"}}},
		{Resource: &cluster.Cluster{Name: "good"}},
	}
	indexed := cache.IndexAndMarshalResourcesByName(items)

	mu.Lock()
	defer mu.Unlock()
	if calls != 1 {
		t.Errorf("OnResourceMarshalError called %d times, want 1", calls)
	}
	if !errors.Is(gotErr, errMarshal) {
		t.Errorf("hook got error %v, want %v", gotErr, errMarshal)
	}
	if _, ok := indexed["bad"]; ok {
		t.Error("unmarshalable resource made it into the snapshot")
	}
	if _, ok := indexed["good"]; !ok {
		t.Error("the good resource was dropped alongside the bad one")
	}
}
