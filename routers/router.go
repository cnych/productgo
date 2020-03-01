package routers

import (
	"productgo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/products", &controllers.ProductController{})
	beego.Router("/products/:id([0-9]+)/vote/", &controllers.ProductVoteController{})
}
