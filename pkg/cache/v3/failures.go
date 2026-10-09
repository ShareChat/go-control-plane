package cache

import (
	"errors"
	stdlog "log"

	"github.com/envoyproxy/go-control-plane/pkg/log"
)

// ErrResponseChannelClosed is returned by respondDelta when the send panicked
// because the stream's response channel was already closed.
//
// It is deliberately distinct from a nil response with a nil error. That pair
// means "nothing to send, or the channel was full - keep the watch and retry on
// the next upsert", and retrying is correct there. A closed channel means the
// stream is gone and no retry will ever happen, so collapsing the two hides a
// dead watch behind a healthy-looking one.
var ErrResponseChannelClosed = errors.New("delta response channel is closed")

// OnResourceMarshalError is called when a resource cannot be marshalled while
// building or updating a snapshot. The resource is dropped: it is not added to
// the snapshot, so it is never pushed, and a proxy already holding it sees it
// disappear on the next diff.
//
// None of those call sites can return an error, so this hook is the only way to
// count the drop. Set it once at startup; it must be safe for concurrent use.
// Leaving it nil sends drops seen by the logger-less snapshot constructors
// (NewSnapshot, NewResources, IndexAndMarshalResourcesByName) to the standard logger.
var OnResourceMarshalError func(typeURL, name string, err error)

// reportMarshalError records a dropped resource. logger may be nil; with no hook
// either, the standard logger is used so the drop is never silent.
func reportMarshalError(logger log.Logger, typeURL, name string, err error) {
	if logger == nil && OnResourceMarshalError == nil {
		logger = log.LoggerFuncs{ErrorFunc: stdlog.Printf}
	}
	if logger != nil {
		logger.Errorf("dropping resource %q of type %q from the snapshot: MarshalVTStrict failed: %v; "+
			"it will not be pushed to any proxy", name, typeURL, err)
	}
	if OnResourceMarshalError != nil {
		OnResourceMarshalError(typeURL, name, err)
	}
}
