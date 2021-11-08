CREATE TABLE `gin_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` varchar(255) DEFAULT NULL COMMENT '关联tag',
  `title` varchar(128) DEFAULT NULL COMMENT '标题',
  `desc` varchar(255) DEFAULT '' COMMENT '描述',
  `content` text COMMENT '内容',
  `created_by` varchar(32) DEFAULT NULL COMMENT '创建人',
  `modified_by` varchar(32) DEFAULT NULL COMMENT '修改人',
  `state` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态 0-正常 1-隐藏',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `gin_auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `state` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态 0-正常 1-隐藏',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `gin_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL COMMENT '名称',
  `created_by` varchar(32) DEFAULT NULL COMMENT '添加人',
  `modified_by` varchar(32) DEFAULT NULL COMMENT '修改人',
  `state` tinyint(2) DEFAULT '0' COMMENT '状态 0-正常 1-隐藏',
  `ctime` int(11) DEFAULT '0' COMMENT '创建时间',
  `mtime` int(11) DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;