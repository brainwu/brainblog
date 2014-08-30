package models

import (
	"github.com/astaxie/beego/orm"
)

type Option struct {
	Id int
	Name string
	Value string
}

func (o *Option) TableName() string {
	return TableName("option")
}

func (o *Option) Insert() error {
	if _, err := orm.NewOrm().Insert(o); err != nil {
		return err
	}
	return nil
}

func (o *Option) Delete() error {
	if _, err := orm.NewOrm().Delete(o); err != nil {
		return err
	}
	return nil
}

func (o *Option) Read(fields ...string) error {
	if err := orm.NewOrm().Read(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *Option) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(o, fields...); err != nil {
		return err
	}
	return nil
}
