package controllers

import (
	"HomeConsole/HomeConsole/models"
	"HomeConsole/HomeConsole/services"
	"strconv"

	"fmt"

	"github.com/astaxie/beego"
)

type EditSettingsController struct {
	beego.Controller
}

func (controller *EditSettingsController) Get() {
	param := controller.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("Error during parsing")
	}

	lights := services.GetLights()

	var light models.Light

	for _, l := range lights {
		if l.ID == id {
			light = l
		}
	}

	controller.Data["light"] = light
	controller.TplName = "settings/edit.html"
	controller.Layout = "_layout.html"
}
