package controllers

import "github.com/astaxie/beego"

type OutTempController struct {
	beego.Controller
}

func (controller *OutTempController) Get() {
	//controller.Data["lights"] = services.GetLights()
	controller.TplName = "outtemp/index.html"
	controller.Layout = "_layout.html"
}
