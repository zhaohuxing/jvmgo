package classfile

/*
	constant_name_and_type_info {
		tag 	u1
		index	u2 名称索引
		index 	u2 修饰词索引
	}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
