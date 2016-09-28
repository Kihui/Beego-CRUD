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

	var usuarios []*models.Usuario
	num, err := o.QueryTable("usuario").All(&usuarios)

	if err != orm.ErrNoRows && num > 0 {
		this.Data["records"] = usuarios
	}
}

func (this *MainController) Delete() {
	this.TplName = "index.tpl"
	
	id, err := this.GetInt("id")
	if err != nil {
		msg := fmt.Sprintf("Parámetro inválido: ", err)
		beego.Debug(msg)
	}

	o := orm.NewOrm()
	o.Using("default")

	if existe := o.QueryTable("usuario").Filter("Id", id).Exist(); existe {
		if num, err := o.Delete(&models.Usuario{Id: id}); err == nil {
			beego.Info("Usuario eliminado. ", num)
		} else {
			beego.Error("Usuario no pudo ser eliminado: ", err)
		}
	} else {
		beego.Info("El usuario no existe.")
	}
	
	var usuarios []*models.Usuario
	num, err := o.QueryTable("usuario").All(&usuarios)

	if err != orm.ErrNoRows && num > 0 {
		this.Data["records"] = usuarios
	}
}

func (this *MainController) GetUpdate() {
	
}

func (this *MainController) Update() {
	this.TplName = "update.tpl"
	this.Data["Form"] = &models.Usuario{}
	o := orm.NewOrm()
	o.Using("default")

	if id, err := this.GetInt("id"); err == nil {
		u := models.Usuario{Id: id}
		if o.Read(&u) != nil {
			beego.Error("El usuario no existe")
		} else {
			this.Data["Nom"] = u.Nombre
			this.Data["App"] = u.App
			this.Data["Apm"] = u.Apm
			this.Data["Correo"] = u.Correo
			this.Data["Edad"] = u.Edad
		}
		if this.Ctx.Input.Method() == "POST" {
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
					id, err := o.Update(&u)
					if err == nil {
						msg := fmt.Sprintf("Usuario actualizado: ", id)
						beego.Debug(msg)
					} else {
						msg := fmt.Sprintf("No se pudo actulizar usuario: ", err)
						beego.Debug(msg)
					}
				}
			}			
			this.Redirect("/", 302)
			return
		}
	} else {
		beego.Error("Parámetro inválido: ", err)
	}		
}
