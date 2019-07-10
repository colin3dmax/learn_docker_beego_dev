package main

import (
	"beego_01/models"
	_ "beego_01/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(models.User))
}

func main() {
	beego.Run()
}
