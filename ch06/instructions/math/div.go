package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type DDIV struct {
	base.NoOperandsInstruction
}

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

type IDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 / v2
	stack.PushInt(result)
}

type LDIV struct {
	base.NoOperandsInstruction
}

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 / v2
	stack.PushLong(result)
}
