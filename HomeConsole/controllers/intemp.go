package controllers

import "github.com/astaxie/beego"

type InTempController struct {
	beego.Controller
}

func (controller *InTempController) Get() {
	//controller.Data["lights"] = services.GetLights()
	controller.TplName = "intemp/index.html"
	controller.Layout = "_layout.html"
}
