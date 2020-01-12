package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"productgo/models"
)

type ProductController struct {
	beego.Controller
}

func( pc *ProductController) Post() {
	name := pc.GetString("name")
	link := pc.GetString("link")
	description := pc.GetString("description")
	user := pc.GetSession("current_user")  // 当前登录用户

	if user != nil {
		product := &models.Product{
			Name: name,
			Link: link,
			Description: description,
			User: user.(*models.User),
		}
		o := orm.NewOrm()
		_, err := o.Insert(product)
		if err != nil {
			pc.Data["json"] = map[string]interface{}{
				"code":    0,
				"message": "添加失败",
			}
		} else {
			pc.Data["json"] = map[string]interface{}{
				"code":    1,
				"message": "添加成功",
			}
		}
	} else {
		pc.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "当前用户未登录",
		}
	}
	pc.ServeJSON()
}