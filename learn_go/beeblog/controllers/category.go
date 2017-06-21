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
	beego.Alert(op)
	switch op {
	case "add":
		name := self.Input().Get("name")
		beego.Alert(name)
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		self.Redirect("/category", 302)
		return
	case "del":
		id := self.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		self.Redirect("/category", 302)
		return
	}
	self.TplName = "category.html"
	self.Data["IsCategory"] = true
	self.Data["IsLogin"] = checkAccount(self.Ctx)

	var err error
	self.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
