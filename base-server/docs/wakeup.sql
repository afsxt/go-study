CREATE DATABASE IF NOT EXISTS wakeup;

DROP TABLE IF EXISTS `wakeup_auth`;
CREATE TABLE `wakeup_auth` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `username` varchar(50) DEFAULT '' COMMENT '账号',
                             `password` varchar(50) DEFAULT '' COMMENT '密码',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `wakeup_auth` (`id`, `username`, `password`) VALUES ('1', 'test', 'test123');

DROP TABLE IF EXISTS `wakeup_article`;
CREATE TABLE `wakeup_label` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `origin_id` int(10) unsigned DEFAULT '0' COMMENT '被标注数据ID',
  `text` varchar(255) DEFAULT '' COMMENT '文本',
  `audio` varchar(255) DEFAULT '' COMMENT '路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标注数据';