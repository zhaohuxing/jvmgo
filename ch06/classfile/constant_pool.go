package classfile

//因常量池是由一个个常量表组成的，故将slice的ConstantInfo作为ConstantPool
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	// 因JVM中不存在任何分割符，所以格式要求十分严格(固定),
	//字节码中版本号之后是常量池的大小（用u2来表示）,再之后是常量池内容
	//常量池的大小是从1开始的，0代表无效。总长度 <= 常量池大小-1
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		//索引从1开始, 0无效
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //占两个位置
		}
	}
	return cp
}

//获取常量池中的信息
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	/*
		JVM规范中，共有14中常量项:
		每个常量项又是一个表，但是共同点就是：表中存在唯一标识该常量的类型的tag
		举例:Constant_Utf8_info的tag = 1, Constant_Integer_info的tag = 3
	*/
	if cpInfo := self[index]; cpInfo != nil { //index 指　常量项表中的tag
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

//从常量池中查找字段和方法名
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	//大概意思是获取ConstantInfo的实现者:ConstantNameAndTypeInfo
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

//从常量池中查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

//从常量池中查找Utf8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
