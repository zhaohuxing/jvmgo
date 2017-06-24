package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (self *ReplyController) Add() {
	tid := self.Input().Get("tid")
	nickname := self.Input().Get("nickname")
	content := self.Input().Get("content")
	err := models.AddReply(tid, nickname, content)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic/view/"+tid, 302)
}

func (self *ReplyController) Delete() {
	if !checkAccount(self.Ctx) {
		return
	}
	tid := self.Input().Get("tid")
	rid := self.Input().Get("rid")

	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic/view/"+tid, 302)
}
