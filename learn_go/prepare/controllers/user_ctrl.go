package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"prepare/models"
	. "prepare/utils" //Result
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.html")
		t.Execute(w, nil)
	}

	if r.Method == "POST" {
		phoneNumber, password := getAccountMsg(r)
		w.Header().Set("Content-Type", "application/json")
		var result *Result
		if models.GetUser(phoneNumber, password) {
			// result TODO
			result = &Result{
				Code:    1,
				Message: "login successful.",
			}
		} else {
			result = &Result{
				Code:    0,
				Message: "login failed.",
			}
		}
		json, _ := json.Marshal(result)
		w.Write(json)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "POST" {
		phoneNumber, password := getAccountMsg(r)
		w.Header().Set("Content-Type", "application/json")
		var result *Result
		if models.InsertUser(phoneNumber, password) {
			// result TODO
			result = &Result{
				Code:    1,
				Message: "insert successful.",
			}
		} else {
			result = &Result{
				Code:    0,
				Message: "insert failed.",
			}
		}
		json, _ := json.Marshal(result)
		w.Write(json)
	}
}

func getAccountMsg(r *http.Request) (phoneNumber, password string) {
	r.ParseForm()
	phoneNumber = r.PostFormValue("phoneNumber")
	password = r.PostFormValue("password")
	return
}
