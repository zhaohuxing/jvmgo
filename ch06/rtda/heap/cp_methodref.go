package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ReslovedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() {

}
