package controllers

import (
	"HomeConsole/HomeConsole/helper"
	"HomeConsole/HomeConsole/models"
	"HomeConsole/HomeConsole/services"
	"fmt"

	"github.com/astaxie/beego"
)

type SettingsEditController struct {
	beego.Controller
}

func (controller *SettingsEditController) Get() {
	param := controller.Ctx.Input.Param(":id")
	id := helper.StringToInt(param)

	light := services.GetLight(id)

	controller.Data["light"] = light
	controller.TplName = "settings/edit.html"
	controller.Layout = "_layout.html"
}

func (controller *SettingsEditController) Post() {
	var post models.Post
	err := controller.ParseForm(&post)

	fmt.Println(post)

	if err != nil {
		fmt.Println("Error during parsing Form")
		fmt.Println(err)

		controller.Abort("500")
	}

	services.ChangeLight(post)

	controller.Redirect("/settings", 302)
}
