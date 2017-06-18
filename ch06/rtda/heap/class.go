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
//修饰类的访问标志符如下: public final interface abstract super annotation enum synthetic
func (self *Class) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}
func (self *Class) IsFinal() bool {
	return self.accessFlags&ACC_FINAL != 0
}
func (self *Class) IsSuper() bool {
	return self.accessFlags&ACC_SUPER != 0
}
func (self *Class) IsAbstract() bool {
	return self.accessFlags&ACC_ABSTRACT != 0
}
func (self *Class) IsInterface() bool {
	return self.accessFlags&ACC_INTERFACE != 0
}
func (self *Class) IsAnnotation() bool {
	return self.accessFlags&ACC_ANNOTATION != 0
}
func (self *Class) IsEnum() bool {
	return self.accessFlags&ACC_ENUM != 0
}
func (self *Class) IsSynthetic() bool {
	return self.accessFlags & ACC_SYNTHETIC
}

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

//get main method
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

//NewOject
func (self *Class) NewObject() *Object {
	return newObject(self) // newObject(self) 方法位于heap/object.go中
}
