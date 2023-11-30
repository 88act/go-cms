 
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for basic_area
-- ----------------------------
DROP TABLE IF EXISTS `basic_area`;
CREATE TABLE `basic_area` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `pid` int(11) DEFAULT '0' COMMENT '父栏目',
  `areaname` varchar(50) NOT NULL DEFAULT '' COMMENT '地区名称',
  `shortname` varchar(50) DEFAULT '' COMMENT '简称',
  `areacode` varchar(10) DEFAULT '' COMMENT '电话区号',
  `zipcode` varchar(10) DEFAULT '' COMMENT '邮编',
  `pinyin` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '拼音',
  `py` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '拼音简写',
  `lng` varchar(20) DEFAULT '',
  `lat` varchar(20) DEFAULT '',
  `level` tinyint(1) DEFAULT '0' COMMENT '级别',
  `position` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `mergername` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '组合名称',
  `sort` tinyint(3) unsigned DEFAULT '50' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='地区信息表';

-- ----------------------------
-- Table structure for basic_dict
-- ----------------------------
DROP TABLE IF EXISTS `basic_dict`;
CREATE TABLE `basic_dict` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '编码id',
  `type` smallint(6) DEFAULT '0' COMMENT '类别:1=普通 1=系统 2=ims 3 =活动 99=树状',
  `pid` bigint(20) DEFAULT '0' COMMENT '父id',
  `module` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '模块名',
  `label` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '名称',
  `value` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'key',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '图标',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '备注',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` smallint(6) DEFAULT '1' COMMENT '状态',
  `key` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL COMMENT '创建',
  `updated_at` datetime DEFAULT NULL COMMENT '更新',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`) USING BTREE,
  KEY `module` (`module`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `sort` (`sort`),
  KEY `label` (`label`),
  KEY `value` (`value`) USING BTREE,
  KEY `type` (`type`),
  KEY `deleted_at` (`deleted_at`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7273 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='常量字典表';

-- ----------------------------
-- Table structure for basic_email
-- ----------------------------
DROP TABLE IF EXISTS `basic_email`;
CREATE TABLE `basic_email` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录id',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户id',
  `user_id_send` bigint(20) DEFAULT '0' COMMENT '发送人id:0表示系统',
  `email` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收件邮箱',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `verily_type` smallint(6) DEFAULT '0' COMMENT '验证类型:1注册,2密码登录,3手机登录,4修改密码',
  `verily_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '验证码',
  `attachment` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '附件',
  `status` smallint(6) DEFAULT '0' COMMENT '记录状态:0存草稿 1已发送 2发送失败 3被拒收',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_email` (`email`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='邮件记录表';

-- ----------------------------
-- Table structure for basic_file
-- ----------------------------
DROP TABLE IF EXISTS `basic_file`;
CREATE TABLE `basic_file` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `guid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '唯一id',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `user_id_sys` bigint(20) DEFAULT '0' COMMENT '系统用户id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件名',
  `cat_id` bigint(20) DEFAULT '0' COMMENT '栏目id',
  `module` smallint(6) DEFAULT '0' COMMENT '模块id',
  `media_type` smallint(6) DEFAULT '0' COMMENT '媒体类型: 1 图片 2 视频 3音频 4文档 ',
  `driver` smallint(6) DEFAULT '0' COMMENT '存储位置:0本地1阿里云2腾讯云',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件路径',
  `file_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件类型',
  `ext` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件类型',
  `size` int(10) unsigned DEFAULT '0' COMMENT '文件大小[k]',
  `md5` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'md5值',
  `sha1` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'sha散列值',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `download` int(10) unsigned DEFAULT '0' COMMENT '下载次数',
  `used_time` int(11) DEFAULT '0' COMMENT '使用次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `status` smallint(6) DEFAULT '0' COMMENT '状态: 3=删除 ',
  PRIMARY KEY (`id`),
  KEY `guid` (`guid`),
  KEY `md5` (`md5`),
  KEY `userid` (`user_id`),
  KEY `module` (`module`),
  KEY `sha1` (`sha1`),
  KEY `deleted_at` (`deleted_at`),
  KEY `cat_id` (`cat_id`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件附件表';

-- ----------------------------
-- Table structure for basic_file_cat
-- ----------------------------
DROP TABLE IF EXISTS `basic_file_cat`;
CREATE TABLE `basic_file_cat` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0未审核 1审核  2未通过审核 3 草稿',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_sort` (`sort`),
  KEY `idx_pid` (`pid`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件栏目';

-- ----------------------------
-- Table structure for basic_msg
-- ----------------------------
DROP TABLE IF EXISTS `basic_msg`;
CREATE TABLE `basic_msg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '消息id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `send_user_id` bigint(20) unsigned DEFAULT '0' COMMENT '发送id: 根据msg_type可能对应不同的用户 0表示系统',
  `opt_type` smallint(6) DEFAULT '0' COMMENT '操作类型: 1=验证好友 2=名片消息,3= 缴费通知 4= 群成员确认进群 5=群主审核邀请进群 6=主播开播消息',
  `opt_type_old` tinyint(1) DEFAULT '0' COMMENT '操作类型:(原始的opt_type值)',
  `msg_type` smallint(5) unsigned DEFAULT '0' COMMENT '类别:1=系统消息 2=活动消息 3=私人消息 4 =财务信息',
  `detail` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '消息内容 ',
  `detail_id` bigint(20) unsigned DEFAULT '0' COMMENT '消息内容id',
  `rela_id` bigint(20) DEFAULT '0' COMMENT '关联id',
  `params` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '参数',
  `status_read` smallint(6) DEFAULT '0' COMMENT '查看:0未读 1已读',
  `status` smallint(5) unsigned DEFAULT '0' COMMENT '状态: 0未审核 1已审核',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '0正常 大于0表示软删除时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_status` (`status`),
  KEY `idx_opt_type` (`opt_type`),
  KEY `idx_msg_type` (`msg_type`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知信息表';

-- ----------------------------
-- Table structure for basic_sms
-- ----------------------------
DROP TABLE IF EXISTS `basic_sms`;
CREATE TABLE `basic_sms` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录id',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户id',
  `user_id_send` bigint(20) DEFAULT '0' COMMENT '发送人id:0表示系统',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT ' 标题',
  `content` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 内容',
  `templet_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '模板id',
  `verily_type` smallint(6) DEFAULT '0' COMMENT '验证类型:1注册,2密码登录,3手机登录,4修改密码',
  `verily_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '验证码',
  `attachment` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '附件',
  `status` smallint(6) DEFAULT '0' COMMENT '记录状态:0存草稿 1已发送 2发送失败 3被拒收',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_phone` (`phone`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='短信记录表';

-- ----------------------------
-- Table structure for cms_art
-- ----------------------------
DROP TABLE IF EXISTS `cms_art`;
CREATE TABLE `cms_art` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `pid` bigint(20) DEFAULT '0' COMMENT '父id:章节集合',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `cat_id` bigint(20) DEFAULT '0' COMMENT '类别',
  `cat_id_sys` bigint(20) DEFAULT '0' COMMENT '系统类别',
  `type` smallint(6) DEFAULT '0' COMMENT '文章类型',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文章标题',
  `desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文章摘要',
  `tag_list` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标签列表',
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '来源',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '插图',
  `file_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '媒体列表',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '链接地址',
  `be_top` tinyint(1) DEFAULT '0' COMMENT '置顶:0否 1是',
  `total_whole` int(11) DEFAULT '0' COMMENT '综合指数',
  `total_share` int(11) DEFAULT '0' COMMENT '总分享',
  `total_fav` int(11) DEFAULT '0' COMMENT '总收藏',
  `total_discuss` int(11) DEFAULT '0' COMMENT '总评论',
  `total_click` int(11) DEFAULT '0' COMMENT '总点击',
  `total_star` int(11) DEFAULT '0' COMMENT '总星',
  `total_good` int(11) DEFAULT '0' COMMENT '总赞',
  `total_poor` int(11) DEFAULT '0' COMMENT '总踩',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0未审核 1审核通过 2未通过审核 3草稿 ',
  `verify_msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '审核信息',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `cat_id` (`cat_id`),
  KEY `cat_id_sys` (`cat_id_sys`),
  KEY `type` (`type`),
  KEY `total_click` (`total_click`),
  KEY `total_whole` (`total_whole`),
  KEY `status` (`status`),
  KEY `created_at` (`created_at`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章表';

-- ----------------------------
-- Table structure for cms_blog
-- ----------------------------
DROP TABLE IF EXISTS `cms_blog`;
CREATE TABLE `cms_blog` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户id',
  `title` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标题',
  `desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '简介',
  `image` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '缩略图',
  `file_list` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件列表',
  `email` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '邮件',
  `area_id` int(11) DEFAULT '0' COMMENT '地区id',
  `address` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '地址',
  `total_whole` int(11) DEFAULT '0' COMMENT '综合指数',
  `total_click` int(11) DEFAULT '0' COMMENT '总点击',
  `total_fan` int(11) DEFAULT '0' COMMENT '总粉丝',
  `total_good` int(11) DEFAULT '0' COMMENT '总赞',
  `total_poor` int(11) DEFAULT '0' COMMENT '总踩',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '状态:0未审核 1审核通过 2未通过审核 3黑名单',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `total_whole` (`total_whole`),
  KEY `total_click` (`total_click`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='blog表';

-- ----------------------------
-- Table structure for cms_cat
-- ----------------------------
DROP TABLE IF EXISTS `cms_cat`;
CREATE TABLE `cms_cat` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户id',
  `be_sys` tinyint(1) DEFAULT '0' COMMENT '系统分类: 0=否 1=是',
  `obj_id` bigint(20) DEFAULT '0' COMMENT '专题id:blog/crm/专题',
  `type` smallint(6) DEFAULT '0' COMMENT '栏目类型:1crm ,2 blog 3 group ',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标题',
  `desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '摘要',
  `tag_list` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标签列表:id1,id2,id3,',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '插图',
  `file_list` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件',
  `be_top` tinyint(1) DEFAULT '0' COMMENT '置顶:0否 1是',
  `be_nav` tinyint(1) DEFAULT '0' COMMENT '是否导航',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '状态:0未审核 1审核通过 2未通过审核 3黑名单',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `obj_id` (`obj_id`),
  KEY `sort` (`sort`),
  KEY `type` (`type`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='栏目表';

-- ----------------------------
-- Table structure for cms_cat_art
-- ----------------------------
DROP TABLE IF EXISTS `cms_cat_art`;
CREATE TABLE `cms_cat_art` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户id',
  `cat_id` bigint(20) DEFAULT '0' COMMENT '栏目id',
  `art_id` bigint(20) DEFAULT '0' COMMENT '文章id',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0未审核 1审核 2未通过审核 3草稿',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cat_id` (`cat_id`),
  KEY `art_id` (`art_id`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='栏目文章表';

-- ----------------------------
-- Table structure for cms_detail
-- ----------------------------
DROP TABLE IF EXISTS `cms_detail`;
CREATE TABLE `cms_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `art_id` bigint(20) unsigned NOT NULL COMMENT '文章id',
  `detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '详细',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `art_id` (`art_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='详细内容';

-- ----------------------------
-- Table structure for cms_discuss
-- ----------------------------
DROP TABLE IF EXISTS `cms_discuss`;
CREATE TABLE `cms_discuss` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` bigint(20) DEFAULT '0' COMMENT '父ID',
  `art_id` bigint(20) unsigned DEFAULT '0' COMMENT 'id',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  `user_id_at` bigint(20) unsigned DEFAULT NULL COMMENT '用户At',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标题',
  `detail` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '内容',
  `file_list` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '图片',
  `total_good` int(11) DEFAULT '0' COMMENT '总赞',
  `total_poor` int(11) DEFAULT '0' COMMENT '总踩',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0未审核 1审核通过 2未通过审核  ',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `art_id` (`art_id`),
  KEY `user_id` (`user_id`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论表';

-- ----------------------------
-- Table structure for cms_group
-- ----------------------------
DROP TABLE IF EXISTS `cms_group`;
CREATE TABLE `cms_group` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `type` int(11) DEFAULT '0' COMMENT '群组类别: 1音乐 2舞蹈 3书法 4 棋艺',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '名称',
  `desc` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '内容',
  `tag_list` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标签',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '插图',
  `file_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0未审核 1审核 2未通过审核 3 草稿',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `type` (`type`),
  KEY `user_id` (`user_id`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='圈子群组表';

-- ----------------------------
-- Table structure for cms_tag
-- ----------------------------
DROP TABLE IF EXISTS `cms_tag`;
CREATE TABLE `cms_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0未审核 1审核 2未通过审核 3 草稿',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sort` (`sort`),
  KEY `user_id` (`user_id`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='标签表';

-- ----------------------------
-- Table structure for cms_visitor
-- ----------------------------
DROP TABLE IF EXISTS `cms_visitor`;
CREATE TABLE `cms_visitor` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '访客id',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `type` smallint(6) DEFAULT '0' COMMENT '类型:0文章,1blog,2 team',
  `obj_id` bigint(20) DEFAULT '0' COMMENT '对象Id',
  `plat_type` smallint(6) DEFAULT '0' COMMENT '平台:1 wx 2 web,3 andriod,4 ios',
  `client_ip` bigint(20) DEFAULT '0' COMMENT '客户端ip',
  `status` smallint(6) DEFAULT '0' COMMENT '状态: 0未审核 1审核 2未通过审核 3 草稿',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `type` (`type`),
  KEY `obj_id` (`obj_id`),
  KEY `user_id` (`user_id`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='访客表';

-- ----------------------------
-- Table structure for mem_depart
-- ----------------------------
DROP TABLE IF EXISTS `mem_depart`;
CREATE TABLE `mem_depart` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` bigint(20) unsigned DEFAULT '0' COMMENT '父id',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '管理员id',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '名称',
  `encode` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  `desc` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址',
  `phone` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '联系电话',
  `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '0' COMMENT '配图',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0未审核 1审核通过 2未通过审核 3草稿',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `pid` (`pid`),
  KEY `user_id` (`user_id`),
  KEY `status` (`status`),
  KEY `sort` (`sort`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部门表';

-- ----------------------------
-- Table structure for mem_logs
-- ----------------------------
DROP TABLE IF EXISTS `mem_logs`;
CREATE TABLE `mem_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '商户id',
  `type` smallint(6) DEFAULT '0' COMMENT '类型: 1登录 2退出 3增加用户4 增加任务',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '备注',
  `ip` int(10) unsigned DEFAULT '0' COMMENT 'ip',
  `ip_addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'ip城市',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0=不正常,1=正常',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `type` (`type`),
  KEY `user_id` (`user_id`),
  KEY `deleted_at` (`deleted_at`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户日志';

-- ----------------------------
-- Table structure for mem_user
-- ----------------------------
DROP TABLE IF EXISTS `mem_user`;
CREATE TABLE `mem_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '会员id',
  `user_id` bigint(20) DEFAULT '0' COMMENT '录入人',
  `user_type` smallint(6) DEFAULT '0' COMMENT '类型:1员工 2子账号  3管理员 9 平台号',
  `guid` varchar(32) DEFAULT '' COMMENT '唯一id',
  `username` varchar(40) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) DEFAULT '' COMMENT '密码',
  `password_slat` varchar(32) DEFAULT '' COMMENT '密码盐',
  `nickname` varchar(200) DEFAULT '' COMMENT '昵称',
  `realname` varchar(200) DEFAULT '' COMMENT '真实名',
  `role_list` varchar(255) DEFAULT '' COMMENT '角色list',
  `role` bigint(20) DEFAULT '0' COMMENT '角色',
  `email` varchar(60) DEFAULT '' COMMENT '邮件',
  `mobile` varchar(20) DEFAULT '' COMMENT '手机',
  `card_id` varchar(18) DEFAULT '' COMMENT '身份证',
  `sex` smallint(6) DEFAULT '0' COMMENT '性别: 0保密 1 男 2 女',
  `birthday` datetime DEFAULT NULL COMMENT '生日',
  `avatar` varchar(256) DEFAULT '' COMMENT '头像',
  `job_id` bigint(20) DEFAULT '0' COMMENT '岗位',
  `depart_id` bigint(20) DEFAULT '0' COMMENT '部门',
  `mobile_validated` tinyint(1) DEFAULT '0' COMMENT '验证手机',
  `email_validated` tinyint(1) DEFAULT '0' COMMENT '验证邮箱',
  `cardid_validated` tinyint(1) DEFAULT '0' COMMENT '验证实名',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `status_safe` smallint(6) DEFAULT '0' COMMENT '安全状态:0=正常，1=修改了密码 2=修改了手机号',
  `reg_ip` int(10) unsigned DEFAULT '0' COMMENT '注册ip',
  `login_ip` int(10) unsigned DEFAULT '0' COMMENT '登录ip',
  `login_total` int(11) DEFAULT '0' COMMENT '登录次数',
  `login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `scan` varchar(32) DEFAULT '' COMMENT '扫码',
  `scan_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '扫码',
  `setting` varchar(1000) DEFAULT '' COMMENT '设置值',
  `rtc_model` smallint(6) DEFAULT '0' COMMENT '远程协助模式',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0=审核中,1=审核通过 2=审核未通过3=禁止登录4=非法信息5:已注销6:非法攻击',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`),
  KEY `mobile` (`mobile`),
  KEY `username` (`username`),
  KEY `email` (`email`),
  KEY `card_id` (`card_id`),
  KEY `realname` (`realname`),
  KEY `sex` (`sex`),
  KEY `job_id` (`job_id`),
  KEY `depart_id` (`depart_id`),
  KEY `scan` (`scan`),
  KEY `user_type` (`user_type`),
  KEY `deleted_at` (`deleted_at`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- ----------------------------
-- Table structure for mem_user_safe
-- ----------------------------
DROP TABLE IF EXISTS `mem_user_safe`;
CREATE TABLE `mem_user_safe` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `safe_type` smallint(6) DEFAULT '0' COMMENT '类型:1改密码 2改手机号 3改用户名 4实名认证',
  `val_old` varchar(500) DEFAULT '' COMMENT '旧值',
  `val_new` varchar(500) DEFAULT '' COMMENT '新值',
  `ip` int(10) unsigned DEFAULT '0' COMMENT 'ip',
  `ip_addr` varchar(255) DEFAULT '' COMMENT 'ip城市',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0=不正常,1=正常',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `safe_type` (`safe_type`),
  KEY `user_id` (`user_id`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='安全日志';

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `path` varchar(256) DEFAULT '' COMMENT 'api路径',
  `desc` varchar(256) DEFAULT '' COMMENT 'api中文描述',
  `model` varchar(100) DEFAULT '' COMMENT '所属模块',
  `api_group` varchar(256) DEFAULT '' COMMENT 'api组',
  `method` varchar(256) DEFAULT 'POST' COMMENT '方法',
  `status` smallint(6) DEFAULT '0' COMMENT '状态',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `path` (`path`),
  KEY `method` (`method`),
  KEY `model` (`model`),
  KEY `api_group` (`api_group`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='API表';

-- ----------------------------
-- Table structure for sys_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE `sys_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  `log_type` smallint(6) DEFAULT '0' COMMENT 'log类型',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `error_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '请求ip',
  `ip_addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '地址',
  `status` int(11) DEFAULT NULL COMMENT '请求状态',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `log_type` (`log_type`),
  KEY `user_id` (`user_id`),
  KEY `deleted_at` (`deleted_at`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=906 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统日志';

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由name',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '路由path',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标题',
  `show_link` tinyint(1) DEFAULT '0' COMMENT '是否隐藏',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '前端文件路径',
  `sort` bigint(20) DEFAULT '0' COMMENT '排序',
  `level` int(10) unsigned DEFAULT NULL COMMENT '等级',
  `frame_src` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '外链',
  `param` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '参数',
  `keep_alive` tinyint(1) DEFAULT '0' COMMENT '保持连接',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '图标',
  `auths` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '按钮权限',
  `active_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '活动链接',
  `be_sys` tinyint(1) DEFAULT '0' COMMENT '平台菜单',
  `status` smallint(6) DEFAULT '0' COMMENT '状态',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `sort` (`sort`),
  KEY `deleted_at` (`deleted_at`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `status` (`status`),
  KEY `be_sys` (`be_sys`)
) ENGINE=InnoDB AUTO_INCREMENT=199 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='基础菜单';

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) DEFAULT NULL,
  `pid` bigint(20) unsigned DEFAULT '0' COMMENT '父id',
  `role_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '角色名',
  `role_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '角色编码',
  `api_list` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `menu_list` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `default_menu` bigint(20) DEFAULT '0' COMMENT '默认菜单',
  `status` smallint(6) DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `role_name` (`role_name`),
  KEY `role_code` (`role_code`),
  KEY `deleted_at` (`deleted_at`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色';
