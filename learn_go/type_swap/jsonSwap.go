package main

import (
	"encoding/json"
	"fmt"
)

// 将struct实例解析成json
// 将json数据解析到struct实例

//首先需要一个struct
type Admin struct {
	Id       int `json:",string"` //为啥这么标记后就为0
	Username string
	Password string
}

// Otoj: struct --- > json

func (self *Admin) Otoj() {
	bytes, err := json.Marshal(self)
	if err != nil {
		fmt.Println("解析失败: struct to json")
		return
	}
	fmt.Println(string(bytes))
}

//Jtoo: json --- > struct

func (self *Admin) Jtoo(values string) {
	fmt.Println("将json数据解析到Admin前:")
	fmt.Println("ID:", self.Id)
	fmt.Println("Username:", self.Username)
	fmt.Println("Password:", self.Password)
	json.Unmarshal([]byte(values), self)
	fmt.Println("将json数据解析到Admin后:")
	fmt.Println("ID:", self.Id)
	fmt.Println("Username:", self.Username)
	fmt.Println("Password:", self.Password)
}
