// Copyright 2020 Envoyproxy Authors
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

package cache

import (
	"context"
	"strings"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// groups together resource-related arguments for the createDeltaResponse function
type resourceContainer struct {
	resourceMap   map[string]types.ResourceWithTTL
	versionMap    map[string]string
	systemVersion string
}

func contains(s []string, elem string) bool {
	for _, a := range s {
		if a == elem {
			return true
		}
	}
	return false
}

func combineUnique(str1 []string, str2 []string) []string {
	res := make([]string, 0)
	for _, s := range str1 {
		if !contains(res, s) {
			res = append(res, s)
		}
	}

	for _, s := range str2 {
		if !contains(res, s) {
			res = append(res, s)
		}
	}

	return res
}

func containsPrefixedKey(data map[string]string, keyLike string) ([]string, bool) {
	resNames := make([]string, 0)
	for key := range data {
		if strings.Contains(key, keyLike) {
			resNames = append(resNames, key)
		}
	}

	return resNames, len(resNames) > 0
}

func containsPrefixedKeyResources(data map[string]types.ResourceWithTTL, keyLike string) ([]string, bool) {
	resNames := make([]string, 0)
	for key := range data {
		if strings.Contains(key, keyLike) {
			resNames = append(resNames, key)
		}
	}
	return resNames, len(resNames) > 0
}

func createDeltaResponse(ctx context.Context, req *DeltaRequest, state stream.StreamState, resources resourceContainer) *RawDeltaResponse {
	// variables to build our response with
	var nextVersionMap map[string]string
	var filtered map[string]types.ResourceWithTTL
	var toRemove []string
	switch {
	case state.IsWildcard():
		filtered = make(map[string]types.ResourceWithTTL)
		nextVersionMap = make(map[string]string, len(resources.resourceMap))
		for name, r := range resources.resourceMap {
			// Since we've already precomputed the version hashes of the new snapshot,
			// we can just set it here to be used for comparison later
			version := resources.versionMap[name]
			nextVersionMap[name] = version
			prevVersion, found := state.GetResourceVersions()[name]
			if !found || (prevVersion != version) {
				filtered[name] = r
			}
		}

		// Compute resources for removal
		// The resource version can be set to "" here to trigger a removal even if never returned before
		for name := range state.GetResourceVersions() {
			if _, ok := resources.resourceMap[name]; !ok {
				toRemove = append(toRemove, name)
			}
		}
	default:
		filtered = make(map[string]types.ResourceWithTTL)
		nextVersionMap = make(map[string]string, len(state.GetSubscribedResourceNames()))
		// state.GetResourceVersions() may include resources no longer subscribed
		// In the current code this gets silently cleaned when updating the version map
		for name := range state.GetSubscribedResourceNames() {
			dirResourceName := name
			if strings.Contains(dirResourceName, "*") {
				dirResourceName = strings.Split(dirResourceName, "*")[0]
				prevVersions, _ := containsPrefixedKey(state.GetResourceVersions(), dirResourceName)
				currVersions, _ := containsPrefixedKeyResources(resources.resourceMap, dirResourceName)
				combinedVersions := combineUnique(prevVersions, currVersions)

				for _, versionName := range combinedVersions {
					prevVersion, found := state.GetResourceVersions()[versionName]
					if r, ok := resources.resourceMap[versionName]; ok {
						nextVersion := resources.versionMap[versionName]
						if prevVersion != nextVersion {
							filtered[GetResourceName(r.Resource)] = r
						}
						nextVersionMap[versionName] = nextVersion
					} else if found {
						toRemove = append(toRemove, versionName)
					}
				}
			} else {
				prevVersion, found := state.GetResourceVersions()[name]
				if r, ok := resources.resourceMap[name]; ok {
					nextVersion := resources.versionMap[name]
					if prevVersion != nextVersion {
						filtered[GetResourceName(r.Resource)] = r
					}
					nextVersionMap[name] = nextVersion
				} else if found {
					toRemove = append(toRemove, name)
				}
			}
		}
	}

	filteredResources := make([]types.ResourceWithTTL, 0)
	filteredResourceNames := make([]string, 0)
	for name, r := range filtered {
		filteredResources = append(filteredResources, r)
		filteredResourceNames = append(filteredResourceNames, name)
	}

	return &RawDeltaResponse{
		DeltaRequest:      req,
		Resources:         filteredResources,
		RemovedResources:  toRemove,
		NextVersionMap:    nextVersionMap,
		SystemVersionInfo: resources.systemVersion,
		Ctx:               ctx,
	}
}
