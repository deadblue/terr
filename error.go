package terr

import (
	"iter"
	"strings"
)

// TracedError represent an error with calling stack.
type TracedError struct {
	base   error
	frames []FrameLite
}

// Error returns error message with calling stack.
func (e *TracedError) Error() string {
	sb := &strings.Builder{}
	sb.WriteString(e.base.Error())
	for _, frame := range e.frames {
		sb.WriteRune('\n')
		sb.WriteString(frame.String())
	}
	return sb.String()
}

// Unwrap returns the original error.
func (e *TracedError) Unwrap() error {
	return e.base
}

// Stack returns calling stack.
func (e *TracedError) Stack() iter.Seq[FrameLite] {
	return func(yield func(FrameLite) bool) {
		for _, frame := range e.frames {
			if !yield(frame) {
				break
			}
		}
	}
}
