package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool //TODO
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

//TODO
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	// AccessFlags(), ClassName(), SuperClassName(), InterfaceNames()
	// ConstantPool(), Fields(), Methods()
	//位于/classfile/class_file.go中
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool()) //newConstantPool(...)　位于.../heap/constant_pool.go中
	class.fields = newFields(class, cf.Fields())                   // newFields(...) 位于.../heap/field.go中
	class.methods = newMethods(class, cf.Methods())                // newMethods(...) 位于.../heap/method.go中
	return class
}

//getters
func (self *Class) ConstantPool() *ContantPool {
	return self.constantPool
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

//访问标志

//
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

//
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}
