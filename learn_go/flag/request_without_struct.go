package main

import (
	"flag"
	"fmt"
)

//specified name: 命令行指定的名字，栗子: -get
//default value: -test 的默认值: https://ip:port
//usage: 这条指令的使用说明

var getRequest *string = flag.String("get", "https://ip:port", "发起一个GET请求")
var postRequest *string = flag.String("post", "https://ip:port", "发起一个POST请求")
var putRequest *string = flag.String("put", "https://ip:port", "发起一个PUT请求")
var deleteRequest *string = flag.String("delete", "https://ip:port", "发起一个DELETE请求")

//如果命令行解析error, 将输入所有定义的flag
func RequestWithoutStruct() {
	flag.Parse() //自定义完flag后，需要调用flag.Parse()进行解析：
	if *getRequest != "https://ip:port" {
		fmt.Println("向", *getRequest, "发起了一个GET请求")
	} else {
		fmt.Println("get的默认值为:https//ip:port")
	}

	if *postRequest != "https://ip:port" {
		fmt.Println("向", *postRequest, "发起了一个POST请求")
	}

	if *putRequest != "https://ip:port" {
		fmt.Println("向", *putRequest, "发起了一个PUT请求")
	}

	if *deleteRequest != "https://ip:port" {
		fmt.Println("向", *deleteRequest, "发起了一个DELETE请求")
	}
}
