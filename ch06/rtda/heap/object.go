package rtda

type Object struct {
	class  *Class //TODO
	fields Slots  //TODO
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount), //TODO
	}
}

//getter
func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.fields
}

func (sefl *Object) IsInstanceOf(class *Class) bool {
	return class.isAssginableFrom(self.class)
}
