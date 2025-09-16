package terr

import (
	"errors"
	"fmt"
)

// New makes a traceable error with message.
func New(message string) error {
	return fillFrames(&_TraceableError{
		cause: errors.New(message),
	})
}

// Trace makes a traceable error whose cause is |cause|.
// When |cause| is already a traceable error, it will be returned directly.
func Trace(cause error) error {
	if te, ok := cause.(*_TraceableError); ok {
		return te
	}
	return fillFrames(&_TraceableError{
		cause: cause,
	})
}

// Trace makes a wrapped traceable error.
// When |cause| is already a traceable error, its frames will be copied to new error.
func Wrap(message string, cause error) error {
	if te, ok := cause.(*_TraceableError); ok {
		nte := &_TraceableError{
			cause: fmt.Errorf("%s: %w", message, te.cause),
		}
		// Copy frames
		nte.frames = append(nte.frames, te.frames...)
		return nte
	}
	return fillFrames(&_TraceableError{
		cause: fmt.Errorf("%s: %w", message, cause),
	})
}
