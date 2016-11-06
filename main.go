package main

import (
	_ "beego-crud/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	models "beego-crud/models"
)

func init() {
	// Se registra el driver de la base de datos
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// Se registra la base de dator
	orm.RegisterDataBase("default", "mysql", "root:magdario@/beego?charset=utf8")
	// Se registra el modelo de usuario
	orm.RegisterModel(new(models.Usuario))
	orm.RunCommand()
}

func main() {
	// Se corre la aplicación
	beego.Run()
}

