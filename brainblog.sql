drop table IF exists `blog_article`;
drop table IF exists `blog_user`;

CREATE TABLE `blog_user` (
  `id` int unsigned NOT NULL auto_increment,
  `name` varchar(15) NOT NULL default '' COMMENT '用户名',
  `account` varchar(30) NOT NULL default '' COMMENT '登陆账号',
  `password` varchar(30) NOT NULL default '' COMMENT '密码',
  `avatar` varchar(120) COMMENT '用户头像',
  `personal_profile` varchar(140) COMMENT '个人简介',
  `email` varchar(50) NOT NULL default '' COMMENT '邮箱',
  `country` varchar(30) default '' COMMENT '国家',
  `province` varchar(30) default '' COMMENT '省',
  `city` varchar(30) default '' COMMENT '市',
  `role` int COMMENT '角色',
  `login_count` mediumint(8) unsigned default '0' COMMENT '登录次数',
  `last_ip` varchar(15) default '0' COMMENT '最后登录ip',
  `last_login` datetime default '0000-00-00 00:00:00' COMMENT '最后登录时间',
  PRIMARY KEY  (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into `blog_user` (name, account, password) values('brainwu', 'admin', 'admin');

create table `blog_article` (
  `id` int unsigned not null auto_increment,
  `user_id` int unsigned not null ,
  `user_name` varchar(15) not null,
  `title` varchar(30) not null,
  `url` varchar(140) not null,
  `content` text not null,
  `create_time` datetime default '1992-12-15 00:00:00',
  `update_time` datetime default '1992-12-15 00:00:00',
  `reply_num` int default 0,
  `click_num` int default 0,
  `tags` varchar(140) default '',
  PRIMARY KEY  (`id`),
  foreign key(user_id) references blog_user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into `blog_article` values(1, 1, 'brainwu', 'The five things make go fast', 'http://www.baidu.com',
'<p>Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型&nbsp;</p>','2014-08-25 06:00:00', '2014-08-25 06:00:00', 0, 0, '');
insert into `blog_article` values(2, 1, 'brainwu', 'Go interface and nil', 'http://www.baidu.com',
'<p>Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型&nbsp;</p>','2014-08-25 06:00:00', '2014-08-25 06:00:00', 0, 0, '');
insert into `blog_article` values(3, 1, 'brainwu', 'Go reflect', 'http://www.baidu.com',
'<p>Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型&nbsp;</p>','2014-08-25 06:00:00', '2014-08-25 06:00:00', 0, 0, '');
insert into `blog_article` values(4, 1, 'brainwu', 'Junit', 'http://www.baidu.com',
'<p>Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型&nbsp;</p>','2014-08-25 06:00:00', '2014-08-25 06:00:00', 0, 0, '');
insert into `blog_article` values(5, 1, 'brainwu', 'Hello Girl', 'http://www.baidu.com',
'<p>Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型&nbsp;</p>','2014-08-25 06:00:00', '2014-08-25 06:00:00', 0, 0, '');
insert into `blog_article` values(6, 1, 'brainwu', 'To be NO.1', 'http://www.baidu.com',
'<p>Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型&nbsp;</p>','2014-08-25 06:00:00', '2014-08-25 06:00:00', 0, 0, '');
insert into `blog_article` values(7, 1, 'brainwu', 'I have a dream.', 'http://www.baidu.com',
'<p>Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型&nbsp;</p>','2014-08-25 06:00:00', '2014-08-25 06:00:00', 0, 0, '');