package main

import "jvmgo/learn_go/data_structure/week01/list"

//import "log"

func main() {
	var list list.List = &list.ArrayList{}
	for i := 0; i < 5; i++ {
		list.add(i)
	}

}
