package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "productgo/routers"
	"productgo/models"
	"productgo/utils"
)

func main() {
	// set default database
	if err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/productgo?charset=utf8", 30); err != nil {
		panic(err)
	}
	orm.Debug = true
	// register model
	orm.RegisterModel(new(models.User), new(models.Product))
	// create table
	if err := orm.RunSyncdb("default", false, true); err != nil {
		panic(err)
	}

	if err := utils.InitDB(); err != nil {
		panic(err)
	}

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
