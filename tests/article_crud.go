package main

import (
	"fmt"
	"github.com/brainwu/brainblog/models"
	"time"
)

func main() {
//	Insert()
//	Update()
	Read()
}

func Insert() {
	var tags []*models.Tag = make([]*models.Tag, 1)
	var tag *models.Tag = &models.Tag{Id: 1}
	fmt.Println(tag.Read("id"))
	tags[0] = tag
	fmt.Println(tag)
	var user *models.User = &models.User{Name: "brainwu"}
	user.QueryByName()
	var article *models.Article = &models.Article{
		User:       user,
		UserName:   user.Name,
		Title:      "mac pro",
		Content:    "mac pro",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Tags:       tags,
	}
	fmt.Println(article.Insert())
}

func Update() {
	var tags []*models.Tag = make([]*models.Tag, 2)
	tags[0] = &models.Tag{Id:1}
	tags[1] = &models.Tag{Id:2}
	var a *models.Article = &models.Article{Id: 12}
	fmt.Println(a.Read("Id"))
	a.Tags = tags
	a.Update()
}


func Read() {
	var a *models.Article = &models.Article{Id: 12}
	fmt.Println(a.Read("Id"))
	fmt.Println(a.GetTags())
	fmt.Println(a.Tags[0])
	fmt.Println(a.Tags[1])
}
