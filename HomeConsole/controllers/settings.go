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
	var post models.Post

	err := controller.ParseForm(&post)
	fmt.Println(controller.Ctx.Request.Form)

	if err != nil {
		fmt.Println("Error during parsing Form")
		fmt.Println(err)

		controller.Abort("500")
	}

	// if lights == nil {
	// 	fmt.Println("No lights parse from Form")
	// 	controller.Abort("500")
	// }

	fmt.Println(post)
	// services.SetLights(lights)

	controller.Ctx.Redirect(201, "/settings")
}
