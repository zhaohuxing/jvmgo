package main

import (
	"log"
	"net/http"
	ctrl "prepare/controllers"
)

func main() {
	// 设置静态文件
	log.Println("本地http服务开启,端口:9000")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", ctrl.Index)
	http.HandleFunc("/login", ctrl.Login)
	http.HandleFunc("/register", ctrl.Register)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
