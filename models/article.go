package models

import (
	"gopkg.in/astaxie/beego.v1/orm"
	"time"
	"strconv"
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
	Tags       []*Tag `orm:"rel(m2m);rel_through(github.com/brainwu/brainblog/models.TagArticle)"` // ManyToMany relation
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
	o := orm.NewOrm()
	if err := o.Read(a, fields...); err != nil {
		return err
	}
	return nil
}

//将文章所对应的tag列表注入article中
func (a *Article)GetTags() error{
	o := orm.NewOrm()
	rs := o.Raw("select tag_id from blog_tag_article where article_id=?", a.Id)
	var maps []orm.Params
	if _, err := rs.Values(&maps); err != nil {
		return err
	}
	var tags []*Tag = make([]*Tag, len(maps))
	for i, v := range maps {
		if id, err := strconv.Atoi(v["tag_id"].(string)); err != nil {
			continue
		} else {
			tags[i] = &Tag{Id:id}
			tags[i].Read("Id")
		}
	}
	a.Tags = tags
	return nil
}

//如果想要更新tags那么必须为tags指定值
func (a *Article) Update(fields ...string) error {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return err
	}
	if _, err := o.Update(a, fields...); err != nil {
		o.Rollback()
		return err
	}

	//更新中间表
	if len(a.Tags) != 0 {
		qm2 := o.QueryM2M(a, "Tags")
		qm2.Clear()
		qm2.Add(a.Tags)
	}
	return o.Commit()
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
