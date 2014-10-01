package manage

import (
	"strconv"
	"strings"

	beego "gopkg.in/astaxie/beego.v1"
	"github.com/brainwu/brainblog/models"
)

type baseController struct {
	beego.Controller
	actionName     string
	controllerName string
	options        map[string]string
}

func (base *baseController) Prepare() {
	base.options = models.GetOptions()
	base.Data["options"] = base.options
	base.Data["name"] = "Brain Wu"
	base.controllerName, base.actionName = base.GetControllerAndAction()
	base.auth()
}

func (base *baseController) auth() {
	if base.actionName == "Login" && base.controllerName == "LoginController" {
		return
	}
	if !base.IsCookieValid() {
		base.Ctx.Redirect(302, "/manage/login")
		base.StopRun()
	}
}

//GetUid get the Uid by the params been passed from client user.
func (base *baseController) GetUid() string {
	cookie, _ := base.Ctx.Request.Cookie("token")
	if cookie != nil {
		token := cookie.Value
		arr := strings.Split(token, "|")
		uid := arr[0]
		return uid
	} else {
		return ""
	}
}

//IsCookieValid verifies the cookie been passed by client.
//If cookie is in valid, IsCookieValid return true. Otherwise, return false.
func (base *baseController) IsCookieValid() bool {
	cookie, err := base.Ctx.Request.Cookie("token")
	if err != nil {
		return false
	}
	token := cookie.Value
	arr := strings.Split(token, "|")
	if len(arr) < 2 {
		return false
	}
	uid, reqMd5Pass := arr[0], arr[1]
	var user *models.User
	if iid, err := strconv.Atoi(uid); err != nil {
		return false
	} else {
		user = &models.User{Id: iid}
	}
	if err := user.Read("Id"); err != nil {
		return false
	}
	md5Pass := models.Md5([]byte(base.GetRemoteIp() + "|" + user.Password))
	if reqMd5Pass != md5Pass {
		return false
	}
	return true
}

//get client ip address.
func (base *baseController) GetRemoteIp() string {
	s := strings.Split(base.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
