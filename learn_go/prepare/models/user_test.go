package models

import "testing"
import "fmt"

func TestInsert(t *testing.T) {
	/*	if InsertUser("17862822537", "root123456") {
		fmt.Println("添加成功")
	}*/
	fmt.Println(InsertUser("17862822537", "root123456"))
}

func TestGet(t *testing.T) {
	if GetUser("17862822537", "root123456") {
		fmt.Println("获取成功")
	}
}
