// Copyright 2018 Envoyproxy Authors
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
	"sort"
	"sync"
	"time"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// NodeHash computes string identifiers for Envoy nodes.
type NodeHash interface {
	// ID function defines a unique string identifier for the remote Envoy node.
	ID(node *core.Node) string
	CacheIndex(node *core.Node) int
	CacheIndexFromKey(key string) int
}

// IDHash uses ID field as the node hash.
type IDHash struct{}

// ID uses the node ID field
func (IDHash) ID(node *core.Node) string {
	if node == nil {
		return ""
	}
	return node.GetId()
}

// CacheIndex returns the index of the node in the cache.
// Don't use this default implementation, write your own CacheIndex function.
func (IDHash) CacheIndex(node *core.Node) int {
	if node == nil {
		return 0
	}
	return len(node.GetId())
}

// CacheIndexFromKey returns the index of the node in the cache.
// Don't use this default implementation, write your own CacheIndex function.
func (IDHash) CacheIndexFromKey(key string) int {
	return len(key)
}

var _ NodeHash = IDHash{}

// StatusInfo publishes information about nodes that are watching the xDS cache.
type StatusInfo interface {
	// GetNode returns the node metadata.
	GetNode() *core.Node

	// GetNumWatches returns the number of open watches.
	GetNumWatches() int

	// GetNumDeltaWatches returns the number of open delta watches.
	GetNumDeltaWatches() int

	// GetLastWatchRequestTime returns the timestamp of the last discovery watch request.
	GetLastWatchRequestTime() time.Time

	// GetLastDeltaWatchRequestTime returns the timestamp of the last delta discovery watch request.
	GetLastDeltaWatchRequestTime() time.Time
}

// statusInfo tracks the server state for the remote Envoy node.
type statusInfo struct {
	// node is the constant Envoy node metadata.
	node *core.Node

	// watches are indexed channels for the response watches and the original requests.
	watches        map[int64]ResponseWatch
	orderedWatches keys

	// deltaWatches are indexed channels for the delta response watches and the original requests
	deltaWatches        map[int64]DeltaResponseWatch
	orderedDeltaWatches keys

	// the timestamp of the last watch request
	lastWatchRequestTime time.Time

	// the timestamp of the last delta watch request
	lastDeltaWatchRequestTime time.Time

	// mutex to protect the status fields.
	// should not acquire mutex of the parent cache after acquiring this mutex.
	mu sync.RWMutex
}

// ResponseWatch is a watch record keeping both the request and an open channel for the response.
type ResponseWatch struct {
	// Request is the original request for the watch.
	Request *Request

	// Response is the channel to push responses to.
	Response chan Response
}

// DeltaResponseWatch is a watch record keeping both the delta request and an open channel for the delta response.
type DeltaResponseWatch struct {
	// Request is the most recent delta request for the watch
	Request *DeltaRequest

	// Response is the channel to push the delta responses to
	Response chan DeltaResponse

	// VersionMap for the stream
	StreamState stream.StreamState
}

// newStatusInfo initializes a status info data structure.
func newStatusInfo(node *core.Node) *statusInfo {
	out := statusInfo{
		node:           node,
		watches:        make(map[int64]ResponseWatch),
		orderedWatches: make(keys, 0),
		deltaWatches:   make(map[int64]DeltaResponseWatch),
	}
	return &out
}

func (info *statusInfo) GetNode() *core.Node {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return info.node
}

func (info *statusInfo) GetNumWatches() int {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return len(info.watches)
}

func (info *statusInfo) GetNumDeltaWatches() int {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return len(info.deltaWatches)
}

func (info *statusInfo) GetLastWatchRequestTime() time.Time {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return info.lastWatchRequestTime
}

func (info *statusInfo) GetLastDeltaWatchRequestTime() time.Time {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return info.lastDeltaWatchRequestTime
}

// setLastDeltaWatchRequestTime will set the current time of the last delta discovery watch request.
func (info *statusInfo) setLastDeltaWatchRequestTime(t time.Time) {
	info.mu.Lock()
	defer info.mu.Unlock()
	info.lastDeltaWatchRequestTime = t
}

// Removed: setDeltaResponseWatch. CreateDeltaWatch now holds info.mu across
// "decide whether the client is up to date" and "register the watch" so the two
// are atomic against respondDeltaWatches, and a self-locking setter called from
// under that lock would deadlock. Registration is done inline instead.

// orderResponseWatches will track a list of watch keys and order them if
// true is passed.
func (info *statusInfo) orderResponseWatches() {
	info.orderedWatches = make(keys, len(info.watches))

	var index int
	for id, watch := range info.watches {
		info.orderedWatches[index] = key{
			ID:      id,
			TypeURL: watch.Request.GetTypeUrl(),
		}
		index++
	}

	// Sort our list which we can use in the SetSnapshot functions.
	// This is only run when we enable ADS on the cache.
	sort.Sort(info.orderedWatches)
}

// getOrderedWatches creates and returns a local copy of ordered watch keys.
// This avoids race conditions by not using a shared struct field.
func (info *statusInfo) getOrderedWatches() keys {
	orderedWatches := make(keys, len(info.watches))

	var index int
	for id, watch := range info.watches {
		orderedWatches[index] = key{
			ID:      id,
			TypeURL: watch.Request.GetTypeUrl(),
		}
		index++
	}

	// Sort our list which we can use in the SetSnapshot functions.
	// This is only run when we enable ADS on the cache.
	sort.Sort(orderedWatches)
	return orderedWatches
}

// orderResponseDeltaWatches will track a list of delta watch keys and order them if
// true is passed.
func (info *statusInfo) orderResponseDeltaWatches() {
	info.orderedDeltaWatches = make(keys, len(info.deltaWatches))

	var index int
	for id, deltaWatch := range info.deltaWatches {
		info.orderedDeltaWatches[index] = key{
			ID:      id,
			TypeURL: deltaWatch.Request.GetTypeUrl(),
		}
		index++
	}

	// Sort our list which we can use in the SetSnapshot functions.
	// This is only run when we enable ADS on the cache.
	sort.Sort(info.orderedDeltaWatches)
}

// getOrderedDeltaWatches creates and returns a local copy of ordered delta watch keys.
// This avoids race conditions by not using a shared struct field.
func (info *statusInfo) getOrderedDeltaWatches() keys {
	orderedWatches := make(keys, len(info.deltaWatches))

	var index int
	for id, deltaWatch := range info.deltaWatches {
		orderedWatches[index] = key{
			ID:      id,
			TypeURL: deltaWatch.Request.GetTypeUrl(),
		}
		index++
	}

	// Sort our list which we can use in the SetSnapshot functions.
	// This is only run when we enable ADS on the cache.
	sort.Sort(orderedWatches)
	return orderedWatches
}
