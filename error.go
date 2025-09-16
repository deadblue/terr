package terr

import (
	"strings"
)

type _TraceableError struct {
	cause  error
	frames []_FrameLite
}

func (e *_TraceableError) Error() string {
	sb := &strings.Builder{}
	sb.WriteString(e.cause.Error())
	for _, frame := range e.frames {
		sb.WriteRune('\n')
		sb.WriteString(frame.String())
	}
	return sb.String()
}

func (e *_TraceableError) Unwarp() error {
	return e.cause
}
