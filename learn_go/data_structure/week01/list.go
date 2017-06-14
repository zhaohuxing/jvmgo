package main

// 问题来了？
// next, prev 的类型为啥是*Element, 而不是Element
// list *List 的类型而不是List
type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type List struct {
	root Element
	len  int
}

//Init initializes or clears list l
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List {
	//new 申请内存初始化
	return new(List).Init()
}

func (l *List) Len() {
	return l.len
}

//TODO
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

//TODO
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

//insert e after at
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.size++
	return e

}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	//创建新的节点
	return l.insert(&Element{Value: v}, at)
}

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	//avoid memory leaks
	e.prev = nil
	e.next = nil
	e.list = nil
	l.len--
	return e
}

func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

//TODO 为啥是 &l.root
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

//TODO 为啥是 &l.root.prev
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root.prev)
}

//TODO  为啥是mark.prev
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	//IF mark is not an eleement of l, the list is not modified
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

//TODO 我肯定有个地方弄混了
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

//将e放到链表的头部
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.insert(l.remove(e), &l.root)
}

//是不是需要搞懂它的数据结构呀？

func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.insert(l.remove(e), l.root.prev)
}

//TODO
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.insert(l.remove(e), mark.prev)
}

//TODO
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.insert(l.remove(e), mark)
}

//TODO
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; e = i - 1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
