package cache

import (
	"fmt"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"time"
)

// Resources is a versioned group of resources.
type Resources struct {
	// Version information.
	Version string

	// Items in the group indexed by name.
	Items map[string]VTMarshaledResource
}

type VTMarshaledResource struct {
	Name     string
	Version  string
	Resource []byte
	TTL      *time.Duration
}

// IndexAndMarshalResourcesByName creates a map from the resource name to the marshaled resource.
func IndexAndMarshalResourcesByName(items []types.ResourceWithTTL) map[string]VTMarshaledResource {
	indexed := make(map[string]VTMarshaledResource, len(items))
	for _, item := range items {
		out, err := item.Resource.MarshalVTStrict()
		if err != nil {
			fmt.Printf("failed to MarshalVTStrict resource %s: %v\n", GetResourceName(item.Resource), err)
			continue
		}
		indexed[GetResourceName(item.Resource)] = VTMarshaledResource{
			Name:     GetResourceName(item.Resource),
			Version:  item.Version,
			Resource: out,
		}
	}
	return indexed
}

// IndexResourcesByName creates a map from the resource name to the resource.
func IndexResourcesByName(items []types.ResourceWithTTL) map[string]types.ResourceWithTTL {
	indexed := make(map[string]types.ResourceWithTTL, len(items))
	for _, item := range items {
		indexed[GetResourceName(item.Resource)] = item
	}
	return indexed
}

// IndexRawResourcesByName creates a map from the resource name to the resource.
func IndexRawResourcesByName(items []types.Resource) map[string]types.Resource {
	indexed := make(map[string]types.Resource, len(items))
	for _, item := range items {
		indexed[GetResourceName(item)] = item
	}
	return indexed
}

// NewResources creates a new resource group.
func NewResources(version string, items []types.Resource) Resources {
	itemsWithTTL := make([]types.ResourceWithTTL, 0, len(items))
	for _, item := range items {
		itemsWithTTL = append(itemsWithTTL, types.ResourceWithTTL{Resource: item})
	}
	return NewResourcesWithTTL(version, itemsWithTTL)
}

// NewResourcesWithTTL creates a new resource group.
func NewResourcesWithTTL(version string, items []types.ResourceWithTTL) Resources {
	return Resources{
		Version: version,
		Items:   IndexAndMarshalResourcesByName(items),
	}
}
