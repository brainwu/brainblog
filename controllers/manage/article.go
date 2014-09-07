package manage

import (
	"github.com/brainwu/brainblog/models"
	"time"
	"strconv"
	"strings"
	"os"
	"html/template"
)

type ArticleController struct {
	baseController
}

//名字相同 不能插入 做个提示
//标题内容不为空 做个提示
//用户id要是正确的 做个提示
func (ac *ArticleController) Post() {
	title := ac.GetString("title")
	content := ac.GetString("content")
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == ""  {
		ac.Ctx.Redirect(304, "")
		return
	}
	t := time.Now()
	uid, err := strconv.Atoi(ac.GetUid())
	if err != nil {
		ac.Ctx.WriteString("无效的用户操作")
		return;
	}
	var u models.User = models.User{Id:uid}
	if err := u.Read("Id"); err != nil {
		ac.Ctx.WriteString("无效的用户操作")
		return
	}

	// sUrl such as /2009/11/10/title
	arr := strings.Split(t.String(), " ")
	sDate := arr[0]
	sDate = strings.Replace(sDate, "-", "/", -1)
	sUrl := sDate + string(os.PathSeparator) + title
	var a *models.Article = &models.Article{
		Title : title,
		Content: template.HTML(content),
		CreateTime:t,
		UpdateTime:t,
		User:&u,
		UserName:u.Name,
		Url:sUrl,
	}
	if err := a.Insert(); err == nil {
		ac.Ctx.Redirect(302, "/articles/" + sUrl)
	} else {
		ac.Ctx.Redirect(304, "")
	}
}


func (ac *ArticleController) List() {
	ac.Layout = "manage/manage.html"
	ac.TplNames="manage/article_list.html"
}

