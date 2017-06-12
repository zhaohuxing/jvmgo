package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}

func (self *ClassMember) IsPrivate() bool {
	return self.accessFlags&ACC_PRIVATE != 0
}

func (self *ClassMember) IsProtected() bool {
	return self.accessFlags&ACC_PROTECTED != 0
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

//getter
func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}
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
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
