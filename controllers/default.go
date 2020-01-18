package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"productgo/models"
	"productgo/utils"
)

type MainController struct {
	beego.Controller
}

type ProductDateItem struct {
	Date string
	Products []models.Product
}

func (c *MainController) Get() {
	//c.Data["Website"] = "youdianzhishi.com"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"

	uid := c.GetSession("login_uid")
	if uid != nil {
		c.Data["Uid"] = uid
	}

	var items []ProductDateItem

	o := orm.NewOrm()
	// 获取 QuerySeter 对象，product 为表名
	qs := o.QueryTable("product")

	deltas := []int{0, 1, 2}
	for _, delta := range deltas {
		// 今天
		todayDate := utils.DateDelta(-1 * delta)
		// 明天
		tomorrowDate := utils.DateDelta(-1 * delta + 1)

	//	2019-12-31 03:26:06 containers 2019-12-31
		var products []models.Product
		num , err := qs.Filter("created__gte", todayDate).Filter("created__lt", tomorrowDate).
			OrderBy("-created").All(&products)
		if err != nil {
			fmt.Printf("Returned Rows Num: %d, Err: %s\n", num, err)
		} else {
			item := ProductDateItem {
				Date: utils.DateFormat(todayDate),
				Products: products,
			}
			items = append(items, item)
		}
	}

	fmt.Println(items)

	c.Data["Items"] = items

}
