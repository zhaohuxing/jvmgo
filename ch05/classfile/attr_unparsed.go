package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

//TODO 难道UnparsedAttribute 实现了AttributeInfo和ConstantInfo
func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
