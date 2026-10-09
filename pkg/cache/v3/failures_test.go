package cache_test

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"sync"
	"testing"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
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

// The snapshot constructors know the type a resource would be pushed under, so
// the hook must get it rather than "".
func TestDroppedSnapshotResourceReportsTypeURL(t *testing.T) {
	original := cache.OnResourceMarshalError
	t.Cleanup(func() { cache.OnResourceMarshalError = original })
	var gotType string
	cache.OnResourceMarshalError = func(typeURL, _ string, _ error) { gotType = typeURL }

	bad := unmarshalableResource{Cluster: &cluster.Cluster{Name: "bad"}}
	for name, build := range map[string]func() error{
		"NewSnapshot": func() error {
			_, err := cache.NewSnapshot("1", map[resource.Type][]types.Resource{resource.ClusterType: {bad}})
			return err
		},
		"NewSnapshotWithTTLs": func() error {
			_, err := cache.NewSnapshotWithTTLs("1", map[resource.Type][]types.ResourceWithTTL{resource.ClusterType: {{Resource: bad}}})
			return err
		},
	} {
		gotType = ""
		if err := build(); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if gotType != resource.ClusterType {
			t.Errorf("%s: hook got type %q, want %q", name, gotType, resource.ClusterType)
		}
	}
}

// With no hook set, the free constructors have no logger either; the drop must
// still reach the standard logger instead of vanishing.
func TestDroppedResourceWithoutHookIsLogged(t *testing.T) {
	original := cache.OnResourceMarshalError
	t.Cleanup(func() { cache.OnResourceMarshalError = original })
	cache.OnResourceMarshalError = nil
	var buf bytes.Buffer
	prev := log.Writer()
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(prev) })

	cache.IndexAndMarshalResourcesByName([]types.ResourceWithTTL{
		{Resource: unmarshalableResource{Cluster: &cluster.Cluster{Name: "bad"}}},
	})

	if !strings.Contains(buf.String(), `dropping resource "bad"`) {
		t.Errorf("standard logger got %q, want the dropped resource reported", buf.String())
	}
}
