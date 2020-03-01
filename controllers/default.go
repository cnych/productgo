package controllers

import (
	"net/http"
	"productgo/models"
	"productgo/utils"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type ProductDateItem struct {
	Date     string
	Products []models.Product
	Uid      int
}

func (c *MainController) Get() {
	currentUser := c.GetSession("current_user")
	uid := 0
	var user *models.User
	if currentUser != nil {
		user = currentUser.(*models.User)
		uid = user.Id
	}

	lastDateStr := c.GetString("last_dt")
	if lastDateStr == "" { // 首页
		var items []ProductDateItem

		deltas := []int{0, 1, 2}
		for _, delta := range deltas {
			// 今天
			now := time.Now()
			todayDate := utils.DateDelta(now, -1*delta)
			// 明天
			tomorrowDate := utils.DateDelta(now, -1*delta+1)

			//	2019-12-31 03:26:06 containers 2019-12-31
			item, err := getProductsByDate(todayDate, tomorrowDate, uid)
			if err != nil {
				http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
				return
			} else {
				items = append(items, *item)
			}
		}

		c.TplName = "index.html"
		c.Data["Items"] = items
		c.Data["User"] = user
	} else { // 前一天的数据
		lastDate, err := utils.Str2Date(lastDateStr)
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
			return
		}
		todayDate := utils.DateDelta(lastDate, -1)
		// 2020-02-04
		// 获取2020-02-03的数据
		// 02-04 (lastDate)
		// 02-03 (>=02-03，<02-04)
		// 02-02
		item, err := getProductsByDate(todayDate, lastDate, uid)
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
			return

		} else {
			c.TplName = "partials/products.html"
			c.Data["Date"] = item.Date
			c.Data["Products"] = item.Products
			c.Data["Uid"] = item.Uid
		}
	}

}

func getProductsByDate(todayDate, tomorrowDate time.Time, uid int) (*ProductDateItem, error) {
	o := utils.GetOrmer()
	// 获取 QuerySeter 对象，product 为表名
	qs := o.QueryTable("product")

	var products []models.Product
	_, err := qs.RelatedSel("user").Filter("created__gte", todayDate).Filter("created__lt", tomorrowDate).
		OrderBy("-created").All(&products)

	if err != nil {
		return nil, err
	}
	item := ProductDateItem{
		Date:     utils.DateFormat(todayDate),
		Products: products,
		Uid:      uid,
	}
	return &item, nil
}
