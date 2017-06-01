//取反
package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DNEG struct {
	base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := frame.PopFloat()
	stack.PushDouble(-val)
}

type INEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := frame.PopInt()
	stack.PushInt(-val)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := frame.PopLong()
	stack.PushLong(-val)
}
