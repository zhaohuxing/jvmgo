package classfile

import "math"

/*
	Constant_Integer_info {
		tag 	u1
		bytes	u4
	}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) { //实现ConstantInfo
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

/*
	Constant_Float_info {
		tag 	u1
		bytes	u4
	}
*/
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/*
	Constant_Long_info {
		tag 	u1
		bytes	u4
	}
*/
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

/*
	Constant_Double_info {
		tag		u1
		bytes	u4
	}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
