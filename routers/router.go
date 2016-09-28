package routers

import (
	"beego-crud/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/delete", &controllers.MainController{}, "get:Delete")
	beego.Router("/", &controllers.MainController{}, "get:View;post:Add")
	
}
