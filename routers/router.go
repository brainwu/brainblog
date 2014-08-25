package routers

import (
	"github.com/brainwu/brainblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
}
