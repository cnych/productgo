package main

import (
	"encoding/gob"
	"fmt"
	"productgo/models"
	_ "productgo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

var (
	mysqluser  = beego.AppConfig.String("mysqluser")
	mysqlpass  = beego.AppConfig.String("mysqlpass")
	mysqlhost  = beego.AppConfig.String("mysqlhost")
	mysqldb    = beego.AppConfig.String("mysqldb")
	dataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mysqluser, mysqlpass, mysqlhost, mysqldb)
)

func init() {
	gob.Register(&models.User{})
}

func main() {
	// set default database
	if err := orm.RegisterDataBase("default", "mysql", dataSource, 30); err != nil {
		panic(err)
	}
	if beego.AppConfig.String("runmode") == beego.DEV {
		orm.Debug = true
	}
	// register model
	orm.RegisterModel(new(models.User), new(models.Product), new(models.ProductVote))
	// create table
	if err := orm.RunSyncdb("default", false, true); err != nil {
		panic(err)
	}

	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = dataSource
	beego.Run()
}
