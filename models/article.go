package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id int
	UserId int
	UserName string
	Title string
	Url string
	Content string
	CreateTime time.Time
	UpdateTime time.Time
	ReplyNum int
	ClickNum int
	Tags string
}

func (a *Article) TableName() string {
	return TableName("article")
}

func (a *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(a)
}

func (a *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(a); err != nil {
		return err
	}
	return nil
}

func (a *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
func (a *Article) List() []Article {
	var articles []Article
	orm.NewOrm().QueryTable(a).All(&articles)
	return articles
}
