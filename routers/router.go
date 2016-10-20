package routers

import (
	"beego-crud/controllers"
	"github.com/astaxie/beego"
)

// Asignación de metodos de un controlador para cada método HTTP de un url
func init() {
	beego.Router("/update", &controllers.MainController{}, "*:Update")
	beego.Router("/delete", &controllers.MainController{}, "*:Delete")	
	beego.Router("/", &controllers.MainController{}, "get:View;post:Add")	
}
