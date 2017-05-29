package rtda

type Frame struct {
	lower        *Frame        `next frame`
	localVars    LocalVars     `局部变量表`
	operandStack *OperandStack `操作数栈`
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
