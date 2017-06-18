package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		method[i] = &Method{}
		method[i].class = class
		method[i].copyMemberInfo(cfMethod)
		method[i].copy(cfMethod)
	}
	return methods
}

func (self *Method) copyAttribute(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

//getters
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}

//访问标志: synchronized, bridge, varargs, native, abstract, strict
func (self *Method) IsSynchronized() bool {
	return self.accessFlags&ACC_SYNCHRONIZED != 0
}
func (self *Method) IsBridge() bool {
	return self.accessFlags&ACC_BRIDGE != 0
}
func (self *Method) IsVarargs() bool {
	return self.accessFlags&ACC_VARARGS != 0
}
func (self *Method) IsNative() bool {
	return self.accessFlags&ACC_NATIVE != 0
}
func (self *Method) IsAbstract() bool {
	return self.accessFlags&ACC_ABSTRACT != 0
}
func (self *Method) IsStrict() bool {
	return self.accessFlags&ACC_STRICT != 0
}
