package routers

import (
	"github.com/brainwu/brainblog/controllers/blog"
	"github.com/brainwu/brainblog/controllers/manage"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &blog.IndexController{})

	ns := beego.NewNamespace("/manage",
		beego.NSRouter("/", &manage.ManageController{}, "*:Manage"),
		beego.NSRouter("/login", &manage.LoginController{}, "*:Login"),
	)
	beego.AddNamespace(ns)
}
