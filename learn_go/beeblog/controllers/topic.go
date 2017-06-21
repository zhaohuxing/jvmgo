package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (self *TopicController) Get() {
	self.Data["IsToptic"] = true
	self.TplName = "topic.html"
	self.Data["IsLogin"] = checkAccount(self.Ctx)
}

func (self *TopicController) Post() {

	if !checkAccount(self.Ctx) {
		self.Redirect("/login", 302)
	}

	title := self.Input().Get("title")
	content := self.Input().Get("content")

	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	self.Redirect("/topic", 302)
}
func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
}
