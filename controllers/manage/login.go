package manage

import (
	"net/http"
	"strconv"
	"time"

	"github.com/brainwu/brainblog/models"
)

type LoginController struct {
	baseController
}

func (l *LoginController) Login() {
	l.TplNames = "manage/login.html"
	//if cookie is valid, redirect to /manage
	if l.IsCookieValid() {
		l.Ctx.Redirect(302, "/manage")
		return
	}
	//just allow POST method to login
	if l.Ctx.Request.Method != "POST" {
		return
	}

	var user models.User = models.User{
		Account : l.GetString("account"),
		Password :  l.GetString("password"),
	}

	if err := user.Read("account", "password"); err == nil {
		var cip string = l.GetRemoteIp()
		var token string = strconv.Itoa(user.Id) + "|" + models.Md5([]byte(cip + "|" + user.Password))
		//http.SetCookie(this.ResponseWriter,c)
		//this.ResponseWriter.Header().Add("Set-Cookie", c.String())
		var cookie http.Cookie
		if l.GetString("remember") == "yes" {
			cookie = http.Cookie{
				Name:    "token",
				Value:   token,
				Expires: time.Now().AddDate(0, 0, 7),
			}
		} else {
			cookie = http.Cookie{
				Name:    "token",
				Value:   token,
			}
		}
		l.Ctx.ResponseWriter.Header().Add("Set-Cookie", cookie.String())
		l.Ctx.Redirect(302, "/manage")
	} else {
		l.Data["errMsg"] = "密码错误或者账号未注册"
	}
}
