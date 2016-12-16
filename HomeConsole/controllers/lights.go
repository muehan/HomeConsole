package controllers

import (
	"HomeConsole/HomeConsole/services"

	"github.com/astaxie/beego"
)

type LightsController struct {
	beego.Controller
}

func (controller *LightsController) Get() {

	controller.Data["lights"] = services.GetLights()
	controller.TplName = "lights/index.html"
	controller.Layout = "_layout.html"	
}
