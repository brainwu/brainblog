package blog

import (
	"github.com/brainwu/brainblog/models"
	beego "gopkg.in/astaxie/beego.v1"
)

type baseController struct {
	beego.Controller
	actionName     string
	controllerName string
	options        map[string]string
}

func (base *baseController) Prepare() {
	base.controllerName, base.actionName = base.GetControllerAndAction()
	base.options = models.GetOptions()
	base.Data["options"] = base.options
}
