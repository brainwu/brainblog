package main

import (
	_ "brainblog/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static","views/static")
	beego.Run()
}
