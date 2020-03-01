package main

import (
	"encoding/gob"
	"productgo/models"
	_ "productgo/routers"
	"productgo/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

var (
	dataSource = "root:root@tcp(127.0.0.1:3306)/productgo?charset=utf8"
)

func init() {
	gob.Register(&models.User{})
}

func main() {
	// set default database
	if err := orm.RegisterDataBase("default", "mysql", dataSource, 30); err != nil {
		panic(err)
	}
	orm.Debug = true
	// register model
	orm.RegisterModel(new(models.User), new(models.Product), new(models.ProductVote))
	// create table
	if err := orm.RunSyncdb("default", false, true); err != nil {
		panic(err)
	}

	if err := utils.InitDB(); err != nil {
		panic(err)
	}

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = dataSource
	beego.Run()
}
