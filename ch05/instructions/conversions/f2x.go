package conversions

//TODO
import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//Convert float to double

type F2D struct {
	base.NoOperandsInstruction
}

func (self *F2D) Execute(frame *rtda.Frame) {

	d := float64(getFloat(frame))
	frame.OperandStack.PushDouble(d)
}

//Convert float to int

type F2I struct {
	base.NoOperandsInstruction
}

func (self *F2I) Execute(frame *rtda.Frame) {
	i := int32(getFloat(frame))
	frame.OperandStack.PushInt(i)
}

//Convert float to long
type F2L struct {
	base.NoOperandsInstruction
}

func (self *F2L) Execute(frame *rtda.Frame) {
	i := int64(getFloat(frame))
	frame.OperandStack().PushLong(i)

}

func getFloat(frame *rtda.Frame) float32 {
	stack := frame.OperandStack()
	return stack.PopFloat()
}
