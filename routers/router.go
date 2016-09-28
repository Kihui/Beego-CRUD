package routers

import (
	"beego-crud/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/update", &controllers.MainController{}, "*:Update")
	beego.Router("/delete", &controllers.MainController{}, "get:Delete;post:Add")	
	beego.Router("/", &controllers.MainController{}, "get:View;post:Add")
	
}
