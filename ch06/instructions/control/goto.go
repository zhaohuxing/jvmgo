package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
