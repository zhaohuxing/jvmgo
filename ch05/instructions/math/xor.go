package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// boolean xor int
type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

//boolean xor long
type LXOR struct {
	base.NoOperandsInstrcution
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stakc.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
