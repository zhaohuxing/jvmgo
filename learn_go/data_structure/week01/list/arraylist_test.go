package list

import "testing"
import "log"

func TestArrayList(t *testing.T) {
	var list List = &ArrayList{}
	for i := 0; i < 5; i++ {
		list.add(i)
	}

	if !list.add(5) {
		log.Fatal("list.add(5)操作失败")
	}

	if list.Size() != 6 {
		log.Fatal("list.Size()操作失败")
	}

	if list.get(3) != 3 {
		log.Fatal("list.Get(", 3, ")操作失败")
	}

	list.remove(3)
	log.Println(list)
}
