package rtda

import "jvmgo/ch06/rtda/heap"

type Thread struct {
	pc    int `存放字节码执行的行号`
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

//存放栈帧
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

//获取栈帧
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

//获取top栈帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

//NewFrame 方法
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(method)
}
