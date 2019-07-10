package main

import (
	"beego_00/models"
	_ "beego_00/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(models.User))
}

func main() {
	beego.Run()
}
