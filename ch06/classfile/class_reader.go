package classfile

import "encoding/binary"

/*
	定义ClassReader结构体，里面存放bytes
	定义各种bytes向uint8, uint16, uint32, uint64, []uint16转换的方法
*/
type ClassReader struct {
	data []byte
}

//u1 无符号1字节
func (self *ClassReader) readUint8() uint8 {
	//读取u1类型数据
	val := self.data[0]
	self.data = self.data[1:] //读取后指针后移1byte
	return val

}

//u2 无符号2字节
func (self *ClassReader) readUint16() uint16 {
	//读取u2类型数据
	//Go标准库encoding/binary包中定义了一个变量BigEndian，
	//正好可以从[]byte中解码多字节数据.

	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//u4 无符号4字节
func (self *ClassReader) readUint32() uint32 {
	//读取U4类型数据
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

//u8 无符号8字节
func (self *ClassReader) readUint64() uint64 {
	//读取U8类型数据
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val

}

//读取uint16表,表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

//读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
