package terr

import (
	"fmt"
)

type _FrameLite struct {
	PC   uintptr
	Func string
	File string
	Line int
}

func (f *_FrameLite) String() string {
	return fmt.Sprintf("%#x: %s(%s:%d)", f.PC, f.Func, f.File, f.Line)
}
