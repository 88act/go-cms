CREATE TABLE `col_collect` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `userid` int(10) DEFAULT '0' COMMENT '用户id',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '名称',
  `url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '路径',
  `url_page` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '翻页路径',
  `to_table` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应表',
  `now_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '当前id',
  `page_now` int(10) DEFAULT '0'  COMMENT '当前页码',
  `page_start` int(10) DEFAULT '0'  COMMENT '开始页码',
  `page_end` int(10) DEFAULT '0'  COMMENT '结束页码',
  `status_run` smallint(2) DEFAULT '0' COMMENT '运行状态:0停止 1采集中',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未审核 1有效  2未通过审核 3草稿',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='采集表';


CREATE TABLE `col_key_field` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id', 
  `col_id` int(11) DEFAULT '0' COMMENT '采集id', 
  `to_key` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标识',
  `to_field` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '字段',  
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未审核 1有效  2未通过审核 3草稿',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`), 
  KEY `idx_status` (`status`) ,
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='采集标识字段对应表';

CREATE TABLE `col_hsj` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `col_id` int(11) DEFAULT '0' COMMENT '采集id',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '原URL',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '内容',
  `pub_time` datetime DEFAULT NULL COMMENT '发布时间',
  `pub_unit` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '发布单位',
  `log` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '经度',
  `lat` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '纬度',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未审核 1有效  2未通过审核 3草稿',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='海事信息内容表';

