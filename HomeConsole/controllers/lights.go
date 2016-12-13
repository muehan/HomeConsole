package controllers

import (
	"github.com/astaxie/beego"
)

type LightsController struct {
	beego.Controller
}

func (c *LightsController) Get() {

	c.TplName = "lights/index.html"
	c.Layout = "_layout.html"
}
