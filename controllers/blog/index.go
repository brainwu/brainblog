package blog

import (
	"github.com/brainwu/brainblog/models"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Get() {
	var article *models.Article = new(models.Article)
	articles := article.List()
	this.Data["articles"] = articles
	this.TplNames = "index.html"
}

