package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"productgo/models"
	"productgo/utils"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")

	fmt.Println(username, password)

	user := models.QueryUserWithParam(username, utils.MD5(password))
	if user.Id > 0 {
		//	匹配密码
		l.Data["json"] = map[string]interface{}{
			"code":    1,
			"message": "登录成功",
		}
		l.SetSession("current_user", user)
	} else {
		l.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "登录失败",
		}
	}
	l.ServeJSON()
}

// 注册

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")

	fmt.Println(username, password, repassword)

	if password != repassword {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "两次输入密码不一致~",
		}
		r.ServeJSON()
		return
	}

	// 注册之前先判断用户名是否已经存在
	id := models.QueryUserWithUsername(username)
	if id > 0 {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "用户名已存在",
		}
		r.ServeJSON()
		return
	}

	// 用户名不存在，添加到数据了
	user := models.User{0, username, utils.MD5(password), time.Now()}
	_, err := models.InsertUser(&user)
	if err != nil {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "注册失败",
		}
	} else {
		r.Data["json"] = map[string]interface{}{
			"code":    1,
			"message": "注册成功",
		}
	}
	r.ServeJSON()
}
