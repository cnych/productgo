package routers

import (
	"github.com/astaxie/beego"
	"productgo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/products", &controllers.ProductController{})
	beego.Router("/products/:id([0-9]+)/vote/", &controllers.ProductVoteController{})
}
