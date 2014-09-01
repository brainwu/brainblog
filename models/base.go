package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/astaxie/beego.v1/orm"
	beego "gopkg.in/astaxie/beego.v1"
	"fmt"
	"runtime"
	"path/filepath"
	"crypto/md5"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User), new(Article), new(Option), new(Tag), new(TagArticle))

}

func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}

func GetOptions() map[string]string {
	var o *Option = new(Option)
	var result []Option;
	orm.NewOrm().QueryTable(o).All(&result)
	var options map[string]string = make(map[string]string)
	for _, v := range result {
		options[v.Name] = v.Value
	}
	return options
}

func Md5(data []byte) string {
	//将返回的16字节的数据转换成32个字节的16进制表示的字符串
	return fmt.Sprintf("%x", md5.Sum(data))
}
