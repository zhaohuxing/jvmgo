package list

type List interface {
	//先小写试试
	//小写在当前包下可以list.add(e)调用， 当换了一个包需要list.Add(e)
	Add(e interface{}) bool
	Get(index int32) interface{}
	Remove(index int32) interface{}
	Size() int32
}
