package routers

import (
	"HomeConsole/HomeConsole/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/lights", &controllers.LightsController{})
	beego.Router("/outtemp", &controllers.OutTempController{})
	beego.Router("/intemp", &controllers.InTempController{})
	beego.Router("/settings", &controllers.SettingsController{})
	beego.Router("/settings/edit/:id", &controllers.EditSettingsController{})
}
