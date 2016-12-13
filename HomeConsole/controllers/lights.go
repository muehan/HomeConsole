package controllers

import (
	"HomeConsole/HomeConsole/services"

	"github.com/astaxie/beego"
)

type LightsController struct {
	beego.Controller
}

func (c *LightsController) Get() {

	c.Data["lights"] = services.GetLights()
	c.TplName = "lights/index.html"
	c.Layout = "_layout.html"
}
