package terr

import (
	"errors"
	"fmt"
)

// Trace returns the original value and a [TracedError] if `err` is not nil.
//
// Example:
//
//	resp, err := terr.Trace(http.Get("https://example.com"))
func Trace[V any](val V, err error) (V, error) {
	return val, TraceError(err)
}

// TraceError returns a [TracedError] if `cause` is not nil.
//
// When `cause` is already a [TracedError], it will be returned directly.
func TraceError(cause error) error {
	if cause == nil {
		return nil
	}
	if te, ok := cause.(*TracedError); ok {
		return te
	}
	return fillFrames(&TracedError{
		base: cause,
	}, 4)
}

// New makes a [TracedError] with specific message.
func New(message string) error {
	return fillFrames(&TracedError{
		base: errors.New(message),
	}, 3)
}

// Errorf makes a [TracedError] with a formatted error.
func Errorf(format string, a ...any) error {
	return fillFrames(&TracedError{
		base: fmt.Errorf(format, a...),
	}, 3)
}

// Wrap makes a [TracedError] whose original error is wrapped from `cause`
// with `message`.
//
// When `cause` is a TracedError, its frames will be copied to new error.
func Wrap(message string, cause error) error {
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
