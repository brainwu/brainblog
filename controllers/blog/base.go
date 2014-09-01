package blog

import (
	beego "gopkg.in/astaxie/beego.v1"
	"github.com/brainwu/brainblog/models"
)

type baseController struct {
	beego.Controller
	options        map[string]string
}


func (base *baseController) Prepare() {
	base.options = models.GetOptions()
	base.Data["options"] = base.options
}
