package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (self *CategoryController) Get() {
	op := self.Input().Get("op")
	switch op {
	case "add":
		name := self.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		self.Redirect("/category", 301)
		return
	case "del":
		id := self.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}
	self.TplName = "category.html"
	self.Data["IsCategory"] = true

	var err error
	self.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
