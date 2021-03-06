package blog

import (
	"strings"

	"github.com/brainwu/brainblog/models"
	beego "gopkg.in/astaxie/beego.v1"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Get() {
	var article *models.Article = new(models.Article)
	articles := article.List(true)
	this.Data["articles"] = articles
	this.TplNames = "index.html"
}

func (this *IndexController) GetArticle() {
	sPath := this.Ctx.Request.URL.Path
	arr := strings.Split(sPath, "/")
	if len(arr) <= 0 {
		this.Redirect("/", 302)
		return
	}
	sTitle := arr[len(arr)-1]
	var a *models.Article = &models.Article{
		Title: sTitle,
	}
	if err := a.Read("Title"); err != nil {
		beego.Error("GetArticle read " + sTitle + " : " + err.Error())
		this.Redirect("/", 302)
		return
	} else {
		this.Data["title"] = a.Title
		this.Data["content"] = a.Content
		this.TplNames = "article.html"
	}
}
