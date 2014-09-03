package main

import (
	"fmt"
	"github.com/brainwu/brainblog/models"
)

func main() {
	getArticles()
}

func getArticles() {
	var t *models.Tag = &models.Tag{Id:1}
	t.Read("Id")
	t.GetArticles()
	fmt.Println(len(t.Articles))
}
