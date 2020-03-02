package controllers

import (
	"fmt"
	"net/http"
	"productgo/manager/product"
	"productgo/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ProductController struct {
	beego.Controller
	proMgr product.Manager
}

func (pc *ProductController) Prepare() {
	pc.Controller.Prepare()
	pc.proMgr = product.Mgr
}

func (pc *ProductController) CreateProduct() {
	name := pc.GetString("name")
	link := pc.GetString("link")
	description := pc.GetString("description")
	user := pc.GetSession("current_user") // 当前登录用户

	if user != nil {
		_, err := pc.proMgr.Create(name, description, link, user.(*models.User))
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

type productVoteData struct {
	VoteCount int `json:"vote_count"`
}
type productVoteResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message,omitempty"`
	Data    *productVoteData `json:"data,omitempty"`
}

func (pv *ProductController) VoteProduct() {
	idStr := pv.Ctx.Input.Param(":id")
	fmt.Printf("idStr=%s\n", idStr)
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(pv.Ctx.ResponseWriter, fmt.Sprintf("invalid product id %s", idStr), http.StatusBadRequest)
		return
	}

	user := pv.GetSession("current_user") // 当前登录用户
	if user == nil {                      // 未登录
		pvr := productVoteResponse{
			Code:    0,
			Message: "当前用户未登录",
		}
		pv.Data["json"] = pvr
		pv.ServeJSON()
		return
	}

	product, err := pv.proMgr.Get(id64)
	if err != nil {
		http.Error(pv.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	// product、user 都获取到了，去执行点赞操作
	err = pv.proMgr.Vote(user.(*models.User), product)
	if err != nil {
		http.Error(pv.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
	} else {
		pvr := productVoteResponse{
			Code: 1,
			Data: &productVoteData{VoteCount: product.VoteCount},
		}
		pv.Data["json"] = pvr
		pv.ServeJSON()
	}

}
