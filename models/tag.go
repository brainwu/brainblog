package models

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

func (t *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}
