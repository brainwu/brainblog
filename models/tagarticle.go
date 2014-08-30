package models

import (
)

type TagArticle struct {
	Id int
	Tag *Tag `orm:"rel(fk)"`
	Article *Article `orm:"rel(fk)"`
}

func(ta *TagArticle) TableName() string {
	return TableName("tag_article")
}

