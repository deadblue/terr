package terr

import (
	"errors"
	"fmt"
)

// New makes a [TracedError] with specific message.
func New(message string) *TracedError {
	return fillFrames(&TracedError{
		base: errors.New(message),
	}, 3)
}

// TraceError wraps |err| to [TracedError] when it is not nil.
//
// When |base| is already a TracedError, it will be returned directly.
//
// Example:
//
//	resp, err := terr.TraceError(http.Get("https://example.com"))
func TraceError[V any](val V, err error) (V, error) {
	return val, traceError(err)
}

// Trace makes a [TracedError] whose original error is |base|.
//
// When |base| is already a TracedError, it will be returned directly.
func Trace(base error) *TracedError {
	return traceError(base)
}

func traceError(base error) *TracedError {
	if base == nil {
		return nil
	}
	if te, ok := base.(*TracedError); ok {
		return te
	}
	return fillFrames(&TracedError{
		base: base,
	}, 4)
}

// Wrap makes a [TracedError] whose original error is wrapped from cause with message.
//
// When |cause| is a TracedError, its frames will be copied to new error.
func Wrap(message string, cause error) *TracedError {
	if te, ok := cause.(*TracedError); ok {
		nte := &TracedError{
			base: fmt.Errorf("%s: %w", message, te.base),
		}
		// Copy frames
		nte.frames = append(nte.frames, te.frames...)
		return nte
	}
	return fillFrames(&TracedError{
		base: fmt.Errorf("%s: %w", message, cause),
	}, 3)
}
