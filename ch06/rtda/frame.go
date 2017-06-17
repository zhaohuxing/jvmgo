package rtda

type Frame struct {
	lower        *Frame        `next frame`
	localVars    LocalVars     `局部变量表`
	operandStack *OperandStack `操作数栈`
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (frame Frame) LocalVars() LocalVars {
	return frame.localVars
}

func (frame Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
