package terr

import "runtime"

const (
	_MaxFrames = 16
)

func fillFrames(te *TracedError, skip int) *TracedError {
	// Get calling stack, skip this function
	callers := make([]uintptr, _MaxFrames)
	// Skip 3 frames
	if n := runtime.Callers(skip, callers); n > 0 {
		te.frames = make([]FrameLite, 0, n)
		frames := runtime.CallersFrames(callers[:n])
		for {
			frame, more := frames.Next()
			te.frames = append(te.frames, FrameLite{
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
