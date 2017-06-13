package heap

type Object struct {
	//TODO
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

//getter

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.fields
}
