package terr

import "runtime"

const (
	_MaxFrames = 16
)

func fillFrames(te *_TraceableError) *_TraceableError {
	// Get calling stack, skip this function
	callers := make([]uintptr, _MaxFrames)
	// Skip 3 frames
	if n := runtime.Callers(3, callers); n > 0 {
		te.frames = make([]_FrameLite, 0, n)
		frames := runtime.CallersFrames(callers[:n])
		for {
			frame, more := frames.Next()
			te.frames = append(te.frames, _FrameLite{
				PC:   frame.PC,
				Func: frame.Function,
				File: frame.File,
				Line: frame.Line,
			})
			if !more {
				break
			}
		}
	}
	return te
}
