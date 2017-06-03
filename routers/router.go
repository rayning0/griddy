package routers

import (
	"griddy/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get,post:Get")
	beego.Router("/prices", &controllers.MainController{}, "post:View")
	beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
}
