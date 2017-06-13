package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField) //TODO
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) IsVolatile() bool {
	return self.accessFlags&ACC_VOLATILE != 0
}

func (self *Field) IsTransient() bool {
	return self.accessFlags&ACC_TRANSIENT != 0
}

func (self *Field) IsEnum() bool {
	return self.accessFlags&ACC_ENUM != 0
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}
