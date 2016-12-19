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
	var config models.Config

	err := controller.ParseForm(&config)
	fmt.Println(controller.Ctx.Request.Form)

	if err != nil {
		fmt.Println("Error during parsing Form")
		fmt.Println(err)

		controller.Abort("500")
	}

	if config.Lights == nil {
		fmt.Println("No lights parse from Form")
		controller.Abort("500")
	}

	services.SetLights(&config.Lights)

	controller.Ctx.Redirect(201, "/settings")
}
