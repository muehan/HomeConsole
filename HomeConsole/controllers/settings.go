package controllers

import (
	"HomeConsole/HomeConsole/models"
	"HomeConsole/HomeConsole/services"

	"github.com/astaxie/beego"
)

type SettingsController struct {
	beego.Controller
}

func (controller *SettingsController) Get() {
	var config models.Config
	config.Lights = services.GetLights()

	controller.Data["config"] = config
	controller.TplName = "settings/index.html"
	controller.Layout = "_layout.html"
}
