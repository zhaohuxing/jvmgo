package main

import "testing"
import "fmt"

func TestFunc(t *testing.T) {
	/*	l := New()
		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		l.PushFront(4)

		fmt.Println("root.prev", l.Back().Value)
		fmt.Println("root.next", l.Front().Value)
		fmt.Println("root", l.root.Value, "prev:", *l.root.prev, "next:", *l.root.next, "list:", l.root)
		// Iterate through list and print its contents.
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}

		for e := l.Back(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}*/
}

func TestFuncc1(t *testing.T) {
	l := New()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	fmt.Println("l.Front()方法:", l.Front())
}

func TestFunc2(t *testing.T) {

	fmt.Println("TESTFun2")
	l := New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	fmt.Println("root信息:", &l.root)
}

func TestFun3(t *testing.T) {
	l := New()
	l.PushFront(1)
	l.PushFront(2.3)
	l.PushFront("你好")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestFun4(t *testing.T) {
	l := New()
	l.PushFront(1)
	e := l.Front()
	fmt.Println(e.prev, ":", e.next)
}
