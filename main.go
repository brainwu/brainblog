package main

import (
	_ "github.com/brainwu/brainblog/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static","views/static")
	beego.Run()
}
