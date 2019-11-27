package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "youdianzhishi.com"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	uid := c.GetSession("login_uid")
	if uid != nil {
		c.Data["Uid"] = uid
	}

}
