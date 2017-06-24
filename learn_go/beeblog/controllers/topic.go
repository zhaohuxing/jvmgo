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
	topics, err := models.GetAllTopics("", false)
	if err != nil {
		beego.Error(err)
	} else {
		self.Data["Topics"] = topics
	}
}

func (self *TopicController) Post() {

	if !checkAccount(self.Ctx) {
		self.Redirect("/login", 302)
		return
	}

	title := self.Input().Get("title")
	content := self.Input().Get("content")
	tid := self.Input().Get("tid")
	category := self.Input().Get("category")
	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}

	if err != nil {
		beego.Error(err)
	}

	self.Redirect("/topic", 302)
}

func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
}

func (self *TopicController) View() {
	self.TplName = "topic_view.html"

	tid := self.Ctx.Input.Param("0")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = topic
	self.Data["Tid"] = self.Ctx.Input.Param("0")

	replies, err1 := models.GetAllReplies(tid)
	if err1 != nil {
		beego.Error(err1)
		return
	}
	self.Data["Replies"] = replies
	self.Data["IsLogin"] = checkAccount(self.Ctx)
}

func (self *TopicController) Modify() {
	self.TplName = "topic_modify.html"
	tid := self.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}

	self.Data["Topic"] = topic
	self.Data["Tid"] = tid
}

func (self *TopicController) Delete() {
	if !checkAccount(self.Ctx) {
		self.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(self.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/", 302)
}
