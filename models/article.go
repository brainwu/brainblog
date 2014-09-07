package models

import (
	"gopkg.in/astaxie/beego.v1/orm"
	"time"
	"strconv"
	"html/template"
)

type Article struct {
	Id         int
	User       *User `orm:"rel(fk)"`
	UserName   string
	Title      string
	Url        string
	Content    template.HTML
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

//Delete delete an article and delete tag_article‘s relationship in blog_tag_article table.
func (a *Article) Delete() error {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return err
	}
	if _, err := o.QueryM2M(a, "Tags").Clear(); err != nil {
		return err
	}
	if _, err := o.Delete(a); err != nil {
		o.Rollback();
		return err
	}
	return o.Commit()
}

//Read read an article by specified field without Article.Tags.
//Such as --> a := Article{Id:1} a.Read("Id").
//It's equivalent to “select * from blog_article where id = 1”.
func (a *Article) Read(fields ...string) error {
	o := orm.NewOrm()
	if err := o.Read(a, fields...); err != nil {
		return err
	}
	return nil
}

//GetTags get tags of article.
func (a *Article) GetTags() error {
	o := orm.NewOrm()
	rs := o.Raw("select tag_id from "+(&TagArticle{}).TableName()+" where article_id=?", a.Id)
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
	qm2 := o.QueryM2M(a, "Tags")
	qm2.Clear()
	//update blog_tag_article table
	if len(a.Tags) != 0 {
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
