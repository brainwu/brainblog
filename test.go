package main

import (
	"github.com/brainwu/brainblog/models"
	"time"
	"fmt"
)

func main() {
	var tags []*models.Tag = make([]*models.Tag, 1)
	var tag *models.Tag = &models.Tag{Id:1}
	fmt.Println(tag.Read("id"))
	tags[0] = tag
	fmt.Println(tag)
	var user *models.User = &models.User{Name:"brainwu"}
	user.QueryByName()
	var article *models.Article = &models.Article {
		User:user,
		UserName:user.Name,
		Title:"mac pro",
		Content:"mac pro",
		CreateTime:time.Now(),
		UpdateTime:time.Now(),
		Tags : tags,
	}
	fmt.Println(article.Insert())
}
