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

func (c *MainController) View() {
	c.TplName = "index.tpl"
	c.Data["Form"] = &models.Usuario{}
	o := orm.NewOrm()
	o.Using("default")

	var usuarios []*models.Usuario
	num, err := o.QueryTable("usuario").All(&usuarios)

	if err != orm.ErrNoRows && num > 0 {
		c.Data["records"] = usuarios
	}
}

func (this *MainController) Add() {
	this.TplName = "index.tpl"
	o := orm.NewOrm()
	o.Using("default")
	
	u := models.Usuario{}
	if err := this.ParseForm(&u); err != nil {
		beego.Error("Error: ", err)
	} else {
		this.Data["Usuarios"] = u
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

