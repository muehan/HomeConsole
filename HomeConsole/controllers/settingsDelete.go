package controllers

import (
	"HomeConsole/HomeConsole/helper"
	"HomeConsole/HomeConsole/models"
	"HomeConsole/HomeConsole/services"
	"fmt"

	"github.com/astaxie/beego"
)

type SettingsDeleteController struct {
	beego.Controller
}

func (controller *SettingsDeleteController) Get() {
	param := controller.Ctx.Input.Param(":id")
	id := helper.StringToInt(param)

	light := services.GetLight(id)

	controller.Data["light"] = light
	controller.TplName = "settings/delete.html"
	controller.Layout = "_layout.html"
}

func (controller *SettingsDeleteController) Post() {
	var post models.Post
	err := controller.ParseForm(&post)

	fmt.Println(post)

	if err != nil {
		fmt.Println("Error during parsing Form")
		fmt.Println(err)

		controller.Abort("500")
	}

	services.DeleteLight(helper.StringToInt(post.ID))

	controller.Redirect("/settings", 302)
}
