package main

import (
	_ "beego-crud/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	models "beego-crud/models"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:magdario@/beego?charset=utf8")
    orm.RegisterModel(new(models.Usuario))
    orm.RunCommand()
}

func main() {
	beego.Run()
}

