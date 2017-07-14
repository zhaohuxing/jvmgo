package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"prepare/models"
	. "prepare/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.html")
		t.Execute(w, nil)
	}

	if r.Method == "POST" {
		// GetUser
		r.ParseForm()
		//		phoneNumber := r.Form["phoneNumber"]
		//		password := r.Form["password"]
		phoneNumber := r.PostFormValue("phoneNumber")
		password := r.PostFormValue("password")
		if models.InsertUser(phoneNumber, password) {
			// 返回给json: code : 1, msg: insert success
			//设置文本类型
			w.Header().Set("Content-Type", "application/json")
			result := &Result{
				Code:    1,
				Message: "响应成功",
			}
			json, _ := json.Marshal(result)
			w.Write(json)
		}
	}
}
