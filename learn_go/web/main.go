package main

import (
	"jvmgo/learn_go/web/handler"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", handler.PhotoHandler())
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
