package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         int
	User       *User `orm:"rel(fk)"`
	UserName   string
	Title      string
	Url        string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
	ReplyNum   int
	ClickNum   int
	TagsStr    string
	Tags 	   []*Tag `orm:"rel(m2m);rel_through(github.com/brainwu/brainblog/models.TagArticle)"` // ManyToMany relation
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


func (a *Article) Insert() error {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return err
	}
	if _, err := o.Insert(a); err != nil {
		o.Rollback()
		return err
	}
	if len(a.Tags) != 0 {
		if _, err := o.QueryM2M(a, "Tags").Add(a.Tags); err != nil {
			o.Rollback()
			return err
		}
	}
	return o.Commit()
}
