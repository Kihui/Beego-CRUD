package routers

import (
	"beego-crud/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/crea", &controllers.Crea{})
}
