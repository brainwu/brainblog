package controllers

import (
	"github.com/astaxie/beego"
	"github.com/brainwu/brainblog/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	var article *models.Article = new(models.Article)
	articles := article.List()
	this.Data["articles"] = articles
	this.TplNames = "index.html"
}
