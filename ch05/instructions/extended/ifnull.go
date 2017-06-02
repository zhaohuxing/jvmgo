package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *frame.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

type IFNOTNULL struct {
	base.BranchInstruction
}

func (self *IFNOTNULL) Execute(frame *frame.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
