package routers

import (
	"github.com/brainwu/brainblog/controllers/blog"
	"github.com/brainwu/brainblog/controllers/manage"

	beego "gopkg.in/astaxie/beego.v1"
)

func init() {
	ns := beego.NewNamespace("/manage",
		beego.NSRouter("/", &manage.ManageController{}, "*:Manage"),
		beego.NSNamespace("/article",
			beego.NSRouter("/", &manage.ArticleController{}),
			beego.NSRouter("/list", &manage.ArticleController{}, "*:List"),
		),
		beego.NSRouter("/login", &manage.LoginController{}, "*:Login"),
		beego.NSRouter("/logout", &manage.LoginController{}, "*:Logout"),
	)
	beego.AddNamespace(ns)
	beego.Router("/articles/*", &blog.IndexController{}, "get:GetArticle")
	beego.Router("/", &blog.IndexController{})
}
