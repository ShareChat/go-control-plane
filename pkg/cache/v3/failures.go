package cache

import (
	"errors"
	stdlog "log"

	"github.com/envoyproxy/go-control-plane/pkg/log"
)

// ErrResponseChannelClosed is returned by respond and respondDelta when the send
// panicked because the stream's response channel was already closed.
//
// It is deliberately distinct from a nil response with a nil error. That pair
// means "nothing to send, or the channel was full - keep the watch and retry on
// the next upsert", and retrying is correct there. A closed channel means the
// stream is gone and no retry will ever happen, so collapsing the two hides a
// dead watch behind a healthy-looking one.
var ErrResponseChannelClosed = errors.New("response channel is closed")

// OnResourceMarshalError is called when a resource cannot be marshalled while
// building or updating a snapshot. When building one (NewSnapshot and friends)
// the resource is left out, so a proxy holding it sees it removed on the next
// diff. When updating one (UpsertResources, BatchUpsertResources) the update is
// skipped: a new resource is not added, and an existing one keeps its previous
// value and is still served, while the type version is still bumped.
//
// None of those call sites can return an error, so this hook is the only way to
// count the drop. Set it once at startup; it must be safe for concurrent use.
// Leaving it nil sends drops seen by the logger-less snapshot constructors
// (NewSnapshot, NewResources, IndexAndMarshalResourcesByName) to the standard logger.
var OnResourceMarshalError func(typeURL, name string, err error)

// reportMarshalError records a dropped resource. logger may be nil, or the no-op
// default a cache gets when built without one; with no hook either, the standard
// logger is used so the drop is never silent.
func reportMarshalError(logger log.Logger, typeURL, name string, err error) {
	if _, noop := logger.(*log.DefaultLogger); (logger == nil || noop) && OnResourceMarshalError == nil {
		logger = log.LoggerFuncs{ErrorFunc: stdlog.Printf}
	}
	if logger != nil {
		logger.Errorf("dropping resource %q of type %q: MarshalVTStrict failed: %v; "+
			"this value is not pushed (an existing resource keeps its previous value)", name, typeURL, err)
	}
	if OnResourceMarshalError != nil {
		OnResourceMarshalError(typeURL, name, err)
	}
}
