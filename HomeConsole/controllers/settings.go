package controllers

import (
	"HomeConsole/HomeConsole/models"
	"HomeConsole/HomeConsole/services"

	"fmt"

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

func (controller *SettingsController) Post() {
	lights := new([]models.Light)

	err := controller.ParseForm(&lights)
	fmt.Println(controller.Ctx.Request.Form)

	if err != nil {
		fmt.Println("Error during parsing Form")
		fmt.Println(err)
	}

	services.SetLights(lights)

	controller.Ctx.Redirect(201, "/settings")
}
