package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}

//aconst_null指令把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

type DCONST_0 struct {
	base.NoOperandsInstruction
}

//dconst_0指令把double型0推入操作数栈顶
func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

//dconst_1指令把double型1推入操作数栈顶
func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_0 struct {
	base.NoOperandsInstruction
}

//fconst_0指令把float型0推入操作数栈顶
func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

//fconst_1指令把float型1推入操作数栈顶
func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

//fconst_2指令把float型2推入操作数栈顶
func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

type ICONST_M1 struct {
	base.NoOperandsInstruction
}

//iconst_m1指令把int型-1推入操作数栈顶
func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
