package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (controller *MainController) Get() {
	controller.TplName = "index.html"
	controller.Layout = "_layout.html"
}
