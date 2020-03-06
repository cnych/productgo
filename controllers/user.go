package controllers

import (
	"fmt"
	"productgo/manager/user"

	"github.com/astaxie/beego"
)

type userController struct {
	beego.Controller
	userMgr user.Manager
}

func (u *userController) Prepare() {
	u.Controller.Prepare()
	u.userMgr = user.Mgr
}

type LoginController struct {
	userController
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

func (l *LoginController) Logout() {
	l.DelSession("current_user")
	l.Ctx.Redirect(302, "/")
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")
	user, err := l.userMgr.Get(username, password)
	if err != nil {
		l.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "登录失败",
		}
	} else {
		if user.Id > 0 {
			l.SetSession("current_user", user)
			//	匹配密码
			l.Data["json"] = map[string]interface{}{
				"code":    1,
				"message": "登录成功",
			}
		} else {
			l.Data["json"] = map[string]interface{}{
				"code":    0,
				"message": "登录失败",
			}
			fmt.Println(3)
		}
	}
	l.ServeJSON()
}

// 注册

type RegisterController struct {
	userController
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")

	if password != repassword {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "两次输入密码不一致~",
		}
		r.ServeJSON()
		return
	}

	// 注册之前先判断用户名是否已经存在
	if r.userMgr.IsExists(username) {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "用户名已存在",
		}
		r.ServeJSON()
		return
	}

	// 用户名不存在，添加到数据了
	_, err := r.userMgr.Create(username, password)
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
