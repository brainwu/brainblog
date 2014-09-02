package models

import "strconv"

import "gopkg.in/astaxie/beego.v1/orm"

import (
)

type Tag struct {
	Id int
	Name string
	Count int
	Articles []*Article `orm:"reverse(many);"`
}


func (t *Tag) TableName() string {
	return TableName("tag")
}


func (t *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Delete() error {
	if _, err := orm.NewOrm().Delete(t); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(t, fields...); err != nil {
		return err
	}
	return nil
}

//获取该标签对应的文章
func (t *Tag) GetArticles() error {
	o := orm.NewOrm()
	rs := o.Raw("select article_id from blog_tag_article where tag_id=?", t.Id)
	var maps []orm.Params
	if _, err := rs.Values(&maps); err != nil {
		return err
	}
	var articles []*Article = make([]*Article, len(maps))
	for i, v := range maps {
		if id, err := strconv.Atoi(v["article_id"].(string)); err != nil {
			continue
		} else {
			articles[i] = &Article{Id:id}
			articles[i].Read("Id")
		}
	}
	t.Articles = articles
	return nil
}

func (t *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}
