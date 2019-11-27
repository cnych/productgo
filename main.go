package main

import (
	"github.com/astaxie/beego"
	_ "productgo/routers"
	"productgo/utils"
)

func main() {
	if err := utils.InitMySQL(); err != nil {
		panic(err)
	}
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
