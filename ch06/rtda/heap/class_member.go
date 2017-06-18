package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

//将MemberInfo中的信息赋值到ClassMember
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

//getters
func (self *ClassMember) Name() {
	return self.name
}

func (self *ClassMember) Descriptor() {
	return self.descriptor
}

func (self *ClassMember) Class() {
	return self.class
}

//方法标志
// public protected private static final synthetic
func (self *ClassMember) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}

func (self *ClassMember) IsProtected() bool {
	return self.accessFlags&ACC_PROTECTED != 0
}

func (self *ClassMember) IsPrivate() bool {
	return self.accessFlags&ACC_PRIVATE != 0
}

func (self *ClassMember) IsStatic() bool {
	return self.accessFlags&ACC_STATIC != 0
}

func (self *ClassMember) IsFinal() bool {
	return self.accessFlags&ACC_FINAL != 0
}

func (self *ClassMember) IsSynthetic() bool {
	return self.accessFlags&ACC_SYNTHETIC != 0
}

//jvm 5.4.4 什么鬼?
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}

	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackgaeName()
	}
	return d == c
}
