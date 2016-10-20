package controllers

import (
	"github.com/astaxie/beego"
	models "beego-crud/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"fmt"
)

// Definición de un controlador
type MainController struct {
	beego.Controller
}

// Método que muestra todos los usuarios
func (c *MainController) View() {
	// Nombre del template
	c.TplName = "index.tpl"
	// Se le asigna un modelo al formulario de la vista
	c.Data["Form"] = &models.Usuario{}
	// Se crea una conexión a la base
	o := orm.NewOrm()
	// Se indica que se usará la base por default
	o.Using("default")

	// Arreglo donde se guardarán los usuarios
	var usuarios []*models.Usuario
	// Se guardan los usuarios de la base en el arreglo
	num, err := o.QueryTable("usuario").All(&usuarios)

	if err != orm.ErrNoRows && num > 0 {
		// Se pasa el arreglo al contexto de la vista
		c.Data["records"] = usuarios
	}
}

// Método que permite registrar un nuevo usuario
func (this *MainController) Add() {
	this.TplName = "index.tpl"
	o := orm.NewOrm()
	o.Using("default")

	// Se crea un nuevo usuario
	u := models.Usuario{}
	// Se guardan los datos del formulario en el nuevo usuario
	if err := this.ParseForm(&u); err != nil {
		// Si hay errores en el formulario
		beego.Error("Error: ", err)
	} else {
		this.Data["Usuarios"] = u
		// Se validan los datos del usuario
		valid := validation.Validation{}
		isValid, _ := valid.Valid(u)
		if !isValid {
			this.Data["Errors"] = valid.ErrorsMap
			beego.Error("Datos inválidos")
		} else {
			// Se inserta el usuario nuevo en la base de datos
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

// Método que permite eliminar un usuario
func (this *MainController) Delete() {
	this.TplName = "delete.tpl"
	// Se obtiene el parámetro id del url
	id, err := this.GetInt("id")
	if err != nil {
		msg := fmt.Sprintf("Parámetro inválido: ", err)
		beego.Debug(msg)
	}

	o := orm.NewOrm()
	o.Using("default")

	// Se obtiene el usuario con el id del parámetro
	if existe := o.QueryTable("usuario").Filter("Id", id).Exist(); existe {
		// Se elimina el usuario de la base de datos
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
	this.Redirect("/",302)
}

// Método que permite modificar los dato de un usuario
func (this *MainController) Update() {
	this.TplName = "update.tpl"
	this.Data["Form"] = &models.Usuario{}
	o := orm.NewOrm()
	o.Using("default")

	// Se obtiene el parámetro id del url
	if id, err := this.GetInt("id"); err == nil {
		// Se obtiene el usuario con el id del parámetro
		u := models.Usuario{Id: id}
		if o.Read(&u) != nil {
			beego.Error("El usuario no existe")
		} else {
			// Se pasan los datos del usuario al contexto de la vista
			this.Data["Nom"] = u.Nombre
			this.Data["App"] = u.App
			this.Data["Apm"] = u.Apm
			this.Data["Correo"] = u.Correo
			this.Data["Edad"] = u.Edad
		}
		// Si se efectúa el método POST en la vista
		if this.Ctx.Input.Method() == "POST" {
			// Se guardan los datos del formulario en el usuario
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
					// Se actualiza el usuario con los nuevos datos en la base
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
