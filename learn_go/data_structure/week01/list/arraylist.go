package list

type ArrayList struct {
	size int32
	data []interface{}
}

func (self *ArrayList) add(e interface{}) bool {
	//先不设边界条件
	if e == nil {
		return false
	}
	self.data = append(self.data, e)
	self.size++
	return true
}

func (self *ArrayList) get(index int32) interface{} {
	if self.size == 0 || self.size < index {
		return nil
	}

	return self.data[index]
}

func (self *ArrayList) remove(index int32) interface{} {
	//假设index范围合法
	old := self.data[index]
	self.data = append(self.data[:index], self.data[index+1:]...)
	self.size--
	return old
}

func (self *ArrayList) Size() int32 {
	return self.size
}
