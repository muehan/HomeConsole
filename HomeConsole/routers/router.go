package routers

import (
	"HomeConsole/HomeConsole/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/lights", &controllers.LightsController{})
}
