DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_label` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `origin_id` int(10) unsigned DEFAULT '0' COMMENT '被标注数据ID',
  `text` varchar(255) DEFAULT '' COMMENT '文本',
  `audio` varchar(255) DEFAULT '' COMMENT '路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标注数据';