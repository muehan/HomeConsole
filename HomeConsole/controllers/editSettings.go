package controllers

import (
	"HomeConsole/HomeConsole/models"
	"HomeConsole/HomeConsole/services"

	"fmt"

	"HomeConsole/HomeConsole/helper"

	"github.com/astaxie/beego"
)

type EditSettingsController struct {
	beego.Controller
}

func (controller *EditSettingsController) Get() {
	param := controller.Ctx.Input.Param(":id")
	id := helper.StringToInt(param)

	light := services.GetLight(id)

	controller.Data["light"] = light
	controller.TplName = "settings/edit.html"
	controller.Layout = "_layout.html"
}

func (controller *EditSettingsController) Post() {
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
