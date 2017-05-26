package rtda

type Frame struct {
	lower        *Frame        //下一节点
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
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
