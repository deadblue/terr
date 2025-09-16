package terr

import (
	"errors"
	"fmt"
)

// New makes a [TracedError] with specific message.
func New(message string) *TracedError {
	return fillFrames(&TracedError{
		base: errors.New(message),
	})
}

// Trace makes a [TracedError] whose original error is |base|.
//
// When |base| is already a TracedError, it will be returned directly.
func Trace(base error) *TracedError {
	if te, ok := base.(*TracedError); ok {
		return te
	}
	return fillFrames(&TracedError{
		base: base,
	})
}

// Wrap makes a [TracedError] whose original error is wrapped from cause with message.
//
// When |cause| is already a TracedError, its frames will be copied to new error.
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
	})
}
