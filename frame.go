package terr

import (
	"fmt"
)

// FrameLite is a lite version of [runtime.Frame].
type FrameLite struct {
	// PC is the program counter for the location in this frame.
	PC uintptr
	// Func is the function name when available.
	Func string
	// File is the name of source file when available.
	File string
	// Line is the line number in source file when available.
	Line int
}

func (f *FrameLite) String() string {
	return fmt.Sprintf("%#x: %s(%s:%d)", f.PC, f.Func, f.File, f.Line)
}
