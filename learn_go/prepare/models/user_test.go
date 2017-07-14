package models

import "testing"
import "fmt"

func TestInsert(t *testing.T) {
	user := User{
		PhoneNumber: "17862822537",
		Password:    "root123456",
	}

	if !InsertUser(user) {
		t.Log("添加失败")
	}
}

func TestGet(t *testing.T) {
	if !GetUser("17862822537", "root123456") {
		fmt.Println("获取失败")
		t.Log("添加失败")
	}
}
