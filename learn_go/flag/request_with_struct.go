package main

import (
	"flag"
	"fmt"
)

type RequestType struct {
	getRequest    string
	postRequest   string
	putRequest    string
	deleteRequest string
}

func RequestWithStruct() {
	req := &RequestType{}
	flag.StringVar(&req.getRequest, "get1", "https://ip:port", "发起一个GET请求")
	flag.StringVar(&req.postRequest, "post1", "https://ip:port", "发起一个POST请求")
	flag.StringVar(&req.putRequest, "put1", "https://ip:port", "发起一个PUT请求")
	flag.StringVar(&req.deleteRequest, "delete1", "https://ip:port", "发起一个DELETE请求")

	//解析flag
	flag.Parse()

	if req.getRequest != "https://ip:port" {
		fmt.Println("向", req.getRequest, "发起了一个GET请求")
	}
	if req.postRequest != "https://ip:port" {
		fmt.Println("向", req.postRequest, "发起了一个POST请求")
	}
	if req.putRequest != "https://ip:port" {
		fmt.Println("向", req.putRequest, "发起了一个PUT请求")
	}
	if req.deleteRequest != "https://ip:port" {
		fmt.Println("向", req.deleteRequest, "发起了一个DELETE请求")
	}

}
