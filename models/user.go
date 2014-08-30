package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	Name string `orm:"size(15);unique"`
	Account string `orm:"size(30);unique"`
	Password string `orm:"size(30)"`
	Avatar string	`orm:"size(120);null"`
	PersonalProfile string `orm:"size(140);null"`
	Email string `orm:"size(30);null"`
	Country string `orm:"size(30);null"`
	Province string `orm:"size(30);null"`
	City string `orm:"size(30);null"`
	Role int `orm:"null"`
	LastLogin time.Time `orm:"type(datetime);"`
	LoginCount int
	LastIp int
}


func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}

func (u *User) QueryByName() error {
	return orm.NewOrm().QueryTable(u).Filter("Name", u.Name).One(u)
}

func (u *User) QueryByAccount() error {
	return orm.NewOrm().QueryTable(u).Filter("Account", u.Account).One(u)
}

func (u *User)TableName() string {
	return TableName("user")
}

