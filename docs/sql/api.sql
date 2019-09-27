CREATE TABLE `blog_auth` (
  `id` int(10) unsigned not null auto_increment,
  `username` varchar(30)  not null default '' comment '账号',
  `password` varchar(50) not null default '' comment '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');