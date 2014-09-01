package main

import (
	_ "github.com/brainwu/brainblog/routers"

	beego "gopkg.in/astaxie/beego.v1"
)

func main() {
	beego.SetStaticPath("/static", "views/static")
	beego.Run()
}
