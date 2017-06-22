package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
}
