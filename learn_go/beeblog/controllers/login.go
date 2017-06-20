package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	isExit := self.Input().Get("exit") == "true"
	if isExit {
		self.Ctx.SetCookie("uname", "", -1, "/")
		self.Ctx.SetCookie("pwd", "", -1, "/")
		self.Redirect("/", 301)
		return
	}
	self.TplName = "login.html"
}

func (self *LoginController) Post() {
	uname := self.Input().Get("uname")
	pwd := self.Input().Get("pwd")
	autoLogin := self.Input().Get("autoLogin") == "on"
	if check(uname, pwd) {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<32 - 1
		}
		self.Ctx.SetCookie("uname", uname, maxAge, "/")
		self.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	self.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return check(uname, pwd)
}

func check(uname, pwd string) bool {
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		return true
	}
	return false
}
