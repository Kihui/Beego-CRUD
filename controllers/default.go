package controllers

import (
	"github.com/astaxie/beego"
	models "beego-crud/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Crea struct {
	beego.Controller
}

func (c *Crea) Get() {
	c.TplName = "crea.tpl"
	c.Data["Form"] = &models.Usuario{}
}

func (this *Crea) Post() {
	o := orm.NewOrm()
	o.Using("default")
	
	u := models.Usuario{}
	if err := this.ParseForm(&u); err != nil {
		beego.Error("Error: ", err)
	} else {
		manage.Data["Usuarios"] = u
		valid := validation.Validation{}
		isValid, _ := valid.Valid(u)
		if !isValid {
			this.Data["Errors"] = valid.ErrorsMap
			beego.Error("Datos inválidos")
		} else {	
			id, err := o.Insert(&u)
			if err == nil {
				msg := fmt.Sprintf("Nuevo usuario con id: ", id)
				beego.Debug(msg)
			} else {
				msg := fmt.Sprintf("No se pudo registrar usuario: ", err)
				beego.Debug(msg)
			}
		}
	}
}
