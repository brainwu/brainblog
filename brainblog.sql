drop table if exists `blog_tag_article`;
drop table if exists `blog_article`;
drop table if exists `blog_user`;
drop table if exists `blog_option`;
drop table if exists `blog_tag`;


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
  `tags_str` varchar(140) default '',
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

create table `blog_option` (
  `id` int unsigned not null auto_increment,
  `name` varchar(30) not null default '',
  `value` text not null,
  primary key (`id`),
  unique key `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into `blog_option` values(1, 'about_me', '90后 Developer 非单身');
insert into `blog_option` values(2, 'facebook', 'https://www.facebook.com/profile.php?id=100006836388062');
insert into `blog_option` values(3, 'sina', 'http://weibo.com/1793804425/profile?topnav=1&wvr=5&user=1');
insert into `blog_option` values(4, 'twitter', 'https://twitter.com/Brain_Wu');


create table `blog_tag` (
  `id` int unsigned not null auto_increment,
  `name` varchar(15) not null default '',
  `Count` int not null default 0 COMMENT '使用当前标签的文章数量',
  primary key  (`id`),
  unique key `name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into `blog_tag` values(1, 'android', 0);
insert into `blog_tag` values(2, 'golang', 0);

create table `blog_tag_article` (
  `id` int unsigned not null auto_increment,
  `tag_id` int unsigned not null ,
  `article_id` int unsigned not null ,
  primary key (`id`),
  foreign key(article_id) references blog_article(id),
  foreign key(tag_id) references blog_tag(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;