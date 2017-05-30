package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	//执行指令逻辑
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
	//没有操作数执行，所以没有定义任何字段
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

type BranchInstruction struct { //跳转指令
	Offset int `跳转偏移量`
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
