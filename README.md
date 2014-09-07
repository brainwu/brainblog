#Brain Blog

##基于beego写的博客项目

[我的博客地址](http://www.brainwu.com)

##如何使用
* go get github.com/brainwu/brainblog
* 修改conf/app.conf 的配置 根据自己的配置修改
* 现在数据库中创建数据库 create database brainblog;
* 执行mysql -u root -p -D brainblog < brainblog.sql
* go run main.go
* http://localhost/manage 登录到后台
* 账号密码都是admin
