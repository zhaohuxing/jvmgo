package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember     // 来自.../heap/class_member.go文件夹下
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	//newFields没有为slotId赋值
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

//访问标志
func (self *Field) IsVolatile() bool {
	return self.accessFlags&ACC_VOLATILE != 0
}

func (self *Field) IsTransient() bool {
	return self.accessFlags&ACC_TRANSIENT != 0
}

func (self *Field) IsEnum() bool {
	return self.accessFlags&ACC_ENUM != 0
}

//getters
func (self *Field) ConstantValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}

//
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}
