/*
Navicat MySQL Data Transfer

Source Server         : 88act-120.24.150.47
Source Server Version : 80013
Source Host           : 120.24.150.47:53306
Source Database       : go-cms-test

Target Server Type    : MYSQL
Target Server Version : 80013
File Encoding         : 65001

Date: 2022-04-21 14:06:56
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for act_act
-- ----------------------------
DROP TABLE IF EXISTS `act_act`;
CREATE TABLE `act_act` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '活动id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `cat_id` int(11) NOT NULL DEFAULT '0' COMMENT '系统类别ID ',
  `cat_id_user` int(11) DEFAULT '0' COMMENT '用户自建类别ID',
  `be_online` tinyint(1) DEFAULT '0' COMMENT '类型：0=线上  1=线下',
  `act_type` varchar(500) DEFAULT '' COMMENT '类型（多选）：1=视频直播，2=语音直播，3=视频点播 4=文本文档 5=群组聊天  6=室内活动 7=户外活动',
  `be_multi` tinyint(1) DEFAULT '0' COMMENT '多期活动类型：0=否 1=是',
  `title` varchar(150) DEFAULT NULL COMMENT '标题',
  `desc` varchar(500) DEFAULT NULL COMMENT '简介',
  `thumb` bigint(20) DEFAULT '0' COMMENT '课程缩略图',
  `media_list` varchar(255) DEFAULT '' COMMENT '媒体列表（id列 1,2,3）',
  `map_address` varchar(500) DEFAULT NULL COMMENT '上课地址',
  `area_id` int(11) DEFAULT '0' COMMENT '地区id',
  `lng` decimal(9,6) DEFAULT '0.000000' COMMENT '经度',
  `lat` decimal(9,6) DEFAULT '0.000000' COMMENT '纬度',
  `begin_time` int(11) unsigned DEFAULT '0' COMMENT '活动开始时间',
  `end_time` int(11) unsigned DEFAULT '0' COMMENT '活动结束时间',
  `live_time` int(11) unsigned DEFAULT '0' COMMENT '直播开始时间）',
  `live_end` int(11) unsigned DEFAULT '0' COMMENT '直播结束时间）',
  `presenter` varchar(200) DEFAULT NULL COMMENT '主持人/讲师/嘉宾',
  `phone_kefu` varchar(100) DEFAULT NULL COMMENT '报名客服电话',
  `phone_hezu` varchar(100) DEFAULT NULL COMMENT '合作联系电话',
  `wx` varchar(100) DEFAULT NULL COMMENT '微信',
  `qq` varchar(100) DEFAULT NULL COMMENT 'QQ',
  `group_id` varchar(100) DEFAULT NULL COMMENT '群组id',
  `end_join_time` int(11) unsigned DEFAULT '0' COMMENT '活动结束报名',
  `price` int(11) DEFAULT '0' COMMENT ' 默认票价（价格:分），0表示免费',
  `refund_type` tinyint(1) DEFAULT '0' COMMENT '退费类型：1=不可退款 2=活动开始前都可退款 3=活动开始前n天可退费',
  `refund_days` tinyint(2) DEFAULT '0' COMMENT ' 活动开始前n天退费',
  `be_showjoin` tinyint(1) DEFAULT '0' COMMENT '是否显示报名人数：0=显示 1=不显示',
  `act_status` tinyint(1) DEFAULT '0' COMMENT '状态：0未开始,1报名中,2停止报名,3进行中,4已结束,5评价期，6结束',
  `max_join` int(11) DEFAULT '0' COMMENT '最大报名人数',
  `now_join` int(11) DEFAULT '0' COMMENT '当前报名人数',
  `be_choice` tinyint(1) DEFAULT '0' COMMENT '是否精选：0=否 1=是',
  `be_week` tinyint(1) DEFAULT '0' COMMENT '是否周末：0=否 1=是',
  `be_vip` tinyint(1) DEFAULT '0' COMMENT '是否会员价：0=否 1=是',
  `total_whole` int(11) DEFAULT '0' COMMENT '综合指数',
  `total_hot` tinyint(1) DEFAULT '0' COMMENT '推荐级别 1～10级',
  `total_share` int(11) DEFAULT '0' COMMENT '分享数',
  `total_fav` int(11) DEFAULT '0' COMMENT '收藏数',
  `total_join` int(11) DEFAULT '0' COMMENT '总报名人数',
  `total_discuss` int(11) DEFAULT '0' COMMENT '总评论',
  `total_click` int(11) DEFAULT '0' COMMENT '总点击（热门）',
  `total_star` int(11) DEFAULT '0' COMMENT '总评',
  `total_star_1` smallint(6) DEFAULT NULL COMMENT '总星评1（教学水平）',
  `total_star_2` smallint(6) DEFAULT NULL COMMENT '总星评2（课程质量）',
  `total_star_3` smallint(6) DEFAULT NULL COMMENT '总星评3（服务态度）',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0未审核 1审核通过 2未通过审核 3草稿 ',
  `status_msg` varchar(255) DEFAULT NULL COMMENT '审核原因',
  `create_time` int(11) unsigned DEFAULT '0' COMMENT '创建时间',
  `delete_time` int(11) unsigned DEFAULT '0' COMMENT '0正常,大于0表示软删除时间',
  `update_time` int(11) unsigned DEFAULT '0' COMMENT '更新时间',
  `createat` time DEFAULT NULL,
  `deleteat` time DEFAULT NULL,
  `updateat` time DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_cat_id` (`cat_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_total_click` (`total_click`),
  KEY `idx_total_discuss` (`total_discuss`),
  KEY `idx_total_join` (`total_join`),
  KEY `idx_be_online` (`be_online`),
  KEY `idx_be_multi` (`be_multi`),
  KEY `idx_be_choice` (`be_choice`),
  KEY `idx_be_vip` (`be_vip`),
  KEY `idx_price` (`price`),
  KEY `idx_refund_type` (`refund_type`),
  KEY `idx_act_status` (`act_status`),
  KEY `idx_status` (`status`),
  KEY `idx_total_hot` (`total_hot`),
  KEY `idx_total_whole` (`total_whole`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_update_time` (`update_time`),
  KEY `idx_delete_time` (`delete_time`),
  KEY `idx_act_act_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8 COMMENT='活动表';

-- ----------------------------
-- Records of act_act
-- ----------------------------
INSERT INTO `act_act` VALUES ('1', '0', '553', '0', '1', '633,634,635', '0', '广州周末相亲活动2222', '网络儿童兴趣课网络儿童兴趣课网络儿童兴趣课网络儿童兴趣课', '10361', '', '', '0', '0.000000', '0.000000', '1589922000', '1987200', '0', '0', '', '', '', '', '', '0', '604800', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '', '1587109224', '0', '1587109224', null, null, null, '2021-03-03 22:59:46', null, '0000-00-00 00:00:00');
INSERT INTO `act_act` VALUES ('2', '0', '553', '0', '0', '', '0', '网络儿童兴趣课2', '音乐会  也要炫酷到底，我们不“箭”不散', '10361', '', '', '0', '0.000000', '0.000000', '0', '0', '0', '0', '', '', '', '', '', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '', '1587109393', '0', '1587109393', null, null, null, null, null, null);
INSERT INTO `act_act` VALUES ('3', '0', '545', '0', '1', '625,626,627,633,634,635', '0', '网络儿童兴趣课3', '音乐会 也要炫酷到底，我们不“箭”不散', '10361', '', '广州市天河区广州天河体育中心', '120103', '113.331575', '23.143232', '0', '0', '0', '0', '', '', '', '', '', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '', '1587110147', '0', '1587110147', null, null, null, null, null, null);
INSERT INTO `act_act` VALUES ('4', '0', '545', '0', '0', '626,628,629,630', '0', '互联网啊啊', ' 也要炫酷到底，我们不“箭”不散网络儿童兴趣课网络儿童兴趣课', '10361', '', '广东省广州市天河区', '130103', '113.368509', '23.130358', '1589922306', '0', '0', '0', '', '', '', '', '', '0', '0', '1212', '2', '12', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '', '1587368879', '0', '1587368879', null, null, null, null, null, null);
INSERT INTO `act_act` VALUES ('5', '0', '544', '0', '0', '625,626,627,628,629,630', '0', '测试 测试 测试 ', ' 也要炫酷到底，我们不“箭”不散', '10361', '', '广州市白云区天河电脑分店', '110102', '113.224061', '23.233246', '1589922300', '1592427960', '0', '0', '', '', '', '', '', '0', '1592611560', '0', '0', '0', '0', '0', '100', '0', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '', '1587374036', '0', '1587374036', null, null, null, null, null, null);
INSERT INTO `act_act` VALUES ('6', '0', '520', '0', '0', '626,628,629,630', '0', '地方 abc study', '射箭：要运动，也要炫酷到底，我们不“箭”不散', '10361', '', '广州市天河区广州天河体育中心', '120103', '113.331575', '23.143232', '1589922302', '1592514360', '0', '0', '', '', '', '', '', '0', '1589922300', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '2', '0', '0', '0', '0', '0', '1', '32323', '1587375584', '0', '1587375584', null, null, null, null, null, null);
 

-- ----------------------------
-- Table structure for basic_file
-- ----------------------------
DROP TABLE IF EXISTS `basic_file`;
CREATE TABLE `basic_file` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `guid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '唯一id',
  `user_id` bigint(10) unsigned DEFAULT '0' COMMENT '用户id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '文件名',
  `module` smallint(2) DEFAULT '0' COMMENT '模块名',
  `media_type` smallint(2) DEFAULT '0' COMMENT '媒体类型',
  `driver` smallint(2) DEFAULT '0' COMMENT '存储位置:0本地1阿里云2腾讯云',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '文件路径',
  `ext` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '文件类型',
  `size` int(10) unsigned DEFAULT '0' COMMENT '文件大小[k]',
  `md5` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT 'md5值',
  `sha1` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT 'sha散列值',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `download` int(10) unsigned DEFAULT '0' COMMENT '下载次数',
  `used_time` int(11) DEFAULT '0' COMMENT '使用次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `guid` (`guid`),
  KEY `md5` (`md5`),
  KEY `userid` (`user_id`),
  KEY `module` (`module`),
  KEY `sha1` (`sha1`),
  KEY `idx_basic_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=123 DEFAULT CHARSET=utf8 COMMENT='文件附件表';

-- ----------------------------
-- Records of basic_file
-- ----------------------------
INSERT INTO `basic_file` VALUES ('99', '85e3ab6bfa3447bebd7365c6a527c971', null, 'image.png', '0', '2', '0', 'res/sys/0/20211021/85e3ab6bfa3447bebd7365c6a527c971.png', 'png', '10585', '', '9fcc669d5704a983019202e7c43f8231d90778d1', null, null, null, '2021-10-21 10:37:03', '2021-10-21 10:37:03', '2021-10-21 10:47:35');
INSERT INTO `basic_file` VALUES ('100', 'ef2a204318a54e79875e89e0b74f9c43', null, '9127206164_1127227391.jpg', '0', '0', '0', 'res/sys/0/20211021/ef2a204318a54e79875e89e0b74f9c43.jpg', 'jpg', '216', '08815fb28531f8ca34253e757c94cf57', '254328283fd1abd5a2c28d05c6aca7c602b3af92', null, null, null, '2021-10-21 10:37:41', '2021-10-21 10:37:41', '2021-10-21 10:47:35');
INSERT INTO `basic_file` VALUES ('101', '06d392e31927424992f1bd00c2db859e', null, '9127185622_1127227391.jpg', '0', '0', '0', 'res/sys/0/20211021/06d392e31927424992f1bd00c2db859e.jpg', 'jpg', '87', '9d6153661084c1ae9a0b980e93b91d11', '5d745a0bba60ad130e2a9724e84297f5ba200a22', null, null, null, '2021-10-21 10:45:56', '2021-10-21 10:45:56', null);
INSERT INTO `basic_file` VALUES ('102', 'e8fb941c4b544cbb887637ae5fb24445', null, '9149743923_1127227391.jpg', '0', '0', '0', 'res/sys/0/20211021/e8fb941c4b544cbb887637ae5fb24445.jpg', 'jpg', '122', 'd48aba5ea473cc8fd98451b90b86a7c7', '407a2d0bf962aff21f1b5c4c8ee1e1753879b68f', null, null, null, '2021-10-21 10:46:00', '2021-10-21 10:46:00', null);
INSERT INTO `basic_file` VALUES ('103', '9857f17b0948498fbef0cee383e1cf77', null, '9168530821_1127227391.jpg', '5', '1', '0', 'res/sys/5/20211021/9857f17b0948498fbef0cee383e1cf77.jpg', 'jpg', '134', '0488731ab1116c18944e182d7a9880c0', '9a7abd40ae4ad508fb775b3725c0c867280ecccf', null, null, null, '2021-10-21 10:48:53', '2021-10-21 10:48:53', null);
INSERT INTO `basic_file` VALUES ('104', '62fe6c98c9ca4720affdd33a8eb72726', null, '9149734951_1127227391.jpg', '5', '1', '0', 'res/sys/5/20211021/62fe6c98c9ca4720affdd33a8eb72726.jpg', 'jpg', '109', '00af06203ae07d0a517f3e5238f2c988', '3b240e3e28e9d6cc654aa4b01d0f6308b177ac58', null, null, null, '2021-10-21 10:48:58', '2021-10-21 10:48:58', null);
INSERT INTO `basic_file` VALUES ('105', '416486ffa3f6437399aa575b20ac3289', null, '9127194473_1127227391.jpg', '5', '1', '0', 'res/sys/5/20211021/416486ffa3f6437399aa575b20ac3289.jpg', 'jpg', '222', '73f5a2dfcce23723b01ed1f202fc3858', '7ac4dd1a2f91adb88c7d78e2cd3ef95aaf47fe4f', null, null, null, '2021-10-21 10:49:00', '2021-10-21 10:49:00', null);
INSERT INTO `basic_file` VALUES ('106', 'a7868d8f10104352a6b662a61ee0bbaa', null, '9168557283_1127227391.jpg', '5', '1', '0', 'res/sys/5/20211021/a7868d8f10104352a6b662a61ee0bbaa.jpg', 'jpg', '174', '88ac247a94b124f6b6e911baad20a9ee', 'c2aff61a441a3d806288a42e4a5ac6497624a92d', null, null, null, '2021-10-21 10:49:04', '2021-10-21 10:49:04', null);
INSERT INTO `basic_file` VALUES ('107', '1d2f623ac3384ce5b9d9a9ae0626f236', null, 'image.png', '0', '2', '0', 'res/sys/0/20211021/1d2f623ac3384ce5b9d9a9ae0626f236.png', 'png', '56516', '', '510958b0850cd8c581e9cdaaa5f631bd7bd7bb12', null, null, null, '2021-10-21 10:49:33', '2021-10-21 10:49:33', null);
INSERT INTO `basic_file` VALUES ('108', 'cf7ce924cb5243d8ad817bb40a8af4d0', null, '9149782097_1127227391.jpg', '0', '2', '0', 'res/sys/0/20211025/cf7ce924cb5243d8ad817bb40a8af4d0.jpg', 'jpg', '72', 'dab5f972e1683ea073b0fa8c138b3d79', '4e16c2584381e3a0f7fa86e762ed4cbd7e4ea962', null, null, null, '2021-10-25 10:23:52', '2021-10-25 10:23:52', null);
INSERT INTO `basic_file` VALUES ('109', 'df11c3eb767444da84acd55e6ae7ca44', null, 'IMG_00000002.jpeg', '0', '0', '0', 'res/sys/0/20211025/df11c3eb767444da84acd55e6ae7ca44.jpeg', 'jpg', '42', '7a9501970d452b6ead874a6cac152275', 'f87d20f4bce3265f641dcbc1351f549fa1369ae9', null, null, null, '2021-10-25 14:35:14', '2021-10-25 14:35:14', null);
INSERT INTO `basic_file` VALUES ('110', 'ba73f34fe70a4b9081b73e74df4d8c3a', null, 'IMG_00000001.jpg', '0', '0', '0', 'res/sys/0/20211025/ba73f34fe70a4b9081b73e74df4d8c3a.jpg', 'jpg', '46', '3e6a00997c1f06b36c87f76a5c349691', '5249623f4ffc279132d5950078c3ab36a07ed580', null, null, null, '2021-10-25 14:35:31', '2021-10-25 14:35:31', null);
INSERT INTO `basic_file` VALUES ('111', '25fbd22ccc8e43cca38ed2abb1e28dab', null, '2021-08-19 14-14-20 的屏幕截图.png', '0', '2', '0', 'res/sys/0/20211025/25fbd22ccc8e43cca38ed2abb1e28dab.png', 'png', '465537', '', '4ff71b0d69fe9dab235985b0917559a147f7a570', null, null, null, '2021-10-25 15:24:40', '2021-10-25 15:24:40', null);
INSERT INTO `basic_file` VALUES ('112', 'de1f0a1055b24277b29792679c5b6ca2', null, 'IMG_00000001.jpeg', '0', '1', '0', 'res/sys/0/20211028/de1f0a1055b24277b29792679c5b6ca2.jpeg', 'jpg', '42', '4b4c9e82410e0cb0bd93ea9e02b36346', '50ee3c69777670d688239fb737275e67e8ecb53e', null, null, null, '2021-10-28 15:23:50', '2021-10-28 15:23:50', null);
INSERT INTO `basic_file` VALUES ('113', 'a8403b380b71449aa745437918332ccf', null, 'logo.jpg', '0', '1', '0', 'res/sys/0/20211109/a8403b380b71449aa745437918332ccf.jpg', 'jpg', '61', 'b29bdb1017513195077308925966062e', 'd7f4b119d798e7a6a0ea3cc3693d4ea3ae7b503f', null, null, null, '2021-11-09 15:49:33', '2021-11-09 15:49:33', null);
INSERT INTO `basic_file` VALUES ('114', '3caa7304adaa4b7793275df74cec50af', null, '44.png', '0', '0', '0', 'res/sys/0/20211122/3caa7304adaa4b7793275df74cec50af.png', 'png', '554', '2ac15e4f6c94d5882702d216a6e2db15', '9a1bf4ad71c6a0c4fe77bf85a0e8f6167df3473c', null, null, null, '2021-11-22 11:13:27', '2021-11-22 11:13:27', null);
INSERT INTO `basic_file` VALUES ('115', '1e9d63db6a77477d9064ee4da8382511', null, '55.png', '1', '1', '0', 'res/sys/1/20211206/1e9d63db6a77477d9064ee4da8382511.png', 'png', '670', 'aea3f85a128fc823c4f5e9ae646f0d2e', '7d14a654fa7cf1b659fc29b9003dfe0d0456f207', null, null, null, '2021-12-06 07:19:51', '2021-12-06 07:19:51', null);
INSERT INTO `basic_file` VALUES ('116', 'f38b1890c1124a6c96fb5fcf5398ab14', null, '88.jpg', '0', '0', '0', 'res/sys/0/20211206/f38b1890c1124a6c96fb5fcf5398ab14.jpg', 'jpg', '121', 'c5d8b2f061f4423926fffdbb69f94ccc', '9a1e2b23462af2867c36434035ece2229342d878', null, null, null, '2021-12-06 07:20:17', '2021-12-06 07:20:17', null);
INSERT INTO `basic_file` VALUES ('117', 'ad283a5b3ba24be0af845eecdada42d0', null, '66.jpg', '0', '0', '0', 'res/sys/0/20211206/ad283a5b3ba24be0af845eecdada42d0.jpg', 'jpg', '2485', '5610b9e0dd5cea5e738a0e599b77764b', '31b3c297fcb6ae43cbc7d4b897999e5c5deb22dd', null, null, null, '2021-12-06 15:22:51', '2021-12-06 15:22:51', null);
INSERT INTO `basic_file` VALUES ('118', '76893fba23cf4d369144743c48b5329c', null, 'aaa.png', '0', '1', '0', 'res/sys/0/20211206/76893fba23cf4d369144743c48b5329c.png', 'png', '142', 'd57108458e6ba562cf94399a628e7fa7', 'f787204b5240b64624e53504a66ebca756b8b46f', null, null, null, '2021-12-06 15:23:13', '2021-12-06 15:23:13', null);
INSERT INTO `basic_file` VALUES ('119', '96e8c56cff2445af800569932467382c', null, '3333.png', '0', '1', '0', 'res/sys/0/20211206/96e8c56cff2445af800569932467382c.png', 'png', '39', '04538004850ee3af2be951222c58f483', '078a418875a2527493c3b377b42f20b17ae0888b', null, null, null, '2021-12-06 15:27:26', '2021-12-06 15:27:26', null);
INSERT INTO `basic_file` VALUES ('120', '32abd751c01c4466ac4dc4ba581c1663', null, 'ccc.png', '1', '1', '0', 'res/sys/1/20211207/32abd751c01c4466ac4dc4ba581c1663.png', 'png', '494', 'fed9f0e16f676d7b306c27e9c0b1685b', '362a54820982838b8411c9c443b8e5c18b36964a', null, null, null, '2021-12-07 14:44:40', '2021-12-07 14:44:40', null);
INSERT INTO `basic_file` VALUES ('121', '1edff067da1e498a8ccd15d6226198ee', null, '55.png', '0', '0', '0', 'res/sys/0/20211220/1edff067da1e498a8ccd15d6226198ee.png', 'png', '670', 'aea3f85a128fc823c4f5e9ae646f0d2e', '7d14a654fa7cf1b659fc29b9003dfe0d0456f207', null, null, null, '2021-12-20 16:56:29', '2021-12-20 16:56:29', null);
INSERT INTO `basic_file` VALUES ('122', '20a6e4a766d14c30ad79933f0d940600', null, '44.png', '0', '0', '0', 'res/sys/0/20211220/20a6e4a766d14c30ad79933f0d940600.png', 'png', '554', '2ac15e4f6c94d5882702d216a6e2db15', '9a1bf4ad71c6a0c4fe77bf85a0e8f6167df3473c', null, null, null, '2021-12-20 17:15:34', '2021-12-20 17:15:34', null);

-- ----------------------------
-- Table structure for bb_pi_contacts
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_contacts`;
CREATE TABLE `bb_pi_contacts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `bm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '部门',
  `fzrxm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '负责人姓名',
  `lxdh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '联系电话',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成日期时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_contacts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='行政管理通讯表';

-- ----------------------------
-- Records of bb_pi_contacts
-- ----------------------------

-- ----------------------------
-- Table structure for bb_pi_department
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_department`;
CREATE TABLE `bb_pi_department` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `ksbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '科室编码',
  `ksmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '科室名称',
  `bzksdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标准科室代码',
  `ybjdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医保局代码',
  `ksjs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '科室简介',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成日期时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `idx_bb_pi_department_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='科室字典';

-- ----------------------------
-- Records of bb_pi_department
-- ----------------------------

-- ----------------------------
-- Table structure for bb_pi_device
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_device`;
CREATE TABLE `bb_pi_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `sbdh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '设备代号',
  `sbmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '设备名称',
  `tpsbts` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '同批设备台数',
  `cd` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '产地',
  `sccj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '生产厂家',
  `sbxh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '设备型号',
  `sbcs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '设备参数',
  `sblx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '设备类型',
  `gmrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '购买日期',
  `yt` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用途',
  `sbjzje` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '设备价值金额',
  `gjsxqk` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '购进时新旧情况',
  `llsjsm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '理论设计寿命',
  `syqk` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '使用情况 ',
  `bz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '备注',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_device_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='设备资源信息表';

-- ----------------------------
-- Records of bb_pi_device
-- ----------------------------

-- ----------------------------
-- Table structure for bb_pi_institution
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_institution`;
CREATE TABLE `bb_pi_institution` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `zzjgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '统一信用代码',
  `jgmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构名称',
  `xzqhdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '行政区划代码',
  `jglx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构类型',
  `jgclrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构成立日期',
  `dwlsgxdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '单位隶属关系代码',
  `jgflgllbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构分类代码',
  `jgfldm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构分类代码',
  `jjlxdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '经济类型代码',
  `dz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '地址',
  `styyzzjgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '组织机构代码',
  `styymc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医院名称',
  `styljgjb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医疗机构级别',
  `styljgdj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医院机构等级',
  `hlwyywz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医院网址',
  `xkzhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '许可证号码',
  `xkxmmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '许可项目名称',
  `xkzyxq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '许可证有效期',
  `xxaqdjbh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '信息安等级保护',
  `xxaqdjbhbh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '信息安等保编号',
  `kbzjjes` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '开办资金金额数',
  `frdb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '法人代表  ',
  `jgszdmzzzdfbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构民族自治标志',
  `sffzjg` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '是否分支机构',
  `jgdemc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构第二名称',
  `jgms` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构描述',
  `yzbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '邮政编码',
  `dhhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '电话号码',
  `dwdzyx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '电子信箱',
  `dsfjgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '第三方信用代码',
  `dsfjgmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '第三方机构名称',
  `xyyxq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '协议有效期',
  `bspt` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '部署平台',
  `jgmsxx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '架构描述',
  `jfmj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机房面积',
  `sfslgd` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '是否双路发电',
  `bz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '备注',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成日期时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_institution_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='机构信息表';

-- ----------------------------
-- Records of bb_pi_institution
-- ----------------------------
INSERT INTO `bb_pi_institution` VALUES ('1', '0', '4444', '5555', '互联网医院', '440304', '01', '2020-09-10', '40', '1', 'A100', '11', '广东省深圳市福田区', '45574403-0', '中山大学附属医院', '3', '0', 'wwww', '123123', '医疗机构执业许可证', '2033-09-18', '3', 'SZCP20200334', '14697', '王校友', '2', '2', '1111', '', '', '', '', null, '', '', '01', '', '', '1', '', '2022-03-22', '1647936999', '1', null, null, null);
INSERT INTO `bb_pi_institution` VALUES ('2', '0', '4444', '5555', '互联网医院', '440304', '01', '2020-09-10', '40', '1', 'A100', '11', '广东省深圳市福田区', '45574403-0', '中山大学附属医院', '3', '0', 'wwww', '123123', '医疗机构执业许可证', '2033-09-18', '3', 'SZCP20200334', '14697', '王校友', '2', '2', '1111', '', '', '', '', null, '', '', '01', '', '', '1', '', '2022-03-23', '1648022847', '1', null, null, null);
INSERT INTO `bb_pi_institution` VALUES ('3', '0', '4444', '5555', '互联网医院', '440304', '01', '2020-09-10', '40', '1', 'A100', '11', '广东省深圳市福田区', '45574403-0', '中山大学附属医院', '3', '0', 'wwww', '123123', '医疗机构执业许可证', '2033-09-18', '3', 'SZCP20200334', '14697', '王校友', '2', '2', '1111', '', '', '', '', null, '', '', '01', '', '', '1', '', '2022-03-23', '1648023208', '1', null, null, null);
INSERT INTO `bb_pi_institution` VALUES ('4', '0', '4444', '5555', '互联网医院', '440304', '01', '2020-09-10', '40', '1', 'A100', '11', '广东省深圳市福田区', '45574403-0', '中山大学附属医院', '3', '0', 'wwww', '123123', '医疗机构执业许可证', '2033-09-18', '3', 'SZCP20200334', '14697', '王校友', '2', '2', '1111', '', '', '', '', null, '', '', '01', '', '', '1', '', '2022-03-23', '1648026641', '1', null, null, null);
INSERT INTO `bb_pi_institution` VALUES ('5', '0', '4444', '5555', '互联网医院', '440304', '01', '2020-09-10', '40', '1', 'A100', '11', '广东省深圳市福田区', '45574403-0', '中山大学附属医院', '3', '0', 'wwww', '123123', '医疗机构执业许可证', '2033-09-18', '3', 'SZCP20200334', '14697', '王校友', '2', '2', '1111', '', '', '', '', null, '', '', '01', '', '', '1', '', '2022-03-28', '1648438141', '1', null, null, null);
INSERT INTO `bb_pi_institution` VALUES ('6', '0', '4444', '5555', '互联网医院', '440304', '01', '2020-09-10', '40', '1', 'A100', '11', '广东省深圳市福田区', '45574403-0', '中山大学附属医院', '3', '0', 'wwww', '123123', '医疗机构执业许可证', '2033-09-18', '3', 'SZCP20200334', '14697', '王校友', '2', '2', '1111', '', '', '', '', null, '', '', '01', '', '', '1', '', '2022-03-28', '1648438141', '1', null, null, null);

-- ----------------------------
-- Table structure for bb_pi_institution_business
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_institution_business`;
CREATE TABLE `bb_pi_institution_business` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `nf` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '月份',
  `pcjgsl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '派出机构数量',
  `zgzs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '职工总数',
  `khffryzs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '客户服务人员总数',
  `rjzzyspbs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '日均坐诊医生数',
  `ywyfmj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '业务用房面积',
  `zsr` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '总收入',
  `zzc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '总支出',
  `zzch` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '总资产',
  `ldzc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '流动资产',
  `dwtz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对外投资',
  `gdzc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '固定资产',
  `wxzcjkbf` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '无形资产及开办费',
  `fz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '负债',
  `jzc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '净资产',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成日期时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_institution_business_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='机构业务开展情况';

-- ----------------------------
-- Records of bb_pi_institution_business
-- ----------------------------

-- ----------------------------
-- Table structure for bb_pi_person
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_person`;
CREATE TABLE `bb_pi_person` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `userid` int(10) unsigned NOT NULL COMMENT 'userid',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `kh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡号',
  `klx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡类型',
  `zjhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '身份证件号码',
  `zjlbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '身份证件类别',
  `xm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '姓名',
  `xbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '性别代码',
  `xbmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '性别名称',
  `hzsx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '患者属性',
  `hyztdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '婚姻状况代码',
  `hyztmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '婚姻状态名称',
  `csrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生日期',
  `mzdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '民族代码',
  `mzmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '民族名称',
  `gjdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '国籍代码',
  `gjmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '国籍名称',
  `jgnbdah` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构内部档案号',
  `gddhhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '固定电话',
  `sjhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '手机号码',
  `dzyj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '电子邮件',
  `whcddm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文化程度代码',
  `whcdmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文化程度名称',
  `zylbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '职业类别代码',
  `zylbmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '职业类别名称',
  `csdxzqhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生地行政区划码',
  `csd` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生地',
  `jzdxzqhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '居住地行政区划码',
  `jzdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '居住地址',
  `hkdxzqhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '户口地行政区划码',
  `hkdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '户口地址',
  `lxrxm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '联系人姓名',
  `lxrdh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '联系人电话',
  `abo` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'ABO 血型',
  `rh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'RH 血型',
  `cblbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '参保类别代码',
  `grdaid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '个人档案ID',
  `yl1` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '预留一',
  `yl2` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '预留二',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `mj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密级',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `userid` (`userid`),
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_person_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='个人基本信息';

-- ----------------------------
-- Records of bb_pi_person
-- ----------------------------

-- ----------------------------
-- Table structure for bb_pi_service_point
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_service_point`;
CREATE TABLE `bb_pi_service_point` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `zzjgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '统一社会信用代码',
  `fwwddm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务网点代码',
  `fwdmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点名称',
  `xzqhdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '行政区划代码',
  `fwdlx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点类型',
  `fwdclrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点成立日期',
  `dwlsgxdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '单位隶属关系代码',
  `fwdflgllbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点分类管理类别代码',
  `fwdfldm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点分类代码',
  `jjlxdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '经济类型代码',
  `dz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '地址',
  `fwdyyjb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点医院级别',
  `fwdyydj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点医院等级',
  `xkzhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '许可证号码',
  `xkxmmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '许可项目名称',
  `xkzyxq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '许可证有效期',
  `kbzjjes` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '开办资金金额数',
  `frdb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '法人代表',
  `fwdszdmzzzdfbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务点地方标志',
  `sffzjg` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '是否分支机构',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `idx_bb_pi_service_point_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='服务点';

-- ----------------------------
-- Records of bb_pi_service_point
-- ----------------------------

-- ----------------------------
-- Table structure for bb_pi_staff
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_staff`;
CREATE TABLE `bb_pi_staff` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `yhrygh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医护人员账号',
  `yhryxm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医护人员姓名',
  `xb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '性别',
  `csrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生日期',
  `sfzh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '身份证号',
  `zjlbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '身份证件类别代码',
  `ksdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '所属科室代码',
  `zyjszwdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '专业技术职务代码',
  `zyjszwlbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '专业技术职务类别代码',
  `zzlbmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '资质类别名称',
  `zgzsbh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '资格证书编号',
  `zghdsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '资格获得时间',
  `zyzsbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '执业证书编码',
  `fzrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '发证日期',
  `zydd` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '执业地点',
  `zyfw` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '执业范围',
  `zyzyyljgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '主要执业医疗机构代码',
  `zyzyyymc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '主要执业医院名称',
  `qkysbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '全科医生标志',
  `sjhm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '手机号码',
  `zc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '专长',
  `mz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '民族',
  `cjgzrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '参加工作日期',
  `zcsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '注册日期时间',
  `xl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '学历',
  `grzpcfdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '个人照片存放地址',
  `zgzcfdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '资格证存放地址',
  `zyzcfdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '执业证存放地址',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成日期时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_staff_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='人力资源信息';

-- ----------------------------
-- Records of bb_pi_staff
-- ----------------------------
INSERT INTO `bb_pi_staff` VALUES ('1', '0', '123123', '123123', '王小叔', '1', '2022-03-02', '我wwqqq', '01', '', '', '2', '资质类别名称', '123', '2022-03-28', '456', '2022-03-01', '医院', 'A100', '45574403-0', '医院', 'T', '15875310975', '擅长方向', '', '2022-03-03', '2022-03-23', '', '1', '1', '1', '2022-03-29', '1648523641', '1', null, null, null);
INSERT INTO `bb_pi_staff` VALUES ('2', '0', '123123', '123123', '王小叔', '1', '2022-03-02', '我wwqqq', '01', '', '', '2', '资质类别名称', '123', '2022-03-28', '456', '2022-03-01', '医院', 'A100', '45574403-0', '医院', 'T', '15875310975', '擅长方向', '', '2022-03-03', '2022-03-23', '', '1', '1', '1', '2022-03-29', '1648524438', '1', null, null, null);
INSERT INTO `bb_pi_staff` VALUES ('3', '0', '123123', '123123', '王小叔', '1', '2022-03-02', '我wwqqq', '01', '139', '231', '2', '资质类别名称', '123', '2022-03-28', '456', '2022-03-01', '医院', 'A100', '45574403-0', '医院', 'T', '15875310975', '擅长方向', '', '2022-03-03', '2022-03-23', '', '1', '1', '1', '2022-03-29', '1648541978', '1', null, null, null);

-- ----------------------------
-- Table structure for bb_pi_treatment_diagnose
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_treatment_diagnose`;
CREATE TABLE `bb_pi_treatment_diagnose` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识 ',
  `fwwddm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务网点代码',
  `cfbh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方编号',
  `cfmxid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方明细ID',
  `zdxxid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '诊断信息 ID',
  `kh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡号',
  `klx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡类型',
  `mzh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '门诊号',
  `xm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '姓名',
  `xbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '性别',
  `csrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生日期',
  `nls` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄岁',
  `nly` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄月',
  `jzrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '就诊时间',
  `zdlxbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '诊断类型编码',
  `xyzdbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '西医诊断编码',
  `xyzdmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '西医诊断名称',
  `xyzdbmlx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '西医诊断编码类型',
  `zyzdbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中医诊断编码',
  `zyzdmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中医诊断名称',
  `zyzdbmlx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中医诊断编码类型',
  `zdsm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '诊断说明',
  `zdbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '诊断标志',
  `zdysgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '诊断医生工号',
  `zdysxm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '诊断医生姓名',
  `zdsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '诊断时间',
  `yl1` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '预留一',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `mj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密级',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_treatment_diagnose_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='医学诊断';

-- ----------------------------
-- Records of bb_pi_treatment_diagnose
-- ----------------------------
INSERT INTO `bb_pi_treatment_diagnose` VALUES ('1', '0', '121212', '我问鹅鹅鹅', '', '', '70550', '我巍峨巍峨', '9', '61', '我哦', '1', '1991-04-20', '30', '11 16/30', '2022-03-07', '21', '', '', '', 'JB:22.48', '出疹', '05', '', '1', '3', '我哦哦', '2022-03-07', '', '2022-04-06', '1649225544', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_diagnose` VALUES ('2', '0', '121212', '我问鹅鹅鹅', '1111222', '111', '70550', '我巍峨巍峨', '9', '61', '我哦', '1', '1991-04-20', '30', '11 16/30', '2022-03-07', '21', '', '', '', 'JB:22.48', '出疹', '05', '', '1', '3', '我哦哦', '2022-03-07', '', '2022-04-06', '1649226954', '1', '01', '0000-00-00 00:00:00', '2022-04-20 11:40:22', null);
INSERT INTO `bb_pi_treatment_diagnose` VALUES ('3', '0', '121212', '我问鹅鹅鹅', '11', '1', '', '我巍峨巍峨', '', '', '我哦', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '我哦哦', '', '', '', '', '', '', '2022-04-20 11:46:16', '2022-04-20 11:46:16', '2022-04-20 11:46:31');

-- ----------------------------
-- Table structure for bb_pi_treatment_order
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_treatment_order`;
CREATE TABLE `bb_pi_treatment_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识 ',
  `fwwddm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务网点代码',
  `cfbh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方编号',
  `cfmxid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方明细ID',
  `kh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡号',
  `klx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡类型',
  `cfklsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方开立日期',
  `cfyxts` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方有效天数',
  `cfklksbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方科室编码',
  `cfklksmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方科室名称',
  `mzh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '门诊号',
  `xm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '姓名',
  `xbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '性别代码',
  `csrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生日期',
  `nls` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄岁',
  `nly` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄月',
  `jzrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '就诊日期时间',
  `yzsm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医嘱说明',
  `pxh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '排序号',
  `yzxmlxdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医嘱项目类型代码',
  `ypcfsx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药品处方属性',
  `zylbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中药类别代码',
  `cfmxxmbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方明细医保编码',
  `ypid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药监药品ID',
  `cfmxmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方明细名称',
  `ywmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药物名称',
  `ypgg` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药品规格',
  `dddz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'DDD值',
  `ywjxdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药物剂型代码',
  `ywsycjl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药物使用次剂量',
  `ywsyjldw` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药物使用剂量单位',
  `ywsypcdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药物使用频次代码',
  `ywsypcmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药物使用频次名称',
  `yytjdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用药途径代码',
  `yytjmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用药途径名称',
  `ywsyzjl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '药物使用总剂量',
  `cfypzh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方药品组号',
  `zyypcf` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中药饮片处方',
  `zyypjs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中药饮片剂数',
  `zyypjzf` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中药饮片煎煮法',
  `zyyyff` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中药用药方法',
  `fyjl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '发药剂量',
  `fyjldw` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '发药剂量单位',
  `dj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '单价',
  `zje` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '总金额',
  `cfklysgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方开立医师工号',
  `cfklysqm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方开立医师签名',
  `cfshyjsgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方审核工号',
  `cfshyjsqm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方审核签名',
  `cftpyjsgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方调配工号',
  `cftpyjsqm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方调配签名',
  `cfhdyjsgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方核对工号',
  `cfhdyjsqm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方核对签名',
  `cffyyjsgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方发药工号',
  `cffyyjsqm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处方发药签名',
  `zxjg` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '执行结果',
  `bz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '备注',
  `qyjgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '取药机构代码',
  `qyjgmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '取药机构名称',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `mj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密级',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_treatment_order_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='互联网诊疗处方';

-- ----------------------------
-- Records of bb_pi_treatment_order
-- ----------------------------
INSERT INTO `bb_pi_treatment_order` VALUES ('2', '0', '404404532750140543', '404404532750140543', '1646284816818', '123', '372522197406087218', '9', '2022-03-03', '3', '139', '肿瘤科', '60', '谭启省', '1', '1974-06-08', '47', '10 10/30', '2022-03-03', '饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后', '0', '01', '1', '', '', '1120202611DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '03', '10', '瓶', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '12', '3', '于晓黎', '2', '药师1', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249270', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('3', '0', '404404532750140543', '404404532750140543', '1646284816818', '123', '372522197406087218', '9', '2022-03-03', '3', '139', '肿瘤科', '60', '谭启省', '1', '1974-06-08', '47', '10 10/30', '2022-03-03', '嘱托嘱托嘱托嘱托嘱托', '1', '01', '1', '', '', '1120202606DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '01', '11', '克', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '39.6', '3', '于晓黎', '2', '药师1', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249270', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('4', '0', '404404532750140543', '404404532750140543', '1646624677742', '124', '372522197204184239', '9', '2022-03-07', '3', '139', '肿瘤科', '61', '房平东', '2', '1972-04-18', '50', '0 0/30', '2022-03-07', '饭后', '0', '01', '1', '', '', '1120202611DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '03', '10', '克', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '12', '3', '于晓黎', '1', '药师wood', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249270', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('5', '0', '404404532750140543', '404404532750140543', '1646624677742', '124', '372522197204184239', '9', '2022-03-07', '3', '139', '肿瘤科', '61', '房平东', '2', '1972-04-18', '50', '0 0/30', '2022-03-07', '44', '1', '01', '1', '', '', '1120202606DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '01', '11', '克', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '39.6', '3', '于晓黎', '1', '药师wood', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249270', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('6', '0', '404404532750140543', '404404532750140543', '1646284816818', '123', '372522197406087218', '9', '2022-03-03', '3', '139', '肿瘤科', '60', '谭启省', '1', '1974-06-08', '47', '10 10/30', '2022-03-03', '饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后', '', '01', '1', '', '', '1120202611DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '03', '10', '瓶', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '12', '3', '于晓黎', '2', '药师1', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249974', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('7', '0', '404404532750140543', '404404532750140543', '1646284816818', '123', '372522197406087218', '9', '2022-03-03', '3', '139', '肿瘤科', '60', '谭启省', '1', '1974-06-08', '47', '10 10/30', '2022-03-03', '嘱托嘱托嘱托嘱托嘱托', '', '01', '1', '', '', '1120202606DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '01', '11', '克', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '39.6', '3', '于晓黎', '2', '药师1', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249974', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('8', '0', '404404532750140543', '404404532750140543', '1646624677742', '124', '372522197204184239', '9', '2022-03-07', '3', '139', '肿瘤科', '61', '房平东', '2', '1972-04-18', '50', '0 0/30', '2022-03-07', '饭后', '', '01', '1', '', '', '1120202611DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '03', '10', '克', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '12', '3', '于晓黎', '1', '药师wood', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249974', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('9', '0', '404404532750140543', '404404532750140543', '1646624677742', '124', '372522197204184239', '9', '2022-03-07', '3', '139', '肿瘤科', '61', '房平东', '2', '1972-04-18', '50', '0 0/30', '2022-03-07', '44', '', '01', '1', '', '', '1120202606DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '01', '11', '克', '', '', '', '', '', '', '', '', '', '', '', '', '1.20', '39.6', '3', '于晓黎', '1', '药师wood', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650249974', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('10', '0', '404404532750140543', '404404532750140543', '1646284816818', '123', '372522197406087218', '9', '2022-03-03', '3', '139', '肿瘤科', '60', '谭启省', '1', '1974-06-08', '47', '10 10/30', '2022-03-03', '饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后饭后', '', '01', '1', '', '', '1120202611DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '03', '10', '瓶', '', '', '', '', '', '', '', '', '', '', '10', '', '1.20', '12', '3', '于晓黎', '2', '药师1', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650250807', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('11', '0', '404404532750140543', '404404532750140543', '1646284816818', '123', '372522197406087218', '9', '2022-03-03', '3', '139', '肿瘤科', '60', '谭启省', '1', '1974-06-08', '47', '10 10/30', '2022-03-03', '嘱托嘱托嘱托嘱托嘱托', '', '01', '1', '', '', '1120202606DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '01', '11', '克', '', '', '', '', '', '', '', '', '', '', '33', '', '1.20', '39.6', '3', '于晓黎', '2', '药师1', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650250807', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('12', '0', '404404532750140543', '404404532750140543', '1646624677742', '124', '372522197204184239', '9', '2022-03-07', '3', '139', '肿瘤科', '61', '房平东', '2', '1972-04-18', '50', '0 0/30', '2022-03-07', '饭后', '', '01', '1', '', '', '1120202611DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '03', '10', '克', '', '', '', '', '', '', '', '', '', '', '10', '', '1.20', '12', '3', '于晓黎', '1', '药师wood', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650250807', '1', '01', null, null, null);
INSERT INTO `bb_pi_treatment_order` VALUES ('13', '0', '404404532750140543', '404404532750140543', '1646624677742', '124', '372522197204184239', '9', '2022-03-07', '3', '139', '肿瘤科', '61', '房平东', '2', '1972-04-18', '50', '0 0/30', '2022-03-07', '44', '', '01', '1', '', '', '1120202606DD', '葡萄糖氯化钠注射液', '葡萄糖氯化钠注射液', '瓶', '', '01', '11', '克', '', '', '', '', '', '', '', '', '', '', '33', '', '1.20', '39.6', '3', '于晓黎', '1', '药师wood', '', '', '', '', '', '', '', '', '', '', '2022-04-18', '1650250807', '1', '01', null, null, null);

-- ----------------------------
-- Table structure for bb_pi_treatment_record
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_treatment_record`;
CREATE TABLE `bb_pi_treatment_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识',
  `fwwddm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务网点代码',
  `jzlx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '就诊类型',
  `kh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡号',
  `klx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡类型',
  `mzh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '门诊号',
  `ksbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '科室编码',
  `ksmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '科室名称',
  `xm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '姓名',
  `xbdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '性别代码',
  `csrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生日期',
  `nls` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄岁',
  `nly` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄月',
  `gmsbs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '过敏史标识',
  `gms` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '过敏史',
  `cblb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '参保类别',
  `jzrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '就诊日期时间',
  `jzzdsm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '就诊诊断说明',
  `czbzdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '初诊标志代码',
  `zs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '主诉',
  `xbs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '现病史',
  `jws` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '既往史',
  `fzjcxm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '辅助检查项目',
  `fzjcjg` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '辅助检查结果',
  `mzzzzddm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '门诊症状-诊断代码',
  `mzzzzdmc` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '门诊症状-诊断名称',
  `mzzzzdbmlx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '门诊症状诊断编码类型',
  `zzms` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '症状描述',
  `bzyj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '辨证依据',
  `zzzf` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '治则治法',
  `jkzd` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '健康指导',
  `czjh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '处置计划',
  `yzysgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '应诊医生工号',
  `yzysqm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '应诊医师签名',
  `czylwsjgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '初诊医疗卫生机构代码',
  `czylwsjgmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '初诊机构名称',
  `jzjssj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '就诊结束时间',
  `zzbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转诊标志',
  `hzmyd` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '患者满意度',
  `yl1` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '预留一',
  `yl2` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '预留二',
  `dzcfwjcfdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '电子处方地址',
  `ysdspwjcfdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '医生端视频地址',
  `hzdspwjcfdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '患者端视频地址',
  `jlypcfdz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '交流音频地址',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成日期时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `mj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密级',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `idx_bb_pi_treatment_record_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='互联网诊疗病历';

-- ----------------------------
-- Records of bb_pi_treatment_record
-- ----------------------------
INSERT INTO `bb_pi_treatment_record` VALUES ('1', '0', '40441111', '40440453275222', '100', '440781222', '9', '61', '139', '肿瘤科', '李医生', '1', '1991-04-20', '30', '11 16/30', 'T', '3', '', '2022-03-07', '', '1', '1', '', '2', '', '', '', '', '05', '', '', '', '', '', '3', '于晓黎', '404404532750140543', '中山大学附属第八医院（深圳福田）', '2022-03-07', '', '', '', '', 'T', 'F', 'F', 'F', '2022-03-07', '1649226954', '01', '1', null, null, null);

-- ----------------------------
-- Table structure for bb_pi_treatment_referral
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_treatment_referral`;
CREATE TABLE `bb_pi_treatment_referral` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '上传:0未上传 1已上传 2上传失败',
  `jgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '机构标识 ',
  `fwwddm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服务网点代码',
  `zzjlid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转诊记录',
  `kh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡号',
  `klx` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '卡类型',
  `mzh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '门诊号',
  `xm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '姓名',
  `xb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '性别',
  `csrq` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '出生日期',
  `nls` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄岁',
  `nly` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '年龄月',
  `jzrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '就诊时间',
  `fzysgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '负责医生工号',
  `fzysxm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '负责医生姓名',
  `fzksbm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '负责科室编码',
  `fzksmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '负责科室名称',
  `zzyy` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转诊原因',
  `zzsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转诊时间',
  `zxjgdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转向机构代码',
  `zxjgmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转向机构名称',
  `zxksdm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转向科室代码',
  `zxksmc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转向科室名称',
  `zxysgh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转向医生工号',
  `zxysxm` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转向医生姓名',
  `zyjws` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '主要既往史',
  `zljg` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '治疗经过',
  `xybzlfa` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '下一步治疗方案',
  `zzbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转诊标志',
  `sjscsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据生成时间',
  `tbrqsj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '填报日期',
  `mj` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密级',
  `cxbz` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '撤销标志',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_treatment_referral_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='转诊记录';

-- ----------------------------
-- Records of bb_pi_treatment_referral
-- ----------------------------

-- ----------------------------
-- Table structure for bb_pi_up_field
-- ----------------------------
DROP TABLE IF EXISTS `bb_pi_up_field`;
CREATE TABLE `bb_pi_up_field` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否上传:0不传 1上传',
  `table_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '表名',
  `table_name_cn` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '表名cn',
  `table_field` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '字段名',
  `table_field_cn` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '字段名cn',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `tab_field_cn` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字段名cn',
  `tab_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名',
  `tab_name_cn` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名cn',
  `tab_field` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字段名',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`),
  KEY `deleted_at` (`deleted_at`),
  KEY `idx_bb_pi_up_field_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='必传字段表';

-- ----------------------------
-- Records of bb_pi_up_field
-- ----------------------------

-- ----------------------------
-- Table structure for business_examples
-- ----------------------------
DROP TABLE IF EXISTS `business_examples`;
CREATE TABLE `business_examples` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `business_example_field` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '仅作示例条目无实际作用',
  PRIMARY KEY (`id`),
  KEY `idx_business_examples_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of business_examples
-- ----------------------------

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/deleteBbPiTreatmentReferralByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/createBbPiPerson', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/getBbPiInstitutionBusinessList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/updateBbPiInstitutionBusiness', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/findBbPiInstitutionBusiness', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusinessByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusiness', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitutionBusiness/createBbPiInstitutionBusiness', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/superBuilder/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/deleteBbPiTreatmentReferral', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/createBbPiTreatmentReferral', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/getBbPiTreatmentRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/findBbPiTreatmentRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/deleteBbPiTreatmentRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/updateBbPiTreatmentRecord', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/deleteBbPiTreatmentRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentRecord/createBbPiTreatmentRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/getBbPiTreatmentOrderList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/findBbPiTreatmentOrder', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/updateBbPiTreatmentOrder', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/deleteBbPiTreatmentOrderByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/deleteBbPiTreatmentOrder', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentOrder/createBbPiTreatmentOrder', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/getBbPiTreatmentDiagnoseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/findBbPiTreatmentDiagnose', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnoseByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/updateBbPiTreatmentDiagnose', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnose', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentDiagnose/createBbPiTreatmentDiagnose', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/updateBbPiStaff', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/findBbPiStaff', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/getBbPiStaffList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/createBbPiStaff', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/deleteBbPiStaff', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiStaff/deleteBbPiStaffByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/getBbPiServicePointList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/findBbPiServicePoint', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/updateBbPiServicePoint', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/deleteBbPiServicePoint', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/deleteBbPiServicePointByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiServicePoint/createBbPiServicePoint', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/getBbPiPersonList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/findBbPiPerson', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/updateBbPiPerson', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/createBbPiPerson', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/deleteBbPiPerson', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiPerson/deleteBbPiPersonByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/findBbPiInstitutionBusiness', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/getBbPiInstitutionBusinessList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/updateBbPiInstitutionBusiness', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusinessByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusiness', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitutionBusiness/createBbPiInstitutionBusiness', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/getBbPiInstitutionList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/findBbPiInstitution', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/updateBbPiInstitution', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/deleteBbPiInstitutionByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/deleteBbPiInstitution', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiInstitution/createBbPiInstitution', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/findBbPiDevice', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/getBbPiDeviceList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/updateBbPiDevice', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/deleteBbPiDeviceByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/deleteBbPiDevice', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDevice/createBbPiDevice', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/getBbPiDepartmentList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/findBbPiDepartment', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/updateBbPiDepartment', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/deleteBbPiDepartmentByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/deleteBbPiDepartment', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiDepartment/createBbPiDepartment', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/getBbPiContactsList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/updateBbPiContacts', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/findBbPiContacts', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/getBbPiInstitutionList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/findBbPiInstitution', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/updateBbPiInstitution', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/deleteBbPiInstitutionByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/createBbPiInstitution', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiInstitution/deleteBbPiInstitution', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/getBbPiDeviceList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/findBbPiDevice', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/updateBbPiDevice', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/deleteBbPiDeviceByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/createBbPiDevice', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDevice/deleteBbPiDevice', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/getBbPiDepartmentList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/updateBbPiDepartment', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/findBbPiDepartment', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/deleteBbPiDepartment', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/deleteBbPiDepartmentByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiDepartment/createBbPiDepartment', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/getBbPiContactsList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/findBbPiContacts', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/updateBbPiContacts', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/deleteBbPiContactsByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/createBbPiContacts', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiContacts/deleteBbPiContacts', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/getImTxFileList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/findImTxFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/updateImTxFile', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/deleteImTxFileByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/deleteImTxFile', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxFile/createImTxFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/getMemUserSafeList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/findMemUserSafe', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/updateMemUserSafe', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/deleteMemUserSafeByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/deleteMemUserSafe', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUserSafe/createMemUserSafe', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/getMemUserList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/updateMemUser', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/findMemUser', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/deleteMemUserByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/createMemUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/memUser/deleteMemUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/findK8sClusters', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/getK8sClustersList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/deleteK8sClustersByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/updateK8sClusters', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/deleteK8sClusters', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/startOrStopCollect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/k8sClusters/createK8sClusters', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/getImTximList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/updateImTxim', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/findImTxim', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/deleteImTximByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/deleteImTxim', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxim/createImTxim', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/getImTxMsgList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/findImTxMsg', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/updateImTxMsg', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/deleteImTxMsgByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/deleteImTxMsg', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsg/createImTxMsg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/deleteBbPiContactsByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/deleteBbPiContacts', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiContacts/createBbPiContacts', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/getBasicAreaList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/findBasicArea', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/updateBasicArea', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/deleteBasicAreaByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/deleteBasicArea', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicArea/createBasicArea', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/getMemAddressList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/findMemAddress', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/updateMemAddress', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/deleteMemAddressByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/deleteMemAddress', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memAddress/createMemAddress', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/getMemUserLogList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/findMemUserLog', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/updateMemUserLog', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/deleteMemUserLogByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/deleteMemUserLog', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserLog/createMemUserLog', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/getSysSuperBuilderHistoriesList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/findSysSuperBuilderHistories', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/updateSysSuperBuilderHistories', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/deleteSysSuperBuilderHistoriesByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/deleteSysSuperBuilderHistories', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysSuperBuilderHistories/createSysSuperBuilderHistories', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/getImTxFileList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/findImTxFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/getImTxMsgFileList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/findImTxMsgFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/updateImTxMsgFile', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/deleteImTxMsgFileByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/deleteImTxMsgFile', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/imTxMsgFile/createImTxMsgFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/excelList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/getColHsjList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/updateColHsj', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/findColHsj', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/deleteColHsjByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/deleteColHsj', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colHsj/createColHsj', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/startOrStopCollect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/getColCollectList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/updateColCollect', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/findColCollect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/deleteColCollectByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/createColCollect', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colCollect/deleteColCollect', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/findColKeyField', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/getColKeyFieldList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/updateColKeyField', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/deleteColKeyFieldByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/deleteColKeyField', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/colKeyField/createColKeyField', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/getCmsArticleList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/findCmsArticle', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/updateImTxFile', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/deleteImTxFileByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/deleteImTxFile', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxFile/createImTxFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/getMemUserSafeList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/findMemUserSafe', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/updateMemUserSafe', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/deleteMemUserSafeByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/deleteMemUserSafe', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUserSafe/createMemUserSafe', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/getMemUserList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/findMemUser', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/updateMemUser', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/deleteMemUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/deleteMemUserByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/memUser/createMemUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/findK8sPods', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/getK8sPodsList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/updateK8sPods', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/deleteK8sPodsByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/deleteK8sPods', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sPods/createK8sPods', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/findK8sNodes', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/getK8sNodesList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/updateK8sNodes', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/createK8sNodes', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/deleteK8sNodes', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNodes/deleteK8sNodesByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/findK8sNamespaces', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/getK8sNamespacesList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/updateK8sNamespaces', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/deleteK8sNamespacesByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/deleteK8sNamespaces', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sNamespaces/createK8sNamespaces', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/getK8sDeploymentsList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/findK8sDeployments', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/updateK8sDeployments', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/deleteK8sDeploymentsByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/deleteK8sDeployments', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sDeployments/createK8sDeployments', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/updateK8sClusters', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/findK8sClusters', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/getK8sClustersList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/deleteK8sClustersByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/deleteK8sClusters', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/k8sClusters/createK8sClusters', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/startOrStopCollect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/getImTximList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/updateImTxim', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/findImTxim', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/deleteImTximByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/deleteImTxim', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxim/createImTxim', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/getImTxMsgList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/findImTxMsg', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/updateImTxMsg', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/deleteImTxMsgByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/createImTxMsg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsg/deleteImTxMsg', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/getImTxMsgFileList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/findImTxMsgFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/updateImTxMsgFile', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/deleteImTxMsgFileByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/deleteImTxMsgFile', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/imTxMsgFile/createImTxMsgFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/excelList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/getColHsjList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/findColHsj', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/updateColHsj', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/deleteColHsjByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/deleteColHsj', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colHsj/createColHsj', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/startOrStopCollect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/getColCollectList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/findColCollect', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/updateColCollect', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/deleteColCollectByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/deleteColCollect', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colCollect/createColCollect', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/findColKeyField', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/getColKeyFieldList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/updateColKeyField', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/deleteColKeyFieldByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/deleteColKeyField', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/colKeyField/createColKeyField', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/findCmsArticle', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/getCmsArticleList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/updateCmsArticle', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/deleteCmsArticleByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/deleteCmsArticle', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsArticle/createCmsArticle', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/getCmsAdList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/updateCmsAd', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/findCmsAd', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/deleteCmsAd', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/deleteCmsAdByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAd/createCmsAd', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/updateCmsArticle', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/deleteCmsArticle', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/deleteCmsArticleByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsArticle/createCmsArticle', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/getCmsAdList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/findCmsAd', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/updateCmsAd', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/deleteCmsAdByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/deleteCmsAd', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAd/createCmsAd', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/getCmsAdSeatList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/findCmsAdSeat', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/updateCmsAdSeat', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/deleteCmsAdSeatByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/deleteCmsAdSeat', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsAdSeat/createCmsAdSeat', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/commDb/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/commFile/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/getBasicFileList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/findBasicFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/updateBasicFile', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/deleteBasicFileByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/deleteBasicFile', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/basicFile/createBasicFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/getCmsCatList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/findCmsCat', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/updateCmsCat', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/deleteCmsCatByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/deleteCmsCat', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/excel/downloadTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/cmsCat/createCmsCat', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/excel/exportExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/excel/loadExcel', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/email/emailTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/excel/importExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/findCmsAdSeat', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/getCmsAdSeatList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/updateCmsAdSeat', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/deleteCmsAdSeatByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/deleteCmsAdSeat', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsAdSeat/createCmsAdSeat', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/commDb/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/commFile/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/getBasicFileList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/findBasicFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/updateBasicFile', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/deleteBasicFileByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/deleteBasicFile', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/basicFile/createBasicFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/getCmsCatList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/findCmsCat', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/updateCmsCat', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/deleteCmsCatByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/deleteCmsCat', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/cmsCat/createCmsCat', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/downloadTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/exportExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/loadExcel', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/importExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/email/emailTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/superBuilder/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/casbinTest/:pathParam', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/superBuilder/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/casbinTest/:pathParam', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/updateBbPiTreatmentReferral', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/findBbPiTreatmentReferral', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/getBbPiTreatmentReferralList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiTreatmentReferral/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/createBbPiUpField', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/deleteBbPiUpField', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/deleteBbPiUpFieldByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/updateBbPiUpField', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/findBbPiUpField', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/getBbPiUpFieldList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/bbPiUpField/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/deleteBbPiPerson', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/deleteBbPiPersonByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/updateBbPiPerson', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/findBbPiPerson', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/getBbPiPersonList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiPerson/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/createBbPiServicePoint', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/deleteBbPiServicePoint', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/deleteBbPiServicePointByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/updateBbPiServicePoint', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/findBbPiServicePoint', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/getBbPiServicePointList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiServicePoint/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/createBbPiStaff', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/deleteBbPiStaff', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/deleteBbPiStaffByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/updateBbPiStaff', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/findBbPiStaff', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/getBbPiStaffList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiStaff/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/createBbPiTreatmentDiagnose', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnose', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnoseByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/updateBbPiTreatmentDiagnose', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/findBbPiTreatmentDiagnose', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/getBbPiTreatmentDiagnoseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentDiagnose/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/createBbPiTreatmentOrder', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/deleteBbPiTreatmentOrder', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/deleteBbPiTreatmentOrderByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/updateBbPiTreatmentOrder', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/findBbPiTreatmentOrder', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/getBbPiTreatmentOrderList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentOrder/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/createBbPiTreatmentRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/deleteBbPiTreatmentRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/deleteBbPiTreatmentRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/updateBbPiTreatmentRecord', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/findBbPiTreatmentRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/getBbPiTreatmentRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentRecord/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/createBbPiTreatmentReferral', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/deleteBbPiTreatmentReferral', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/deleteBbPiTreatmentReferralByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/updateBbPiTreatmentReferral', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/findBbPiTreatmentReferral', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/getBbPiTreatmentReferralList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiTreatmentReferral/excelList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/createBbPiUpField', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/deleteBbPiUpField', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/deleteBbPiUpFieldByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/updateBbPiUpField', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/findBbPiUpField', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/getBbPiUpFieldList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/quickEdit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/bbPiUpField/excelList', 'GET', '', '', '');

-- ----------------------------
-- Table structure for cms_ad
-- ----------------------------
DROP TABLE IF EXISTS `cms_ad`;
CREATE TABLE `cms_ad` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '广告id',
  `seat_id` int(11) unsigned DEFAULT '0' COMMENT '广告位置',
  `name` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '广告名称',
  `media_type` smallint(2) unsigned DEFAULT '0' COMMENT '广告类型:0=文本，1=图片，2=视频，3=flash，4=音频',
  `be_target` tinyint(1) DEFAULT '0' COMMENT '是否新窗: 0=新窗体 1=本页',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '广告图片',
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '广告链接',
  `detail` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '广告内容',
  `exp_detail` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '过期内容',
  `start_time` datetime DEFAULT NULL COMMENT '投放时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `total_click` int(11) unsigned DEFAULT '0' COMMENT '点击量',
  `sort` int(11) DEFAULT '50' COMMENT '排序',
  `status` smallint(2) unsigned DEFAULT '0' COMMENT '状态:0未审核 1审核 2未通过审核 3 草稿',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_seat_id` (`seat_id`),
  KEY `idx_cms_ad_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='广告表';

-- ----------------------------
-- Records of cms_ad
-- ----------------------------
INSERT INTO `cms_ad` VALUES ('4', '0', '3333', '3', '0', '', '', '3333', '33', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '2323', '2323', '1', '2021-07-09 16:36:57', '2021-07-19 13:51:20', '2021-07-09 16:36:57');
INSERT INTO `cms_ad` VALUES ('5', '23423', '4234234', '3', '1', '', '', '23423', '4234234234', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '234', '234', '1', '2021-07-09 16:37:20', '2021-07-19 13:51:20', '2021-07-16 15:21:45');
INSERT INTO `cms_ad` VALUES ('6', '2322', '232', '232', '0', '', '', '232', '3232', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '3223', '232', '2', '2021-07-16 15:22:03', '2021-07-19 15:28:09', '2021-07-16 15:22:03');
INSERT INTO `cms_ad` VALUES ('7', '2222', '2222', '3', '1', '', '', '222', '22', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '222', '22', '2', '2021-07-19 13:51:03', '2021-07-19 15:28:09', '2021-07-19 14:06:44');
INSERT INTO `cms_ad` VALUES ('8', '11', '111', '1', '1', '20a6e4a766d14c30ad79933f0d940600', '', '<p>1<img src=\"http://localhost:40050/res/sys/0/20211020/d3c7afed79404c8db00cf22ca1b8e575.jpg\" alt=\"Smiley face\" width=\"467\" height=\"467\" />11汪汪汪<img src=\"http://localhost:40050/res/sys/0/20211020/4adbb2092f754b07b7e63117d5d84ebd_src.png\" />wwwww汪<img src=\"http://localhost:8888/res/sys/0/20211014/5a4fead22ec44c4398ded2525d7b8956.jpg\" alt=\"Smiley face\" width=\"80\" height=\"80\" /></p>', '<p>11呜呜呜<img src=\"http://localhost:40050/res/sys/0/20211020/3999f6b68f794239a5a30f7113028530.jpg\" alt=\"Smiley face\" width=\"653\" height=\"653\" />www</p>', '2021-12-20 13:29:36', '2021-12-24 13:29:38', '22', '22', '3', '0000-00-00 00:00:00', null, '2021-12-20 17:15:40');
INSERT INTO `cms_ad` VALUES ('9', '222222', '222222222', '1', '1', '1edff067da1e498a8ccd15d6226198ee', '1111', '<p>222222</p>', '<p>2222</p>', '2022-03-03 08:06:43', '2021-12-17 10:01:01', '111', '11', '1', '0000-00-00 00:00:00', null, '2021-12-20 17:15:15');
INSERT INTO `cms_ad` VALUES ('10', null, '', null, '0', 'e8fb941c4b544cbb887637ae5fb24445', '', '<p><img src=\"../res/sys/0/20211021/85e3ab6bfa3447bebd7365c6a527c971_src.png\" /><img src=\"../res/sys/0/20211021/ef2a204318a54e79875e89e0b74f9c43_src.jpg\" alt=\"\" width=\"520\" height=\"801\" /></p>', '<p>213123123</p>\n<p><img src=\"../res/sys/0/20211021/06d392e31927424992f1bd00c2db859e.jpg\" alt=\"Smiley face\" width=\"219\" height=\"219\" /><img src=\"../res/sys/0/20211021/06d392e31927424992f1bd00c2db859e.jpg\" alt=\"Smiley face\" width=\"551\" height=\"551\" /></p>', null, null, null, null, null, '2021-10-21 10:46:52', '2021-11-04 14:38:34', '2021-10-21 10:46:52');

-- ----------------------------
-- Table structure for cms_ad_seat
-- ----------------------------
DROP TABLE IF EXISTS `cms_ad_seat`;
CREATE TABLE `cms_ad_seat` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '位置id',
  `name` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '位置名称',
  `width` smallint(5) unsigned DEFAULT '0' COMMENT '广告位宽度',
  `height` smallint(5) unsigned DEFAULT '0' COMMENT '广告位高度',
  `desc` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '广告描述',
  `style` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '模板',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0关闭1开启',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_cms_ad_seat_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COMMENT='广告位置表';

-- ----------------------------
-- Records of cms_ad_seat
-- ----------------------------
INSERT INTO `cms_ad_seat` VALUES ('1', '111', '22', '1212', '121212', '123123', '0', null, '2021-11-08 10:41:04', null);
INSERT INTO `cms_ad_seat` VALUES ('2', '777733333', '65535', '777', '77', '77', '3', null, '2021-11-08 10:41:40', null);
INSERT INTO `cms_ad_seat` VALUES ('3', '3434', '65535', '0', '<p>234<img src=\"http://localhost:40050/res/sys/0/20211220/1edff067da1e498a8ccd15d6226198ee.png\" alt=\"Smiley face\" width=\"80\" height=\"80\" /></p>', '234234234', '2', '2021-12-21 10:59:52', null, '0000-00-00 00:00:00');
INSERT INTO `cms_ad_seat` VALUES ('4', '11111', '0', '0', '<p>112222</p>', '1111', '0', '2021-11-08 10:40:30', '2021-11-08 10:42:14', '0000-00-00 00:00:00');
 
-- ----------------------------
-- Table structure for cms_article
-- ----------------------------
DROP TABLE IF EXISTS `cms_article`;
CREATE TABLE `cms_article` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `cat_id` int(11) DEFAULT '0' COMMENT '类别ID',
  `cat_id_sys` int(11) DEFAULT '0' COMMENT '系统类别ID',
  `media_type` smallint(1) DEFAULT '0' COMMENT '文章类型',
  `name` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文章标题',
  `sketch` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文章摘要',
  `detail` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '文章内容:(以后需要分表)',
  `author` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '文章作者',
  `tag_list` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '标签列表:id1,id2,id3,',
  `author_email` varchar(60) DEFAULT '' COMMENT '作者邮箱',
  `referer` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '来源',
  `thumb` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '插图',
  `qrcode` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '二维码图片',
  `image_list` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '图片列表:id列 1,2,3',
  `media_list` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '媒体列表:id列 1,2,3',
  `link` varchar(255) DEFAULT '' COMMENT '链接地址',
  `is_top` tinyint(1) DEFAULT '0' COMMENT '置顶:0否 1是）',
  `is_hot` tinyint(1) DEFAULT '0' COMMENT '热门:(0否  1是)',
  `total_discuss` int(11) DEFAULT '0' COMMENT '总评论',
  `total_click` int(11) DEFAULT '0' COMMENT '总点击',
  `total_star` int(11) DEFAULT '0' COMMENT '总评',
  `total_star_1` bigint(20) DEFAULT NULL COMMENT '总星评1',
  `total_star_2` bigint(20) DEFAULT NULL COMMENT '总星评2',
  `total_star_3` bigint(20) DEFAULT NULL COMMENT '总星评3',
  `total_star_4` bigint(20) DEFAULT NULL COMMENT '总星评4',
  `total_star_5` bigint(20) DEFAULT NULL COMMENT '总星评5',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `pid` bigint(20) DEFAULT '0' COMMENT '父id:章节集合',
  `chapter` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '章节集合:json_id',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未审核 1审核  2未通过审核 3 草稿',
  `verify_msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '审核信息',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_cat_id` (`cat_id`),
  KEY `idx_total_click` (`total_click`),
  KEY `idx_cat_id_sys` (`cat_id_sys`),
  KEY `idx_type` (`media_type`),
  KEY `idx_status` (`status`),
  KEY `idx_is_hot` (`is_hot`),
  KEY `idx_total_star` (`total_star`) USING BTREE,
  KEY `idx_sort` (`sort`) USING BTREE,
  KEY `idx_userid` (`user_id`) USING BTREE,
  KEY `idx_is_top` (`is_top`) USING BTREE,
  KEY `idx_cms_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2578 DEFAULT CHARSET=utf8 COMMENT='文章表';

-- ----------------------------
-- Records of cms_article
-- ----------------------------
INSERT INTO `cms_article` VALUES ('1', '0', '20', '0', '0', '指上篮球', '指上篮球', '<h4 style=\"margin: 1em 0px 16px; padding: 0px; font-size: 1.25em; max-width: 100%; box-sizing: border-box; color: #333333; letter-spacing: 0.544px; text-align: justify; white-space: normal; background-color: #ffffff; font-family: \'Microsoft YaHei\', Helvetica, \'Meiryo UI\', \'Malgun Gothic\', \'Segoe UI\', \'Trebuchet MS\', Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, \'Helvetica Neue\', \'Droid Sans\', \'wenquanyi micro hei\', FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; line-height: 1.4; overflow-wrap: break-word !important;\"><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box !important; overflow-wrap: break-word !important; color: #ff4c41;\">指上篮球</span></h4>\n<p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: #333333; letter-spacing: 0.544px; text-align: justify; white-space: normal; background-color: #ffffff; cursor: default; font-family: \'Microsoft YaHei\', Helvetica, \'Meiryo UI\', \'Malgun Gothic\', \'Segoe UI\', \'Trebuchet MS\', Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, \'Helvetica Neue\', \'Droid Sans\', \'wenquanyi micro hei\', FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; overflow-wrap: break-word !important;\"><img src=\"http://localhost:8888/res/sys/0/20211014/4adecad6d9274ed8b10fc3ec1e2724f9.jpg\" alt=\"Smiley face\" width=\"80\" height=\"80\" />33333</p>\n<p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: #333333; letter-spacing: 0.544px; white-space: normal; background-color: #ffffff; cursor: default; font-family: \'Microsoft YaHei\', Helvetica, \'Meiryo UI\', \'Malgun Gothic\', \'Segoe UI\', \'Trebuchet MS\', Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, \'Helvetica Neue\', \'Droid Sans\', \'wenquanyi micro hei\', FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; text-align: center; overflow-wrap: break-word !important;\">《指上篮球》游戏截图</p>\n<p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: #333333; letter-spacing: 0.544px; text-align: justify; white-space: normal; background-color: #ffffff; cursor: default; font-family: \'Microsoft YaHei\', Helvetica, \'Meiryo UI\', \'Malgun Gothic\', \'Segoe UI\', \'Trebuchet MS\', Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, \'Helvetica Neue\', \'Droid Sans\', \'wenquanyi micro hei\', FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; overflow-wrap: break-word !important;\"><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box !important; overflow-wrap: break-word !important; font-size: 16px;\"><strong style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box; overflow-wrap: break-word !important;\">小编简评</strong>：</span></p>\n<p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: #333333; letter-spacing: 0.544px; text-align: justify; white-space: normal; background-color: #ffffff; cursor: default; font-family: \'Microsoft YaHei\', Helvetica, \'Meiryo UI\', \'Malgun Gothic\', \'Segoe UI\', \'Trebuchet MS\', Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, \'Helvetica Neue\', \'Droid Sans\', \'wenquanyi micro hei\', FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; text-indent: 2em; overflow-wrap: break-word !important;\"><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box !important; overflow-wrap: break-word !important; font-size: 16px;\">《指上篮球》&nbsp; 是一款模拟投&nbsp; 篮的休闲小游戏，玩家扮演的不是投手而是球篮，玩的过程玩家要面对各种刁钻的投篮方向，变幻莫测。游戏看上去比较简单，还是有一些可玩性的。&nbsp; &nbsp;推荐大家体验。&nbsp;</span></p>\n<p>&nbsp;</p>', '猪八戒小游戏', '', '', '猪八戒小游戏', '71391d7c8e3643ca8bd226e40409079f', '0', '', '', '', '1', '0', '0', '50', '0', '0', '0', '0', '0', '0', '0', '0', '', '1', '', '0000-00-00 00:00:00', null, '2021-10-14 23:44:26');
INSERT INTO `cms_article` VALUES ('2', '0', '20', '0', '0', '陀螺大乱斗', '陀螺大乱斗', '<p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: rgb(51, 51, 51); letter-spacing: 0.544px; text-align: justify; white-space: normal; background-color: rgb(255, 255, 255); cursor: default; font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; overflow-wrap: break-word !important;\"><img class=\"\" src=\"https://mmbiz.qpic.cn/mmbiz_png/r3Uu3clxncYU7qRrrQBKYocHvxQ1tX4GglQB4rn9rhskfvLuUYKhsgAqHU6vDJnjYCjCxTOXSmMVWLrHpYU5Vg/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1\"/></p><p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: rgb(51, 51, 51); letter-spacing: 0.544px; white-space: normal; background-color: rgb(255, 255, 255); cursor: default; font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; text-align: center; overflow-wrap: break-word !important;\">《陀螺大乱斗》游戏截图</p><p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: rgb(51, 51, 51); letter-spacing: 0.544px; text-align: justify; white-space: normal; background-color: rgb(255, 255, 255); cursor: default; font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; overflow-wrap: break-word !important;\"><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box !important; overflow-wrap: break-word !important; font-size: 16px;\"><strong style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box; overflow-wrap: break-word !important;\">小编简评</strong>：</span></p><p style=\"margin-top: 0px; margin-bottom: 16px; padding: 0px; max-width: 100%; box-sizing: border-box; clear: both; min-height: 1em; color: rgb(51, 51, 51); letter-spacing: 0.544px; text-align: justify; white-space: normal; background-color: rgb(255, 255, 255); cursor: default; font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; font-size: 14px; text-indent: 2em; overflow-wrap: break-word !important;\"><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box !important; overflow-wrap: break-word !important; font-size: 16px;\">在体验了众多碰碰对撞的游戏，小编认为《陀螺大乱斗》是最有特色的对撞游戏之一，与其它游戏直线对撞不同，该游戏的运动速度更快，方向的灵敏度更高，所以在游戏中可以采用陀螺的曲线式运动轨迹进行对撞，玩家结合曲线的运动规律会更好的赢得比赛</span></p><p><br/></p>', '猪八戒小游戏', '', '', '猪八戒小游戏', '15', '0', '', '', '', '1', '0', '0', '33', '0', '0', '0', '0', '0', '0', '0', '0', '', '1', null, null, null, null);
INSERT INTO `cms_article` VALUES ('3', '0', '20', '0', '0', '吃鸡碰碰车', '吃鸡碰碰车', '<p><span style=\"color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255);\">童年中游乐场中的碰碰车，也在小游戏中火爆起来，小编玩了几十款2D与3D的碰碰对撞游戏，其中优秀的LayaAir 3D碰碰游戏很多，比如《最强碰一碰OL》、《</span><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box; color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255); overflow-wrap: break-word !important;\">球球碰碰车</span><span style=\"color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255);\">》、《</span><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box; color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255); overflow-wrap: break-word !important;\">萌宠碰一碰</span><span style=\"color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255);\">》、《</span><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box; color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255); overflow-wrap: break-word !important;\">球球撞撞撞</span><span style=\"color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255);\">》、《</span><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box; color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255); overflow-wrap: break-word !important;\">碰撞球球</span><span style=\"color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255);\">》、《</span><span style=\"margin: 0px; padding: 0px; max-width: 100%; box-sizing: border-box; color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255); overflow-wrap: break-word !important;\">碰碰车大乱斗</span><span style=\"color: rgb(51, 51, 51); font-family: &quot;Microsoft YaHei&quot;, Helvetica, &quot;Meiryo UI&quot;, &quot;Malgun Gothic&quot;, &quot;Segoe UI&quot;, &quot;Trebuchet MS&quot;, Monaco, monospace, Tahoma, STXihei, 华文细黑, STHeiti, &quot;Helvetica Neue&quot;, &quot;Droid Sans&quot;, &quot;wenquanyi micro hei&quot;, FreeSans, Arimo, Arial, SimSun, 宋体, Heiti, 黑体, sans-serif; letter-spacing: 0.544px; text-align: justify; text-indent: 28px; background-color: rgb(255, 255, 255);\">》等等。《吃鸡碰碰车》相对于其它碰碰车游戏，采用了真联网对战、随机地图、游戏地型丰富多样等特色亮点。可玩性较高。<img src=\"/uploads/images/20190422/dca02d9d2206dfab5c1700ddf58dad31.png\" title=\"微信截图_20190422164259.png\" alt=\"\"/></span></p>', '猪八戒小游戏', '', '', '猪八戒小游戏', '14', '0', '', '', '', '1', '0', '0', '41', '0', '0', '0', '0', '0', '0', '0', '0', '', '0', null, '2021-07-21 00:00:00', null, null);
INSERT INTO `cms_article` VALUES ('4', '0', '20', '0', '0', '萌宠跑酷', '萌宠跑酷', '<p><span style=\"color: rgb(34, 34, 34); font-family: Consolas, &quot;Lucida Console&quot;, &quot;Courier New&quot;, monospace; font-size: 12px; white-space: pre-wrap; background-color: rgb(255, 255, 255);\">《萌宠跑酷》<img src=\"/uploads/images/20190422/89b4560419b6d4d81e6655c911e35f84.jpg\" title=\"萌宠跑酷.jpg\" alt=\"\"/></span></p><p><img src=\"/uploads/images/20190422/89b4560419b6d4d81e6655c911e35f84.jpg\" title=\"萌宠跑酷.jpg\"/></p><p><img src=\"/uploads/images/20190422/40b9d7d985140a868e07bce3c7903104.jpg\" title=\"萌宠跑酷QR.jpg\"/></p><p><span style=\"color: rgb(34, 34, 34); font-family: Consolas, &quot;Lucida Console&quot;, &quot;Courier New&quot;, monospace; font-size: 12px; white-space: pre-wrap; background-color: rgb(255, 255, 255);\"></span><br/></p>', '猪八戒小游戏', '', '', '猪八戒小游戏', '13', '0', '', '', '', '1', '0', '3', '36', '2', '0', '0', '0', '0', '0', '0', '0', '', '0', null, '2021-08-01 00:00:00', null, null);
 
 
-- ----------------------------
-- Table structure for cms_cat
-- ----------------------------
DROP TABLE IF EXISTS `cms_cat`;
CREATE TABLE `cms_cat` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `be_sys` tinyint(1) DEFAULT '0' COMMENT '系统分类: 0=否 1=是',
  `group_id` bigint(20) DEFAULT '0' COMMENT '群组id',
  `media_type` smallint(2) DEFAULT '0' COMMENT '文章类型',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '名称:（栏目/专题）',
  `thumb` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '配图:(可以多个 guid )',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `be_nav` tinyint(1) DEFAULT '0' COMMENT '是否导航',
  `desc` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '描述',
  `keywords` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '关键词',
  `alias` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '别名',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未审核 1审核  2未通过审核 3 草稿',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_cms_cat_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8 COMMENT='文章栏目表 ';

-- ----------------------------
-- Records of cms_cat
-- ----------------------------
INSERT INTO `cms_cat` VALUES ('1', '1', '1', '33', '1', '此3333', '4d68b852efdc438c8f00d680e34ca0cb', '33', '1', 'H5小游戏33', 'H5小游戏333', 'H5小游戏333', '1', '2021-10-12 11:30:22', '0000-00-00 00:00:00', null);
INSERT INTO `cms_cat` VALUES ('2', '0', '1', '0', '6', '漫画111', '7b36124f03934584b6f9401fc4ff2c53', '0', '1', '漫画', '漫画', '漫画', '0', '2021-10-08 23:33:19', '0000-00-00 00:00:00', null);
INSERT INTO `cms_cat` VALUES ('3', '1', '0', '0', '6', '动漫2222', 'd8693ad33267461cb6dba290cbb8c789', '0', '1', '动漫222', '动漫22', '动漫333', '3', '2021-10-09 11:17:07', '0000-00-00 00:00:00', null);
INSERT INTO `cms_cat` VALUES ('4', '0', '0', '0', '6', '游戏资讯', '7b36124f03934584b6f9401fc4ff2c53', '0', '0', '游戏资讯', '新闻资讯', '新闻资讯', '3', '2021-10-09 11:17:15', '0000-00-00 00:00:00', null);
INSERT INTO `cms_cat` VALUES ('5', '0', '1', '0', '3', '传统游戏333', 'c00b7c6993ed4a11a1352ca984823c9a', '0', '0', '传统游戏', '传统游戏', '传统游戏', '3', '2021-10-09 11:17:25', '0000-00-00 00:00:00', null);
 
-- ----------------------------
-- Table structure for cms_cat_art
-- ----------------------------
DROP TABLE IF EXISTS `cms_cat_art`;
CREATE TABLE `cms_cat_art` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  `cat_id` int(11) DEFAULT '0' COMMENT '栏目id',
  `article_id` bigint(20) DEFAULT '0' COMMENT '文章id',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未审核 1审核 2未通过审核 3草稿',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cat_id` (`cat_id`),
  KEY `article_id_` (`article_id`),
  KEY `idx_cms_cat_art_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='栏目与文章关系表';

-- ----------------------------
-- Records of cms_cat_art
-- ----------------------------

-- ----------------------------
-- Table structure for cms_discuss
-- ----------------------------
DROP TABLE IF EXISTS `cms_discuss`;
CREATE TABLE `cms_discuss` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'reply_id',
  `pid` bigint(20) DEFAULT '0' COMMENT '父id:被回复id',
  `article_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `reply_user_id` bigint(20) unsigned DEFAULT '0' COMMENT '被回复用户id',
  `praise` tinyint(1) DEFAULT '0' COMMENT '点赞 0 评论 1 点赞 2 取消点赞',
  `detail` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '内容',
  `status` smallint(2) DEFAULT '1' COMMENT '状态：0未审核 1审核 2未通过审核 3 草稿',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章评论表';

-- ----------------------------
-- Records of cms_discuss
-- ----------------------------

-- ----------------------------
-- Table structure for cms_group
-- ----------------------------
DROP TABLE IF EXISTS `cms_group`;
CREATE TABLE `cms_group` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `type` int(11) DEFAULT '0' COMMENT '群组类别',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(100) DEFAULT '' COMMENT '名称',
  `detail` varchar(2000) DEFAULT '' COMMENT '内容',
  `thumb` varchar(255) DEFAULT '' COMMENT '图片',
  `tag_list` varchar(500) DEFAULT '' COMMENT '标签列表 ,id1,id2,id3',
  `image_list` varchar(500) DEFAULT '' COMMENT '图片列表（id列 1,2,3）',
  `media_list` varchar(500) DEFAULT '' COMMENT '媒体列表（id列 1,2,3）',
  `status` smallint(2) DEFAULT '1' COMMENT '状态：0未审核 1审核 2未通过审核 3 草稿',
  PRIMARY KEY (`id`),
  KEY `idx_userid` (`user_id`),
  KEY `idx_type` (`type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='群组圈子表';

-- ----------------------------
-- Records of cms_group
-- ----------------------------

-- ----------------------------
-- Table structure for cms_tag
-- ----------------------------
DROP TABLE IF EXISTS `cms_tag`;
CREATE TABLE `cms_tag` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '标签名称',
  `status` smallint(2) DEFAULT '0' COMMENT '状态：0未审核 1审核 2未通过审核 3 草稿',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_order` (`sort`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签表';

-- ----------------------------
-- Records of cms_tag
-- ----------------------------

-- ----------------------------
-- Table structure for cms_visitor
-- ----------------------------
DROP TABLE IF EXISTS `cms_visitor`;
CREATE TABLE `cms_visitor` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '访客id',
  `article_id` bigint(20) unsigned NOT NULL COMMENT '文章id',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  `plat_type` smallint(2) DEFAULT '0' COMMENT '客户平台:web,ios,andriod,weixin',
  `client_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '客户端ip',
  PRIMARY KEY (`id`),
  KEY `idx_cat_id` (`article_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章访客表';

-- ----------------------------
-- Records of cms_visitor
-- ----------------------------

-- ----------------------------
-- Table structure for col_collect
-- ----------------------------
DROP TABLE IF EXISTS `col_collect`;
CREATE TABLE `col_collect` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `userid` int(10) DEFAULT '0' COMMENT 'id',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '名称',
  `url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '路径',
  `url_page` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '翻页路径',
  `to_table` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应表',
  `now_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '当前id',
  `page_now` int(10) DEFAULT '0' COMMENT '当前页码',
  `page_start` int(10) DEFAULT '0' COMMENT '开始页码',
  `page_end` int(10) DEFAULT '0' COMMENT '结束页码',
  `status_run` smallint(2) DEFAULT '0' COMMENT '运行状态:0停止 1采集中',
  `run_total` int(10) DEFAULT '0' COMMENT '运行次数',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未审核 1有效  2未通过审核 3草稿',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_col_collect_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='采集表';

-- ----------------------------
-- Records of col_collect
-- ----------------------------
INSERT INTO `col_collect` VALUES ('1', '1', '长江海事局', 'http://www.cjhdj.com.cn/hdfw/hdtg/index.shtml', 'http://www.cjhdj.com.cn/hdfw/hdtg/index_%d.shtml', '', '', null, '0', '20', '0', null, '1', '0000-00-00 00:00:00', null, '2022-04-21 13:53:41');
INSERT INTO `col_collect` VALUES ('2', '2', '22', '22', '2', '', '', '0', '0', '0', '0', '0', '0', '2021-12-23 14:03:24', '2021-12-23 14:03:28', '2021-12-23 14:03:24');

-- ----------------------------
-- Table structure for col_hsj
-- ----------------------------
DROP TABLE IF EXISTS `col_hsj`;
CREATE TABLE `col_hsj` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `col_id` int(11) unsigned DEFAULT '0' COMMENT '采集id',
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
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_col_hsj_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2612 DEFAULT CHARSET=utf8 COMMENT='海事信息内容表';

-- ----------------------------
-- Records of col_hsj
-- ----------------------------

-- ----------------------------
-- Table structure for col_key_field
-- ----------------------------
DROP TABLE IF EXISTS `col_key_field`;
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
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_col_key_field_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='采集标识字段对应表';

-- ----------------------------
-- Records of col_key_field
-- ----------------------------

-- ----------------------------
-- Table structure for im_tx_file
-- ----------------------------
DROP TABLE IF EXISTS `im_tx_file`;
CREATE TABLE `im_tx_file` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `msg_id` bigint(20) unsigned NOT NULL COMMENT '消息id',
  `media_type` smallint(2) DEFAULT NULL COMMENT '文件类型',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '远程地址',
  `local` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '本地路径',
  `status` smallint(2) DEFAULT '0' COMMENT '下载状态',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_im_tx_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8 COMMENT='聊天文件表';

-- ----------------------------
-- Records of im_tx_file
-- ----------------------------
INSERT INTO `im_tx_file` VALUES ('90', '1775', null, 'https://hlwyy.sysu8h.com.cn:7443/uploads/store/comment/20220324/31503ed11cb5fef0bb047f819003fa55.png', 'res/sys/txim/202203/2517/31503ed11cb5fef0bb047f819003fa55.png', '1', '2022-03-28 16:10:21', '2022-03-28 16:10:21', null);
INSERT INTO `im_tx_file` VALUES ('91', '1781', null, 'https://hlwyy.sysu8h.com.cn:7443/uploads/store/comment/20220324/31503ed11cb5fef0bb047f819003fa55.png', 'res/sys/txim/202203/2518/31503ed11cb5fef0bb047f819003fa55.png', '1', '2022-03-28 16:10:41', '2022-03-28 16:10:40', null);

-- ----------------------------
-- Table structure for im_tx_msg
-- ----------------------------
DROP TABLE IF EXISTS `im_tx_msg`;
CREATE TABLE `im_tx_msg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `chat_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '消息类型',
  `msg_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '消息时间',
  `from_account` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '发送人',
  `to_account` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '接收人',
  `msg_timestamp` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '时间撮',
  `msg_seq` bigint(10) DEFAULT '0' COMMENT 'seq',
  `msg_random` bigint(20) DEFAULT NULL COMMENT 'random',
  `msg_type` varchar(100) DEFAULT '' COMMENT '消息类型',
  `msg_text` varchar(4000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '文字',
  `msg_content` text COMMENT '内容',
  `media_list` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '媒体文件',
  `media_list_tx` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '媒体远程文件',
  `status_media` smallint(2) DEFAULT '0' COMMENT '媒体下载:0未下载,1已下载,2 失败',
  `client_ip` varchar(255) DEFAULT NULL COMMENT 'IP地址',
  `msg_from_platform` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '平台',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未上传,1已上传,2 失败',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `order_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '订单id',
  `order_status` smallint(2) DEFAULT '0' COMMENT '订单状态   1开始 2 结束',
  PRIMARY KEY (`id`),
  KEY `idx_im_tx_msg_deleted_at` (`deleted_at`),
  KEY `order_id` (`order_id`),
  KEY `order_status` (`order_status`),
  KEY `from_account` (`from_account`),
  KEY `to_account` (`to_account`),
  KEY `msg_time` (`msg_time`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1787 DEFAULT CHARSET=utf8 COMMENT='消息表';

-- ----------------------------
-- Records of im_tx_msg
-- ----------------------------
INSERT INTO `im_tx_msg` VALUES ('1775', 'C2C', '2022032517', '17153', '16796', '2022-03-24 11:34:59', '1135070002', '91255596', 'TIMCustomElem', '', '{\"Text\":\"\",\"Desc\":\"图片合集\",\"Data\":\"imagesgroup\",\"Ext\":\"{\\\"images\\\":[\\\"https://.com.cn:7443/uploads/store/comment/20220324/31503ed11cb5fef0bb047f819003fa55.png\\\"]}\"}', '', '[\"https:// :7443/uploads/store/comment/20220324/31503ed11cb5fef0bb047f819003fa55.png\"]', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:20', '2022-03-28 16:10:20', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1776', 'C2C', '2022032517', '16796', '17153', '2022-03-24 11:34:32', '1131230003', '76867314', 'TIMCustomElem', '', '{\"Text\":\"\",\"Desc\":\"\",\"Data\":\"systemmessage\",\"Ext\":\"{\\\"showDot\\\":false,\\\"content\\\":\\\"—以上是上次问诊内容—\\\"}\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:21', '2022-03-28 16:10:21', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1777', 'C2C', '2022032517', '16796', '17153', '2022-03-24 11:34:38', '1131230004', '99922577', 'TIMTextElem', '你和我', '{\"Text\":\"你和我\",\"Desc\":\"\",\"Data\":\"\",\"Ext\":\"\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:21', '2022-03-28 16:10:21', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1778', 'C2C', '2022032517', '17153', '16796', '2022-03-24 11:34:59', '1135070001', '64687373', 'TIMCustomElem', '', '{\"Text\":\"\",\"Desc\":\"初诊信息\",\"Data\":\"patientcard\",\"Ext\":\"[{\\\"title\\\":\\\"患者信息\\\",\\\"text\\\":\\\"李小兵 男 48岁\\\",\\\"id\\\":\\\"wx2382634507808276481648092907138\\\"},{\\\"title\\\":\\\"病情描述\\\",\\\"text\\\":\\\"111\\\",\\\"id\\\":\\\"wx238263450780827648\\\"}]\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:21', '2022-03-28 16:10:21', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1779', 'C2C', '2022032517', '17153', '16796', '2022-03-24 11:36:26', '1136340001', '47452544', 'TIMTextElem', '11', '{\"Text\":\"11\",\"Desc\":\"\",\"Data\":\"\",\"Ext\":\"\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:21', '2022-03-28 16:10:21', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1780', 'C2C', '2022032517', '17153', '16796', '2022-03-24 11:45:07', '1145150001', '11188110', 'TIMTextElem', '[敲打]', '{\"Text\":\"[敲打]\",\"Desc\":\"\",\"Data\":\"\",\"Ext\":\"\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:21', '2022-03-28 16:10:21', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1781', 'C2C', '2022032518', '17153', '16796', '2022-03-24 11:34:59', '1135070002', '91255596', 'TIMCustomElem', '', '{\"Text\":\"\",\"Desc\":\"图片合集\",\"Data\":\"imagesgroup\",\"Ext\":\"{\\\"images\\\":[\\\" \\\"]}\"}', '', '[\" \"]', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:40', '2022-03-28 16:10:40', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1782', 'C2C', '2022032518', '16796', '17153', '2022-03-24 11:34:32', '1131230003', '76867314', 'TIMCustomElem', '', '{\"Text\":\"\",\"Desc\":\"\",\"Data\":\"systemmessage\",\"Ext\":\"{\\\"showDot\\\":false,\\\"content\\\":\\\"—以上是上次问诊内容—\\\"}\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:40', '2022-03-28 16:10:40', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1783', 'C2C', '2022032518', '16796', '17153', '2022-03-24 11:34:38', '1131230004', '99922577', 'TIMTextElem', '你和我', '{\"Text\":\"你和我\",\"Desc\":\"\",\"Data\":\"\",\"Ext\":\"\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:41', '2022-03-28 16:10:41', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1784', 'C2C', '2022032518', '17153', '16796', '2022-03-24 11:34:59', '1135070001', '64687373', 'TIMCustomElem', '', '{\"Text\":\"\",\"Desc\":\"初诊信息\",\"Data\":\"patientcard\",\"Ext\":\"[{\\\"title\\\":\\\"患者信息\\\",\\\"text\\\":\\\"李小兵 男 48岁\\\",\\\"id\\\":\\\"wx2382634507808276481648092907138\\\"},{\\\"title\\\":\\\"病情描述\\\",\\\"text\\\":\\\"111\\\",\\\"id\\\":\\\"wx238263450780827648\\\"}]\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:41', '2022-03-28 16:10:41', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1785', 'C2C', '2022032518', '17153', '16796', '2022-03-24 11:36:26', '1136340001', '47452544', 'TIMTextElem', '11', '{\"Text\":\"11\",\"Desc\":\"\",\"Data\":\"\",\"Ext\":\"\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:41', '2022-03-28 16:10:41', null, null, '0');
INSERT INTO `im_tx_msg` VALUES ('1786', 'C2C', '2022032518', '17153', '16796', '2022-03-24 11:45:07', '1145150001', '11188110', 'TIMTextElem', '[敲打]', '{\"Text\":\"[敲打]\",\"Desc\":\"\",\"Data\":\"\",\"Ext\":\"\"}', '', 'null', '0', '113.68.154.51', 'Web', '0', '2022-03-28 16:10:41', '2022-03-28 16:10:41', null, null, '0');

-- ----------------------------
-- Table structure for im_tx_msg_file
-- ----------------------------
DROP TABLE IF EXISTS `im_tx_msg_file`;
CREATE TABLE `im_tx_msg_file` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `chat_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '消息类型',
  `msg_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '消息时间',
  `url` varchar(256) DEFAULT '' COMMENT '下载路径',
  `expire_time` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '过期时间',
  `file_size` int(10) DEFAULT '0' COMMENT '文件大小',
  `file_md5` varchar(100) DEFAULT NULL COMMENT '文件md5',
  `gzip_size` int(10) DEFAULT '0' COMMENT '压缩大小',
  `gzip_md5` varchar(100) DEFAULT NULL COMMENT '压缩md5',
  `error_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '请求code',
  `error_info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '请求信息',
  `action_status` varchar(100) DEFAULT '' COMMENT '请求状态',
  `local_file` varchar(255) DEFAULT '' COMMENT '本地路径',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未下载 1已下载  2已处理',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_im_tx_msg_file_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=396 DEFAULT CHARSET=utf8 COMMENT='消息文件表';

-- ----------------------------
-- Records of im_tx_msg_file
-- ----------------------------
INSERT INTO `im_tx_msg_file` VALUES ('394', 'C2C', '2022032517', 'https://download. /4/cddc88667f8beac1ae6e11ec8724b8cef6d20907.gz', '2022-03-28 16:30:15', '1993', 'c93dac66ab9a1b70e9a71234b656470a', '759', 'f288774d275b2eb9c883f53460cca5c1', '\0', '', 'OK', '', '1', '2022-03-28 16:10:20', '2022-03-28 16:10:20', null);
INSERT INTO `im_tx_msg_file` VALUES ('395', 'C2C', '2022032518', 'https://download. /4/cddc88667f8beac1ae6e11ec8724b8cef6d20907.gz', '2022-03-28 16:30:35', '1993', 'c93dac66ab9a1b70e9a71234b656470a', '759', 'f288774d275b2eb9c883f53460cca5c1', '\0', '', 'OK', '', '1', '2022-03-28 16:10:40', '2022-03-28 16:10:40', null);

-- ----------------------------
-- Table structure for im_txim
-- ----------------------------
DROP TABLE IF EXISTS `im_txim`;
CREATE TABLE `im_txim` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '名称',
  `app_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'appid',
  `identifier` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '管理员帐号',
  `user_sig` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '签名',
  `run_times` int(10) DEFAULT '0' COMMENT '运行次数',
  `begin_time` datetime DEFAULT NULL COMMENT '开始时间',
  `now_time` datetime DEFAULT NULL COMMENT '当前时间',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0未上传,1正常',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `status_run` smallint(6) DEFAULT NULL COMMENT '运行状态',
  `interval` int(10) DEFAULT '0' COMMENT '间隔时间 分钟',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_im_txim_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='腾讯IM帐户信息';

-- ----------------------------
-- Records of im_txim
-- ----------------------------
INSERT INTO `im_txim` VALUES ('1', '测试IM', '1400601443', 'admin', 'eJwtjMEKgkAURf9ltoU97b0hhBaRENgQRbUI2hgz2aMcxEaTon9vUpf3nMv5iIPaB42pRCyiAMS426yNdXzlDme6YDuIp75nZclaxCECSAgRp70xbcmV8ZyIIgDoqePiz6Q-E6GUQ4Vz38Xa3lbN6PQ*T1KdbB*ZoRzbXO3qxeWo1s68ZsuNUzZNCOfi*wNdMjHF', '0', '2022-03-24 16:50:12', '2022-03-24 16:50:12', '1', '2022-03-24 16:50:12', '0000-00-00 00:00:00', null, '0', '0');

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES ('44', '2022-04-21 11:38:05', '2022-04-21 11:38:05', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiODA0ZTM1OGYtYWFhOC00YWU3LTk1NTktZmQxN2IwMmZiOTczIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY1MTExNjQwNSwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjUwNTEwNjA1fQ.bqpOH5wt1SbuMOdDoDUz7uxHmSpddsks7Vto5qthf8Q');
INSERT INTO `jwt_blacklists` VALUES ('45', '2022-04-21 13:30:33', '2022-04-21 13:30:33', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiODA0ZTM1OGYtYWFhOC00YWU3LTk1NTktZmQxN2IwMmZiOTczIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY1MTExNzA4OCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjUwNTExMjg4fQ.4C4bIklsq7mafeFwehHAeVCWlh3DxWLBjdKttB1KFeg');
INSERT INTO `jwt_blacklists` VALUES ('46', '2022-04-21 13:33:57', '2022-04-21 13:33:57', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiODA0ZTM1OGYtYWFhOC00YWU3LTk1NTktZmQxN2IwMmZiOTczIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY1MTEyMzgzOCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjUwNTE4MDM4fQ.7t4ytDvvEQhtLebeCoFgSd8XwZMLqW8qCRJxL1guj5M');

-- ----------------------------
-- Table structure for k8s_clusters
-- ----------------------------
DROP TABLE IF EXISTS `k8s_clusters`;
CREATE TABLE `k8s_clusters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `cluster_name` varchar(191) DEFAULT NULL COMMENT '集群名称',
  `kube_config` varchar(12800) DEFAULT NULL COMMENT 'Kubeconfig内容',
  `cluster_version` float DEFAULT NULL COMMENT '集群版本',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_k8s_clusters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='集群';

-- ----------------------------
-- Records of k8s_clusters
-- ----------------------------
INSERT INTO `k8s_clusters` VALUES ('2', '群集1', 'config', '0', '2021-12-21 14:21:21', '2021-12-21 14:21:21', null);

-- ----------------------------
-- Table structure for k8s_deployments
-- ----------------------------
DROP TABLE IF EXISTS `k8s_deployments`;
CREATE TABLE `k8s_deployments` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `namespace` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '命名空间',
  `deployment` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '应用',
  `replicas` int(11) DEFAULT NULL COMMENT '实例数',
  `create_time` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_k8s_deployments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='应用';

-- ----------------------------
-- Records of k8s_deployments
-- ----------------------------
INSERT INTO `k8s_deployments` VALUES ('2', 'go-cms', 'cms', '434', '34', '2021-12-21 14:21:45', '2021-12-21 14:21:45', null);

-- ----------------------------
-- Table structure for k8s_namespaces
-- ----------------------------
DROP TABLE IF EXISTS `k8s_namespaces`;
CREATE TABLE `k8s_namespaces` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `namespace` varchar(191) DEFAULT NULL COMMENT '命名空间',
  `status` varchar(191) DEFAULT NULL COMMENT '状态',
  `create_time` varchar(191) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_k8s_namespaces_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of k8s_namespaces
-- ----------------------------
INSERT INTO `k8s_namespaces` VALUES ('1', '3434', '343', '434', '2021-12-21 14:21:39', '2021-12-21 14:21:39', null);

-- ----------------------------
-- Table structure for k8s_nodes
-- ----------------------------
DROP TABLE IF EXISTS `k8s_nodes`;
CREATE TABLE `k8s_nodes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `node_name` varchar(191) DEFAULT NULL COMMENT '节点名称',
  `ip` varchar(191) DEFAULT NULL COMMENT 'ip地址',
  `status` varchar(191) DEFAULT NULL COMMENT '节点状态',
  `roles` varchar(191) DEFAULT NULL COMMENT '节点角色',
  `create_time` varchar(191) DEFAULT NULL COMMENT '创建时间',
  `version` varchar(191) DEFAULT NULL COMMENT '节点版本',
  `resource` varchar(191) DEFAULT NULL COMMENT '节点资源',
  `label` varchar(191) DEFAULT NULL COMMENT '节点标签',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_k8s_nodes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of k8s_nodes
-- ----------------------------
INSERT INTO `k8s_nodes` VALUES ('3', 'node1', '333', '33', '333', '33', '33', '33', '333', '2021-12-21 14:21:33', '2021-12-21 14:21:33', null);
INSERT INTO `k8s_nodes` VALUES ('4', 'node2', 'node2', 'node2', '', '', '', '', '', '2021-12-21 14:23:28', '2021-12-21 14:23:28', null);

-- ----------------------------
-- Table structure for k8s_pods
-- ----------------------------
DROP TABLE IF EXISTS `k8s_pods`;
CREATE TABLE `k8s_pods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pod_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '容器名称',
  `pod_ip` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '容器IP',
  `host_ip` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '主机IP',
  `status` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '状态',
  `start_time` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '启动时间',
  `restart_count` int(11) DEFAULT NULL COMMENT '重启次数',
  `name_space` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '命名空间',
  `line` bigint(20) DEFAULT NULL COMMENT 'line',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_k8s_pods_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of k8s_pods
-- ----------------------------

-- ----------------------------
-- Table structure for mem_address
-- ----------------------------
DROP TABLE IF EXISTS `mem_address`;
CREATE TABLE `mem_address` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `realname` varchar(60) DEFAULT NULL COMMENT '收货人',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `email` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '邮箱地址',
  `area_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '地区编码',
  `address` varchar(120) DEFAULT NULL COMMENT '详细地址',
  `zip_code` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '邮政编码',
  `is_default` tinyint(1) DEFAULT '0' COMMENT '默认收货地址',
  `status` smallint(2) DEFAULT '0' COMMENT '状态',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `idx_mem_address_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='会员地址表';

-- ----------------------------
-- Records of mem_address
-- ----------------------------
INSERT INTO `mem_address` VALUES ('1', '0', '小李', '188288388488', '102@qq.com', '120202', '广东省深圳市', 'qweqw', '1', '1', '2022-01-04 16:36:09', null, '0000-00-00 00:00:00');

-- ----------------------------
-- Table structure for mem_user
-- ----------------------------
DROP TABLE IF EXISTS `mem_user`;
CREATE TABLE `mem_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '会员id',
  `username` varchar(40) NOT NULL DEFAULT '' COMMENT '用户名',
  `pws` varchar(32) DEFAULT '' COMMENT '密码',
  `pws_slat` varchar(32) DEFAULT '' COMMENT '密码盐',
  `email` varchar(60) DEFAULT '' COMMENT '邮件',
  `mobile` varchar(20) DEFAULT '' COMMENT '手机',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '昵称',
  `realname` varchar(20) DEFAULT '' COMMENT '真实名',
  `card_id` char(18) DEFAULT '' COMMENT '身份证',
  `sex` smallint(2) unsigned DEFAULT '0' COMMENT '性别:0 保密 1 男 2 女',
  `birthday` datetime DEFAULT NULL COMMENT '生日',
  `avatar` varchar(256) DEFAULT '' COMMENT '头像',
  `mobile_validated` tinyint(1) DEFAULT '0' COMMENT '验证手机',
  `email_validated` tinyint(1) DEFAULT '0' COMMENT '验证邮箱',
  `realname_validated` tinyint(1) DEFAULT '0' COMMENT '验证实名',
  `login_times` int(11) DEFAULT '0' COMMENT '登录次数',
  `recommend_id` bigint(20) DEFAULT '0' COMMENT '推荐人ID',
  `recommend_id2` bigint(20) DEFAULT NULL COMMENT '推荐人ID2',
  `recommend_code` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '推荐码16位（自己的）',
  `status` smallint(2) DEFAULT '1' COMMENT '状态:0=审核中,1=审核通过 2=审核未通过3=禁止登录4=发色情非法信息5:已注销6:非法攻击',
  `status_safe` smallint(2) DEFAULT '0' COMMENT '安全状态:0=正常，1=修改了密码 2=修改了手机号',
  `reg_ip` bigint(20) unsigned DEFAULT '0' COMMENT '注册ip',
  `last_login_ip` bigint(20) unsigned DEFAULT '0' COMMENT '登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_mobile` (`mobile`) USING BTREE,
  KEY `idx_mem_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of mem_user
-- ----------------------------
INSERT INTO `mem_user` VALUES ('1', 'test1', 'test1', 'test1', 'test1', 'test1', '', '', '', '2', '2021-12-22 14:25:10', '20a6e4a766d14c30ad79933f0d940600', '0', '1', '0', '0', '0', '0', '', '0', '1', '0', '0', '2021-12-22 14:25:10', '2021-12-31 10:55:41', null, '0000-00-00 00:00:00');
INSERT INTO `mem_user` VALUES ('2', '22', '22', '22', '222', '2', '2', '22', '2', '1', '2021-12-31 10:54:48', '1edff067da1e498a8ccd15d6226198ee', '0', '0', '0', '0', '0', '0', '222222', '0', '0', '0', '0', '2021-12-31 10:54:48', '2021-12-31 10:55:14', null, '2021-12-31 10:55:14');

-- ----------------------------
-- Table structure for mem_user_log
-- ----------------------------
DROP TABLE IF EXISTS `mem_user_log`;
CREATE TABLE `mem_user_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `type` smallint(2) DEFAULT '0' COMMENT '类型:1登录 2退出',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0=不正常,1=正常',
  `ip` varchar(50) DEFAULT NULL COMMENT 'ip',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `idx_mem_user_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户日志';

-- ----------------------------
-- Records of mem_user_log
-- ----------------------------
INSERT INTO `mem_user_log` VALUES ('1', '1', '1', '0', '222', '2022-01-04 16:36:21', null, '0000-00-00 00:00:00');

-- ----------------------------
-- Table structure for mem_user_safe
-- ----------------------------
DROP TABLE IF EXISTS `mem_user_safe`;
CREATE TABLE `mem_user_safe` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `type` smallint(2) DEFAULT '0' COMMENT '类型:1改密码 2改手机号 3改用户名 4实名认证',
  `val_old` smallint(2) DEFAULT '0' COMMENT '旧值',
  `val_new` smallint(2) DEFAULT '0' COMMENT '新值',
  `status` smallint(2) DEFAULT '0' COMMENT '状态:0=审核中,1=审核通过,2=审核未通过',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_mem_user_safe_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户安全操作表';

-- ----------------------------
-- Records of mem_user_safe
-- ----------------------------
INSERT INTO `mem_user_safe` VALUES ('1', '1', '1', '1', '1', '0', '2021-12-22 14:26:02', null, '2021-12-22 14:26:02');

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=387 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES ('1', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/base/login', '用户登录（必选）', 'base', 'POST');
INSERT INTO `sys_apis` VALUES ('2', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/register', '用户注册（必选）', 'user', 'POST');
INSERT INTO `sys_apis` VALUES ('3', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('4', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('5', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('6', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('7', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('8', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('9', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/authority/createAuthority', '创建角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES ('10', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/authority/deleteAuthority', '删除角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES ('11', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/authority/getAuthorityList', '获取角色列表', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES ('12', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/getMenu', '获取菜单树（必选）', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('13', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('14', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/addBaseMenu', '新增菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('15', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('16', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('17', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/getMenuAuthority', '获取指定角色menu', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('18', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/deleteBaseMenu', '删除菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('19', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/updateBaseMenu', '更新菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('20', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/menu/getBaseMenuById', '根据id获取菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES ('21', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/changePassword', '修改密码（建议选择）', 'user', 'POST');
INSERT INTO `sys_apis` VALUES ('23', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/getUserList', '获取用户列表', 'user', 'POST');
INSERT INTO `sys_apis` VALUES ('24', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/setUserAuthority', '修改用户角色（必选）', 'user', 'POST');
INSERT INTO `sys_apis` VALUES ('25', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/fileUploadAndDownload/upload', '文件上传示例', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES ('26', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/fileUploadAndDownload/getFileList', '获取上传文件列表', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES ('27', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES ('28', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES ('29', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/fileUploadAndDownload/deleteFile', '删除文件', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES ('30', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES ('31', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/authority/setDataAuthority', '设置角色资源权限', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES ('32', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/system/getSystemConfig', '获取配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES ('33', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/system/setSystemConfig', '设置配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES ('34', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/customer/customer', '创建客户', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES ('35', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/customer/customer', '更新客户', 'customer', 'PUT');
INSERT INTO `sys_apis` VALUES ('36', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/customer/customer', '删除客户', 'customer', 'DELETE');
INSERT INTO `sys_apis` VALUES ('37', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/customer/customer', '获取单一客户', 'customer', 'GET');
INSERT INTO `sys_apis` VALUES ('38', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/customer/customerList', '获取客户列表', 'customer', 'GET');
INSERT INTO `sys_apis` VALUES ('39', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/casbin/casbinTest/:pathParam', 'RESTFUL模式测试', 'casbin', 'GET');
INSERT INTO `sys_apis` VALUES ('40', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/createTemp', '自动化代码', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES ('41', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/authority/updateAuthority', '更新角色信息', 'authority', 'PUT');
INSERT INTO `sys_apis` VALUES ('42', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/authority/copyAuthority', '拷贝角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES ('43', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/deleteUser', '删除用户', 'user', 'DELETE');
INSERT INTO `sys_apis` VALUES ('44', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', 'sysDictionaryDetail', 'POST');
INSERT INTO `sys_apis` VALUES ('45', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', 'sysDictionaryDetail', 'DELETE');
INSERT INTO `sys_apis` VALUES ('46', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', 'sysDictionaryDetail', 'PUT');
INSERT INTO `sys_apis` VALUES ('47', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', 'sysDictionaryDetail', 'GET');
INSERT INTO `sys_apis` VALUES ('48', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', 'sysDictionaryDetail', 'GET');
INSERT INTO `sys_apis` VALUES ('49', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionary/createSysDictionary', '新增字典', 'sysDictionary', 'POST');
INSERT INTO `sys_apis` VALUES ('50', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionary/deleteSysDictionary', '删除字典', 'sysDictionary', 'DELETE');
INSERT INTO `sys_apis` VALUES ('51', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionary/updateSysDictionary', '更新字典', 'sysDictionary', 'PUT');
INSERT INTO `sys_apis` VALUES ('52', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionary/findSysDictionary', '根据ID获取字典', 'sysDictionary', 'GET');
INSERT INTO `sys_apis` VALUES ('53', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysDictionary/getSysDictionaryList', '获取字典列表', 'sysDictionary', 'GET');
INSERT INTO `sys_apis` VALUES ('54', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', 'sysOperationRecord', 'POST');
INSERT INTO `sys_apis` VALUES ('55', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', 'sysOperationRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES ('56', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES ('57', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES ('58', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/getTables', '获取数据库表', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES ('59', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/getDB', '获取所有数据库', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES ('60', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/getColumn', '获取所选table的所有字段', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES ('61', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', 'sysOperationRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES ('62', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/simpleUploader/upload', '插件版分片上传', 'simpleUploader', 'POST');
INSERT INTO `sys_apis` VALUES ('63', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/simpleUploader/checkFileMd5', '文件完整度验证', 'simpleUploader', 'GET');
INSERT INTO `sys_apis` VALUES ('64', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/simpleUploader/mergeFileMd5', '上传完成合并文件', 'simpleUploader', 'GET');
INSERT INTO `sys_apis` VALUES ('65', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/setUserInfo', '设置用户信息（必选）', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES ('66', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/system/getServerInfo', '获取服务器信息', 'system', 'POST');
INSERT INTO `sys_apis` VALUES ('67', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES ('80', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/preview', '预览自动化代码', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES ('81', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/excel/importExcel', '导入excel', 'excel', 'POST');
INSERT INTO `sys_apis` VALUES ('82', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/excel/loadExcel', '下载excel', 'excel', 'GET');
INSERT INTO `sys_apis` VALUES ('83', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/excel/exportExcel', '导出excel', 'excel', 'POST');
INSERT INTO `sys_apis` VALUES ('84', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/excel/downloadTemplate', '下载excel模板', 'excel', 'GET');
INSERT INTO `sys_apis` VALUES ('85', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES ('86', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/getSysHistory', '查询回滚记录', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES ('87', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/rollback', '回滚自动生成代码', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES ('88', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/getMeta', '获取meta信息', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES ('89', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/superBuilder/delSysHistory', '删除回滚记录', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES ('90', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/setUserAuthorities', '设置权限组', 'user', 'POST');
INSERT INTO `sys_apis` VALUES ('91', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/user/getUserInfo', '获取自身信息（必选）', 'user', 'GET');
INSERT INTO `sys_apis` VALUES ('92', '2021-09-11 13:58:51', '2021-09-11 13:58:51', null, '/cmsCat/createCmsCat', '新增cmsCat表', 'cmsCat', 'POST');
INSERT INTO `sys_apis` VALUES ('93', '2021-09-11 13:58:51', '2021-09-11 13:58:51', null, '/cmsCat/deleteCmsCat', '删除cmsCat表', 'cmsCat', 'DELETE');
INSERT INTO `sys_apis` VALUES ('94', '2021-09-11 13:58:51', '2021-09-11 13:58:51', null, '/cmsCat/deleteCmsCatByIds', '批量删除cmsCat表', 'cmsCat', 'DELETE');
INSERT INTO `sys_apis` VALUES ('95', '2021-09-11 13:58:51', '2021-09-11 13:58:51', null, '/cmsCat/updateCmsCat', '更新cmsCat表', 'cmsCat', 'PUT');
INSERT INTO `sys_apis` VALUES ('96', '2021-09-11 13:58:51', '2021-09-11 13:58:51', null, '/cmsCat/findCmsCat', '根据ID获取cmsCat表', 'cmsCat', 'GET');
INSERT INTO `sys_apis` VALUES ('97', '2021-09-11 13:58:51', '2021-09-11 13:58:51', null, '/cmsCat/getCmsCatList', '获取cmsCat表列表', 'cmsCat', 'GET');
INSERT INTO `sys_apis` VALUES ('98', '2021-09-13 16:47:57', '2021-09-13 16:47:57', null, '/cmsCat/quickEdit', '快速编辑 ', 'cmsCat', 'POST');
INSERT INTO `sys_apis` VALUES ('99', '2021-09-24 11:18:52', '2021-09-24 11:18:52', null, '/basicFile/createBasicFile', '新增basicFile表', 'basicFile', 'POST');
INSERT INTO `sys_apis` VALUES ('100', '2021-09-24 11:18:52', '2021-09-24 11:18:52', null, '/basicFile/deleteBasicFile', '删除basicFile表', 'basicFile', 'DELETE');
INSERT INTO `sys_apis` VALUES ('101', '2021-09-24 11:18:52', '2021-09-24 11:18:52', null, '/basicFile/deleteBasicFileByIds', '批量删除basicFile表', 'basicFile', 'DELETE');
INSERT INTO `sys_apis` VALUES ('102', '2021-09-24 11:18:52', '2021-09-24 11:18:52', null, '/basicFile/updateBasicFile', '更新basicFile表', 'basicFile', 'PUT');
INSERT INTO `sys_apis` VALUES ('103', '2021-09-24 11:18:52', '2021-09-24 11:18:52', null, '/basicFile/findBasicFile', '根据ID获取basicFile表', 'basicFile', 'GET');
INSERT INTO `sys_apis` VALUES ('104', '2021-09-24 11:18:53', '2021-09-24 11:18:53', null, '/basicFile/getBasicFileList', '获取basicFile表列表', 'basicFile', 'GET');
INSERT INTO `sys_apis` VALUES ('105', '2021-09-24 11:18:53', '2021-09-24 11:18:53', null, '/basicFile/quickEdit', '快速编辑basicFile表', 'basicFile', 'POST');
INSERT INTO `sys_apis` VALUES ('106', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/commFile/upload', 'commFile', 'common', 'POST');
INSERT INTO `sys_apis` VALUES ('107', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '/commDb/quickEdit', 'commDb', 'common', 'POST');
INSERT INTO `sys_apis` VALUES ('108', '2021-10-13 14:24:48', '2021-10-13 14:24:48', null, '/cmsAdSeat/createCmsAdSeat', '新增cmsAdSeat表', 'cmsAdSeat', 'POST');
INSERT INTO `sys_apis` VALUES ('109', '2021-10-13 14:24:48', '2021-10-13 14:24:48', null, '/cmsAdSeat/deleteCmsAdSeat', '删除cmsAdSeat表', 'cmsAdSeat', 'DELETE');
INSERT INTO `sys_apis` VALUES ('110', '2021-10-13 14:24:48', '2021-10-13 14:24:48', null, '/cmsAdSeat/deleteCmsAdSeatByIds', '批量删除cmsAdSeat表', 'cmsAdSeat', 'DELETE');
INSERT INTO `sys_apis` VALUES ('111', '2021-10-13 14:24:48', '2021-10-13 14:24:48', null, '/cmsAdSeat/updateCmsAdSeat', '更新cmsAdSeat表', 'cmsAdSeat', 'PUT');
INSERT INTO `sys_apis` VALUES ('112', '2021-10-13 14:24:48', '2021-10-13 14:24:48', null, '/cmsAdSeat/findCmsAdSeat', '根据ID获取cmsAdSeat表', 'cmsAdSeat', 'GET');
INSERT INTO `sys_apis` VALUES ('113', '2021-10-13 14:24:48', '2021-10-13 14:24:48', null, '/cmsAdSeat/getCmsAdSeatList', '获取cmsAdSeat表列表', 'cmsAdSeat', 'GET');
INSERT INTO `sys_apis` VALUES ('114', '2021-10-13 14:24:48', '2021-10-13 14:24:48', null, '/cmsAdSeat/quickEdit', '快速编辑cmsAdSeat表', 'cmsAdSeat', 'POST');
INSERT INTO `sys_apis` VALUES ('115', '2021-10-14 23:11:23', '2021-10-14 23:11:23', null, '/cmsAd/createCmsAd', '新增cmsAd表', 'cmsAd', 'POST');
INSERT INTO `sys_apis` VALUES ('116', '2021-10-14 23:11:23', '2021-10-14 23:11:23', null, '/cmsAd/deleteCmsAd', '删除cmsAd表', 'cmsAd', 'DELETE');
INSERT INTO `sys_apis` VALUES ('117', '2021-10-14 23:11:23', '2021-10-14 23:11:23', null, '/cmsAd/deleteCmsAdByIds', '批量删除cmsAd表', 'cmsAd', 'DELETE');
INSERT INTO `sys_apis` VALUES ('118', '2021-10-14 23:11:23', '2021-10-14 23:11:23', null, '/cmsAd/updateCmsAd', '更新cmsAd表', 'cmsAd', 'PUT');
INSERT INTO `sys_apis` VALUES ('119', '2021-10-14 23:11:23', '2021-10-14 23:11:23', null, '/cmsAd/findCmsAd', '根据ID获取cmsAd表', 'cmsAd', 'GET');
INSERT INTO `sys_apis` VALUES ('120', '2021-10-14 23:11:23', '2021-10-14 23:11:23', null, '/cmsAd/getCmsAdList', '获取cmsAd表列表', 'cmsAd', 'GET');
INSERT INTO `sys_apis` VALUES ('121', '2021-10-14 23:11:23', '2021-10-14 23:11:23', null, '/cmsAd/quickEdit', '快速编辑cmsAd表', 'cmsAd', 'POST');
INSERT INTO `sys_apis` VALUES ('122', '2021-10-14 23:35:39', '2021-10-14 23:35:39', null, '/cmsArticle/createCmsArticle', '新增cmsArticle表', 'cmsArticle', 'POST');
INSERT INTO `sys_apis` VALUES ('123', '2021-10-14 23:35:39', '2021-10-14 23:35:39', null, '/cmsArticle/deleteCmsArticle', '删除cmsArticle表', 'cmsArticle', 'DELETE');
INSERT INTO `sys_apis` VALUES ('124', '2021-10-14 23:35:39', '2021-10-14 23:35:39', null, '/cmsArticle/deleteCmsArticleByIds', '批量删除cmsArticle表', 'cmsArticle', 'DELETE');
INSERT INTO `sys_apis` VALUES ('125', '2021-10-14 23:35:39', '2021-10-14 23:35:39', null, '/cmsArticle/updateCmsArticle', '更新cmsArticle表', 'cmsArticle', 'PUT');
INSERT INTO `sys_apis` VALUES ('126', '2021-10-14 23:35:39', '2021-10-14 23:35:39', null, '/cmsArticle/findCmsArticle', '根据ID获取cmsArticle表', 'cmsArticle', 'GET');
INSERT INTO `sys_apis` VALUES ('127', '2021-10-14 23:35:39', '2021-10-14 23:35:39', null, '/cmsArticle/getCmsArticleList', '获取cmsArticle表列表', 'cmsArticle', 'GET');
INSERT INTO `sys_apis` VALUES ('128', '2021-10-14 23:35:39', '2021-10-14 23:35:39', null, '/cmsArticle/quickEdit', '快速编辑cmsArticle表', 'cmsArticle', 'POST');
INSERT INTO `sys_apis` VALUES ('129', '2021-11-09 14:58:21', '2021-11-09 14:58:21', '2021-11-09 14:58:32', '11', '111', '1', 'POST');
INSERT INTO `sys_apis` VALUES ('130', '2021-11-09 14:58:26', '2021-11-09 14:58:26', '2021-11-09 14:58:32', '22', '22', '22', 'GET');
INSERT INTO `sys_apis` VALUES ('131', '2021-11-15 15:32:42', '2021-11-15 15:32:42', null, '/colKeyField/createColKeyField', '新增colKeyField表', 'colKeyField', 'POST');
INSERT INTO `sys_apis` VALUES ('132', '2021-11-15 15:32:42', '2021-11-15 15:32:42', null, '/colKeyField/deleteColKeyField', '删除colKeyField表', 'colKeyField', 'DELETE');
INSERT INTO `sys_apis` VALUES ('133', '2021-11-15 15:32:42', '2021-11-15 15:32:42', null, '/colKeyField/deleteColKeyFieldByIds', '批量删除colKeyField表', 'colKeyField', 'DELETE');
INSERT INTO `sys_apis` VALUES ('134', '2021-11-15 15:32:42', '2021-11-15 15:32:42', null, '/colKeyField/updateColKeyField', '更新colKeyField表', 'colKeyField', 'PUT');
INSERT INTO `sys_apis` VALUES ('135', '2021-11-15 15:32:42', '2021-11-15 15:32:42', null, '/colKeyField/findColKeyField', '根据ID获取colKeyField表', 'colKeyField', 'GET');
INSERT INTO `sys_apis` VALUES ('136', '2021-11-15 15:32:42', '2021-11-15 15:32:42', null, '/colKeyField/getColKeyFieldList', '获取colKeyField表列表', 'colKeyField', 'GET');
INSERT INTO `sys_apis` VALUES ('137', '2021-11-15 15:32:42', '2021-11-15 15:32:42', null, '/colKeyField/quickEdit', '快速编辑colKeyField表', 'colKeyField', 'POST');
INSERT INTO `sys_apis` VALUES ('138', '2021-11-15 15:37:22', '2021-11-15 15:37:22', null, '/colCollect/createColCollect', '新增colCollect表', 'colCollect', 'POST');
INSERT INTO `sys_apis` VALUES ('139', '2021-11-15 15:37:22', '2021-11-15 15:37:22', null, '/colCollect/deleteColCollect', '删除colCollect表', 'colCollect', 'DELETE');
INSERT INTO `sys_apis` VALUES ('140', '2021-11-15 15:37:23', '2021-11-15 15:37:23', null, '/colCollect/deleteColCollectByIds', '批量删除colCollect表', 'colCollect', 'DELETE');
INSERT INTO `sys_apis` VALUES ('141', '2021-11-15 15:37:23', '2021-11-15 15:37:23', null, '/colCollect/updateColCollect', '更新colCollect表', 'colCollect', 'PUT');
INSERT INTO `sys_apis` VALUES ('142', '2021-11-15 15:37:23', '2021-11-15 15:37:23', null, '/colCollect/findColCollect', '根据ID获取colCollect表', 'colCollect', 'GET');
INSERT INTO `sys_apis` VALUES ('143', '2021-11-15 15:37:23', '2021-11-15 15:37:23', null, '/colCollect/getColCollectList', '获取colCollect表列表', 'colCollect', 'GET');
INSERT INTO `sys_apis` VALUES ('144', '2021-11-15 15:37:23', '2021-11-15 15:37:23', null, '/colCollect/quickEdit', '快速编辑colCollect表', 'colCollect', 'POST');
INSERT INTO `sys_apis` VALUES ('145', '2021-11-15 17:11:07', '2021-11-15 17:11:07', null, '/colHsj/createColHsj', '新增colHsj表', 'colHsj', 'POST');
INSERT INTO `sys_apis` VALUES ('146', '2021-11-15 17:11:07', '2021-11-15 17:11:07', null, '/colHsj/deleteColHsj', '删除colHsj表', 'colHsj', 'DELETE');
INSERT INTO `sys_apis` VALUES ('147', '2021-11-15 17:11:07', '2021-11-15 17:11:07', null, '/colHsj/deleteColHsjByIds', '批量删除colHsj表', 'colHsj', 'DELETE');
INSERT INTO `sys_apis` VALUES ('148', '2021-11-15 17:11:07', '2021-11-15 17:11:07', null, '/colHsj/updateColHsj', '更新colHsj表', 'colHsj', 'PUT');
INSERT INTO `sys_apis` VALUES ('149', '2021-11-15 17:11:07', '2021-11-15 17:11:07', null, '/colHsj/findColHsj', '根据ID获取colHsj表', 'colHsj', 'GET');
INSERT INTO `sys_apis` VALUES ('150', '2021-11-15 17:11:07', '2021-11-15 17:11:07', null, '/colHsj/getColHsjList', '获取colHsj表列表', 'colHsj', 'GET');
INSERT INTO `sys_apis` VALUES ('151', '2021-11-15 17:11:08', '2021-11-15 17:11:08', null, '/colHsj/quickEdit', '快速编辑colHsj表', 'colHsj', 'POST');
INSERT INTO `sys_apis` VALUES ('152', '2021-11-16 11:58:36', '2021-11-16 11:59:00', null, '/colCollect/startOrStopCollect', 'startOrStopCollect', 'colCollect', 'GET');
INSERT INTO `sys_apis` VALUES ('153', '2021-11-21 11:50:21', '2021-11-21 11:50:36', null, '/colCollect/excelList', '分页excel导出ColCollect列表', 'colCollect', 'GET');
INSERT INTO `sys_apis` VALUES ('154', '2021-11-21 12:45:49', '2021-11-21 17:03:35', null, '/colHsj/excelList', 'colHsj excelList', 'colHsj', 'POST');
INSERT INTO `sys_apis` VALUES ('155', '2021-11-22 14:40:56', '2021-11-22 14:40:56', null, '/colHsj/excelList', '分页excel导出colHsj表列表', 'colHsj', 'GET');
INSERT INTO `sys_apis` VALUES ('156', '2021-12-03 00:38:53', '2021-12-03 00:38:53', null, '/colKeyField/excelList', '分页excel导出colKeyField表列表', 'colKeyField', 'GET');
INSERT INTO `sys_apis` VALUES ('157', '2021-12-20 13:15:59', '2021-12-20 13:15:59', null, '/cmsAd/excelList', '分页excel导出cmsAd表列表', 'cmsAd', 'GET');
INSERT INTO `sys_apis` VALUES ('158', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/createImTxMsgFile', '新增imTxMsgFile表', 'imTxMsgFile', 'POST');
INSERT INTO `sys_apis` VALUES ('159', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/deleteImTxMsgFile', '删除imTxMsgFile表', 'imTxMsgFile', 'DELETE');
INSERT INTO `sys_apis` VALUES ('160', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/deleteImTxMsgFileByIds', '批量删除imTxMsgFile表', 'imTxMsgFile', 'DELETE');
INSERT INTO `sys_apis` VALUES ('161', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/updateImTxMsgFile', '更新imTxMsgFile表', 'imTxMsgFile', 'PUT');
INSERT INTO `sys_apis` VALUES ('162', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/findImTxMsgFile', '根据ID获取imTxMsgFile表', 'imTxMsgFile', 'GET');
INSERT INTO `sys_apis` VALUES ('163', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/getImTxMsgFileList', '获取imTxMsgFile表列表', 'imTxMsgFile', 'GET');
INSERT INTO `sys_apis` VALUES ('164', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/quickEdit', '快速编辑imTxMsgFile表', 'imTxMsgFile', 'POST');
INSERT INTO `sys_apis` VALUES ('165', '2021-12-20 14:35:41', '2021-12-20 14:35:41', null, '/imTxMsgFile/excelList', '分页excel导出imTxMsgFile表列表', 'imTxMsgFile', 'GET');
INSERT INTO `sys_apis` VALUES ('166', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/createImTxMsg', '新增imTxMsg表', 'imTxMsg', 'POST');
INSERT INTO `sys_apis` VALUES ('167', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/deleteImTxMsg', '删除imTxMsg表', 'imTxMsg', 'DELETE');
INSERT INTO `sys_apis` VALUES ('168', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/deleteImTxMsgByIds', '批量删除imTxMsg表', 'imTxMsg', 'DELETE');
INSERT INTO `sys_apis` VALUES ('169', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/updateImTxMsg', '更新imTxMsg表', 'imTxMsg', 'PUT');
INSERT INTO `sys_apis` VALUES ('170', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/findImTxMsg', '根据ID获取imTxMsg表', 'imTxMsg', 'GET');
INSERT INTO `sys_apis` VALUES ('171', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/getImTxMsgList', '获取imTxMsg表列表', 'imTxMsg', 'GET');
INSERT INTO `sys_apis` VALUES ('172', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/quickEdit', '快速编辑imTxMsg表', 'imTxMsg', 'POST');
INSERT INTO `sys_apis` VALUES ('173', '2021-12-20 14:39:11', '2021-12-20 14:39:11', null, '/imTxMsg/excelList', '分页excel导出imTxMsg表列表', 'imTxMsg', 'GET');
INSERT INTO `sys_apis` VALUES ('174', '2021-12-20 16:39:22', '2021-12-20 16:39:22', null, '/cmsArticle/excelList', '分页excel导出cmsArticle表列表', 'cmsArticle', 'GET');
INSERT INTO `sys_apis` VALUES ('175', '2021-12-20 16:40:12', '2021-12-20 16:40:12', null, '/cmsAdSeat/excelList', '分页excel导出cmsAdSeat表列表', 'cmsAdSeat', 'GET');
INSERT INTO `sys_apis` VALUES ('176', '2021-12-20 16:40:39', '2021-12-20 16:40:39', null, '/cmsCat/excelList', '分页excel导出cmsCat表列表', 'cmsCat', 'GET');
INSERT INTO `sys_apis` VALUES ('177', '2021-12-20 16:45:51', '2021-12-20 16:45:51', null, '/basicFile/excelList', '分页excel导出basicFile表列表', 'basicFile', 'GET');
INSERT INTO `sys_apis` VALUES ('178', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/createImTxim', '新增imTxim表', 'imTxim', 'POST');
INSERT INTO `sys_apis` VALUES ('179', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/deleteImTxim', '删除imTxim表', 'imTxim', 'DELETE');
INSERT INTO `sys_apis` VALUES ('180', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/deleteImTximByIds', '批量删除imTxim表', 'imTxim', 'DELETE');
INSERT INTO `sys_apis` VALUES ('181', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/updateImTxim', '更新imTxim表', 'imTxim', 'PUT');
INSERT INTO `sys_apis` VALUES ('182', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/findImTxim', '根据ID获取imTxim表', 'imTxim', 'GET');
INSERT INTO `sys_apis` VALUES ('183', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/getImTximList', '获取imTxim表列表', 'imTxim', 'GET');
INSERT INTO `sys_apis` VALUES ('184', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/quickEdit', '快速编辑imTxim表', 'imTxim', 'POST');
INSERT INTO `sys_apis` VALUES ('185', '2021-12-21 10:29:26', '2021-12-21 10:29:26', null, '/imTxim/excelList', '分页excel导出imTxim表列表', 'imTxim', 'GET');
INSERT INTO `sys_apis` VALUES ('186', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/createK8sClusters', '新增k8sClusters表', 'k8sClusters', 'POST');
INSERT INTO `sys_apis` VALUES ('187', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/deleteK8sClusters', '删除k8sClusters表', 'k8sClusters', 'DELETE');
INSERT INTO `sys_apis` VALUES ('188', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/deleteK8sClustersByIds', '批量删除k8sClusters表', 'k8sClusters', 'DELETE');
INSERT INTO `sys_apis` VALUES ('189', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/updateK8sClusters', '更新k8sClusters表', 'k8sClusters', 'PUT');
INSERT INTO `sys_apis` VALUES ('190', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/findK8sClusters', '根据ID获取k8sClusters表', 'k8sClusters', 'GET');
INSERT INTO `sys_apis` VALUES ('191', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/getK8sClustersList', '获取k8sClusters表列表', 'k8sClusters', 'GET');
INSERT INTO `sys_apis` VALUES ('192', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/quickEdit', '快速编辑k8sClusters表', 'k8sClusters', 'POST');
INSERT INTO `sys_apis` VALUES ('193', '2021-12-21 11:15:40', '2021-12-21 11:15:40', null, '/k8sClusters/excelList', '分页excel导出k8sClusters表列表', 'k8sClusters', 'GET');
INSERT INTO `sys_apis` VALUES ('194', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/createK8sDeployments', '新增k8sDeployments表', 'k8sDeployments', 'POST');
INSERT INTO `sys_apis` VALUES ('195', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/deleteK8sDeployments', '删除k8sDeployments表', 'k8sDeployments', 'DELETE');
INSERT INTO `sys_apis` VALUES ('196', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/deleteK8sDeploymentsByIds', '批量删除k8sDeployments表', 'k8sDeployments', 'DELETE');
INSERT INTO `sys_apis` VALUES ('197', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/updateK8sDeployments', '更新k8sDeployments表', 'k8sDeployments', 'PUT');
INSERT INTO `sys_apis` VALUES ('198', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/findK8sDeployments', '根据ID获取k8sDeployments表', 'k8sDeployments', 'GET');
INSERT INTO `sys_apis` VALUES ('199', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/getK8sDeploymentsList', '获取k8sDeployments表列表', 'k8sDeployments', 'GET');
INSERT INTO `sys_apis` VALUES ('200', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/quickEdit', '快速编辑k8sDeployments表', 'k8sDeployments', 'POST');
INSERT INTO `sys_apis` VALUES ('201', '2021-12-21 11:17:03', '2021-12-21 11:17:03', null, '/k8sDeployments/excelList', '分页excel导出k8sDeployments表列表', 'k8sDeployments', 'GET');
INSERT INTO `sys_apis` VALUES ('202', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/createK8sNamespaces', '新增k8sNamespaces表', 'k8sNamespaces', 'POST');
INSERT INTO `sys_apis` VALUES ('203', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/deleteK8sNamespaces', '删除k8sNamespaces表', 'k8sNamespaces', 'DELETE');
INSERT INTO `sys_apis` VALUES ('204', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/deleteK8sNamespacesByIds', '批量删除k8sNamespaces表', 'k8sNamespaces', 'DELETE');
INSERT INTO `sys_apis` VALUES ('205', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/updateK8sNamespaces', '更新k8sNamespaces表', 'k8sNamespaces', 'PUT');
INSERT INTO `sys_apis` VALUES ('206', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/findK8sNamespaces', '根据ID获取k8sNamespaces表', 'k8sNamespaces', 'GET');
INSERT INTO `sys_apis` VALUES ('207', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/getK8sNamespacesList', '获取k8sNamespaces表列表', 'k8sNamespaces', 'GET');
INSERT INTO `sys_apis` VALUES ('208', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/quickEdit', '快速编辑k8sNamespaces表', 'k8sNamespaces', 'POST');
INSERT INTO `sys_apis` VALUES ('209', '2021-12-21 11:18:17', '2021-12-21 11:18:17', null, '/k8sNamespaces/excelList', '分页excel导出k8sNamespaces表列表', 'k8sNamespaces', 'GET');
INSERT INTO `sys_apis` VALUES ('210', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/createK8sNodes', '新增k8sNodes表', 'k8sNodes', 'POST');
INSERT INTO `sys_apis` VALUES ('211', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/deleteK8sNodes', '删除k8sNodes表', 'k8sNodes', 'DELETE');
INSERT INTO `sys_apis` VALUES ('212', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/deleteK8sNodesByIds', '批量删除k8sNodes表', 'k8sNodes', 'DELETE');
INSERT INTO `sys_apis` VALUES ('213', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/updateK8sNodes', '更新k8sNodes表', 'k8sNodes', 'PUT');
INSERT INTO `sys_apis` VALUES ('214', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/findK8sNodes', '根据ID获取k8sNodes表', 'k8sNodes', 'GET');
INSERT INTO `sys_apis` VALUES ('215', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/getK8sNodesList', '获取k8sNodes表列表', 'k8sNodes', 'GET');
INSERT INTO `sys_apis` VALUES ('216', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/quickEdit', '快速编辑k8sNodes表', 'k8sNodes', 'POST');
INSERT INTO `sys_apis` VALUES ('217', '2021-12-21 11:19:37', '2021-12-21 11:19:37', null, '/k8sNodes/excelList', '分页excel导出k8sNodes表列表', 'k8sNodes', 'GET');
INSERT INTO `sys_apis` VALUES ('218', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/createK8sPods', '新增k8sPods表', 'k8sPods', 'POST');
INSERT INTO `sys_apis` VALUES ('219', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/deleteK8sPods', '删除k8sPods表', 'k8sPods', 'DELETE');
INSERT INTO `sys_apis` VALUES ('220', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/deleteK8sPodsByIds', '批量删除k8sPods表', 'k8sPods', 'DELETE');
INSERT INTO `sys_apis` VALUES ('221', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/updateK8sPods', '更新k8sPods表', 'k8sPods', 'PUT');
INSERT INTO `sys_apis` VALUES ('222', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/findK8sPods', '根据ID获取k8sPods表', 'k8sPods', 'GET');
INSERT INTO `sys_apis` VALUES ('223', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/getK8sPodsList', '获取k8sPods表列表', 'k8sPods', 'GET');
INSERT INTO `sys_apis` VALUES ('224', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/quickEdit', '快速编辑k8sPods表', 'k8sPods', 'POST');
INSERT INTO `sys_apis` VALUES ('225', '2021-12-21 11:21:10', '2021-12-21 11:21:10', null, '/k8sPods/excelList', '分页excel导出k8sPods表列表', 'k8sPods', 'GET');
INSERT INTO `sys_apis` VALUES ('226', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/createMemUser', '新增memUser表', 'memUser', 'POST');
INSERT INTO `sys_apis` VALUES ('227', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/deleteMemUser', '删除memUser表', 'memUser', 'DELETE');
INSERT INTO `sys_apis` VALUES ('228', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/deleteMemUserByIds', '批量删除memUser表', 'memUser', 'DELETE');
INSERT INTO `sys_apis` VALUES ('229', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/updateMemUser', '更新memUser表', 'memUser', 'PUT');
INSERT INTO `sys_apis` VALUES ('230', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/findMemUser', '根据ID获取memUser表', 'memUser', 'GET');
INSERT INTO `sys_apis` VALUES ('231', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/getMemUserList', '获取memUser表列表', 'memUser', 'GET');
INSERT INTO `sys_apis` VALUES ('232', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/quickEdit', '快速编辑memUser表', 'memUser', 'POST');
INSERT INTO `sys_apis` VALUES ('233', '2021-12-22 13:57:11', '2021-12-22 13:57:11', null, '/memUser/excelList', '分页excel导出memUser表列表', 'memUser', 'GET');
INSERT INTO `sys_apis` VALUES ('234', '2021-12-22 13:58:39', '2021-12-22 13:58:39', null, '/memUserSafe/createMemUserSafe', '新增memUserSafe表', 'memUserSafe', 'POST');
INSERT INTO `sys_apis` VALUES ('235', '2021-12-22 13:58:40', '2021-12-22 13:58:40', null, '/memUserSafe/deleteMemUserSafe', '删除memUserSafe表', 'memUserSafe', 'DELETE');
INSERT INTO `sys_apis` VALUES ('236', '2021-12-22 13:58:40', '2021-12-22 13:58:40', null, '/memUserSafe/deleteMemUserSafeByIds', '批量删除memUserSafe表', 'memUserSafe', 'DELETE');
INSERT INTO `sys_apis` VALUES ('237', '2021-12-22 13:58:40', '2021-12-22 13:58:40', null, '/memUserSafe/updateMemUserSafe', '更新memUserSafe表', 'memUserSafe', 'PUT');
INSERT INTO `sys_apis` VALUES ('238', '2021-12-22 13:58:40', '2021-12-22 13:58:40', null, '/memUserSafe/findMemUserSafe', '根据ID获取memUserSafe表', 'memUserSafe', 'GET');
INSERT INTO `sys_apis` VALUES ('239', '2021-12-22 13:58:40', '2021-12-22 13:58:40', null, '/memUserSafe/getMemUserSafeList', '获取memUserSafe表列表', 'memUserSafe', 'GET');
INSERT INTO `sys_apis` VALUES ('240', '2021-12-22 13:58:40', '2021-12-22 13:58:40', null, '/memUserSafe/quickEdit', '快速编辑memUserSafe表', 'memUserSafe', 'POST');
INSERT INTO `sys_apis` VALUES ('241', '2021-12-22 13:58:40', '2021-12-22 13:58:40', null, '/memUserSafe/excelList', '分页excel导出memUserSafe表列表', 'memUserSafe', 'GET');
INSERT INTO `sys_apis` VALUES ('242', '2021-11-16 11:58:36', '2021-11-16 11:59:00', null, '/imTxim/startOrStopCollect', 'startOrStopCollect', 'imTxim', 'GET');
INSERT INTO `sys_apis` VALUES ('243', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/createImTxFile', '新增imTxFile表', 'imTxFile', 'POST');
INSERT INTO `sys_apis` VALUES ('244', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/deleteImTxFile', '删除imTxFile表', 'imTxFile', 'DELETE');
INSERT INTO `sys_apis` VALUES ('245', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/deleteImTxFileByIds', '批量删除imTxFile表', 'imTxFile', 'DELETE');
INSERT INTO `sys_apis` VALUES ('246', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/updateImTxFile', '更新imTxFile表', 'imTxFile', 'PUT');
INSERT INTO `sys_apis` VALUES ('247', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/findImTxFile', '根据ID获取imTxFile表', 'imTxFile', 'GET');
INSERT INTO `sys_apis` VALUES ('248', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/getImTxFileList', '获取imTxFile表列表', 'imTxFile', 'GET');
INSERT INTO `sys_apis` VALUES ('249', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/quickEdit', '快速编辑imTxFile表', 'imTxFile', 'POST');
INSERT INTO `sys_apis` VALUES ('250', '2021-12-27 10:46:58', '2021-12-27 10:46:58', null, '/imTxFile/excelList', '分页excel导出imTxFile表列表', 'imTxFile', 'GET');
INSERT INTO `sys_apis` VALUES ('251', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/createSysSuperBuilderHistories', '新增sysSuperBuilderHistories表', 'sysSuperBuilderHistories', 'POST');
INSERT INTO `sys_apis` VALUES ('252', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/deleteSysSuperBuilderHistories', '删除sysSuperBuilderHistories表', 'sysSuperBuilderHistories', 'DELETE');
INSERT INTO `sys_apis` VALUES ('253', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/deleteSysSuperBuilderHistoriesByIds', '批量删除sysSuperBuilderHistories表', 'sysSuperBuilderHistories', 'DELETE');
INSERT INTO `sys_apis` VALUES ('254', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/updateSysSuperBuilderHistories', '更新sysSuperBuilderHistories表', 'sysSuperBuilderHistories', 'PUT');
INSERT INTO `sys_apis` VALUES ('255', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/findSysSuperBuilderHistories', '根据ID获取sysSuperBuilderHistories表', 'sysSuperBuilderHistories', 'GET');
INSERT INTO `sys_apis` VALUES ('256', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/getSysSuperBuilderHistoriesList', '获取sysSuperBuilderHistories表列表', 'sysSuperBuilderHistories', 'GET');
INSERT INTO `sys_apis` VALUES ('257', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/quickEdit', '快速编辑sysSuperBuilderHistories表', 'sysSuperBuilderHistories', 'POST');
INSERT INTO `sys_apis` VALUES ('258', '2022-01-04 14:06:11', '2022-01-04 14:06:11', null, '/sysSuperBuilderHistories/excelList', '分页excel导出sysSuperBuilderHistories表列表', 'sysSuperBuilderHistories', 'GET');
INSERT INTO `sys_apis` VALUES ('259', '2022-01-04 14:08:06', '2022-01-04 14:08:06', null, '/memUserLog/createMemUserLog', '新增memUserLog表', 'memUserLog', 'POST');
INSERT INTO `sys_apis` VALUES ('260', '2022-01-04 14:08:06', '2022-01-04 14:08:06', null, '/memUserLog/deleteMemUserLog', '删除memUserLog表', 'memUserLog', 'DELETE');
INSERT INTO `sys_apis` VALUES ('261', '2022-01-04 14:08:07', '2022-01-04 14:08:07', null, '/memUserLog/deleteMemUserLogByIds', '批量删除memUserLog表', 'memUserLog', 'DELETE');
INSERT INTO `sys_apis` VALUES ('262', '2022-01-04 14:08:07', '2022-01-04 14:08:07', null, '/memUserLog/updateMemUserLog', '更新memUserLog表', 'memUserLog', 'PUT');
INSERT INTO `sys_apis` VALUES ('263', '2022-01-04 14:08:07', '2022-01-04 14:08:07', null, '/memUserLog/findMemUserLog', '根据ID获取memUserLog表', 'memUserLog', 'GET');
INSERT INTO `sys_apis` VALUES ('264', '2022-01-04 14:08:07', '2022-01-04 14:08:07', null, '/memUserLog/getMemUserLogList', '获取memUserLog表列表', 'memUserLog', 'GET');
INSERT INTO `sys_apis` VALUES ('265', '2022-01-04 14:08:07', '2022-01-04 14:08:07', null, '/memUserLog/quickEdit', '快速编辑memUserLog表', 'memUserLog', 'POST');
INSERT INTO `sys_apis` VALUES ('266', '2022-01-04 14:08:07', '2022-01-04 14:08:07', null, '/memUserLog/excelList', '分页excel导出memUserLog表列表', 'memUserLog', 'GET');
INSERT INTO `sys_apis` VALUES ('267', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/createMemAddress', '新增memAddress表', 'memAddress', 'POST');
INSERT INTO `sys_apis` VALUES ('268', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/deleteMemAddress', '删除memAddress表', 'memAddress', 'DELETE');
INSERT INTO `sys_apis` VALUES ('269', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/deleteMemAddressByIds', '批量删除memAddress表', 'memAddress', 'DELETE');
INSERT INTO `sys_apis` VALUES ('270', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/updateMemAddress', '更新memAddress表', 'memAddress', 'PUT');
INSERT INTO `sys_apis` VALUES ('271', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/findMemAddress', '根据ID获取memAddress表', 'memAddress', 'GET');
INSERT INTO `sys_apis` VALUES ('272', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/getMemAddressList', '获取memAddress表列表', 'memAddress', 'GET');
INSERT INTO `sys_apis` VALUES ('273', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/quickEdit', '快速编辑memAddress表', 'memAddress', 'POST');
INSERT INTO `sys_apis` VALUES ('274', '2022-01-04 14:14:05', '2022-01-04 14:14:05', null, '/memAddress/excelList', '分页excel导出memAddress表列表', 'memAddress', 'GET');
INSERT INTO `sys_apis` VALUES ('275', '2022-01-04 14:39:03', '2022-01-04 14:39:03', null, '/basicArea/createBasicArea', '新增basicArea表', 'basicArea', 'POST');
INSERT INTO `sys_apis` VALUES ('276', '2022-01-04 14:39:03', '2022-01-04 14:39:03', null, '/basicArea/deleteBasicArea', '删除basicArea表', 'basicArea', 'DELETE');
INSERT INTO `sys_apis` VALUES ('277', '2022-01-04 14:39:03', '2022-01-04 14:39:03', null, '/basicArea/deleteBasicAreaByIds', '批量删除basicArea表', 'basicArea', 'DELETE');
INSERT INTO `sys_apis` VALUES ('278', '2022-01-04 14:39:03', '2022-01-04 14:39:03', null, '/basicArea/updateBasicArea', '更新basicArea表', 'basicArea', 'PUT');
INSERT INTO `sys_apis` VALUES ('279', '2022-01-04 14:39:03', '2022-01-04 14:39:03', null, '/basicArea/findBasicArea', '根据ID获取basicArea表', 'basicArea', 'GET');
INSERT INTO `sys_apis` VALUES ('280', '2022-01-04 14:39:03', '2022-01-04 14:39:03', null, '/basicArea/getBasicAreaList', '获取basicArea表列表', 'basicArea', 'GET');
INSERT INTO `sys_apis` VALUES ('281', '2022-01-04 14:39:04', '2022-01-04 14:39:04', null, '/basicArea/quickEdit', '快速编辑basicArea表', 'basicArea', 'POST');
INSERT INTO `sys_apis` VALUES ('282', '2022-01-04 14:39:04', '2022-01-04 14:39:04', null, '/basicArea/excelList', '分页excel导出basicArea表列表', 'basicArea', 'GET');
INSERT INTO `sys_apis` VALUES ('283', '2022-01-13 15:18:21', '2022-01-13 15:18:21', null, '/bbPiContacts/createBbPiContacts', '新增bbPiContacts表', 'bbPiContacts', 'POST');
INSERT INTO `sys_apis` VALUES ('284', '2022-01-13 15:18:21', '2022-01-13 15:18:21', null, '/bbPiContacts/deleteBbPiContacts', '删除bbPiContacts表', 'bbPiContacts', 'DELETE');
INSERT INTO `sys_apis` VALUES ('285', '2022-01-13 15:18:21', '2022-01-13 15:18:21', null, '/bbPiContacts/deleteBbPiContactsByIds', '批量删除bbPiContacts表', 'bbPiContacts', 'DELETE');
INSERT INTO `sys_apis` VALUES ('286', '2022-01-13 15:18:21', '2022-01-13 15:18:21', null, '/bbPiContacts/updateBbPiContacts', '更新bbPiContacts表', 'bbPiContacts', 'PUT');
INSERT INTO `sys_apis` VALUES ('287', '2022-01-13 15:18:21', '2022-01-13 15:18:21', null, '/bbPiContacts/findBbPiContacts', '根据ID获取bbPiContacts表', 'bbPiContacts', 'GET');
INSERT INTO `sys_apis` VALUES ('288', '2022-01-13 15:18:21', '2022-01-13 15:18:21', null, '/bbPiContacts/getBbPiContactsList', '获取bbPiContacts表列表', 'bbPiContacts', 'GET');
INSERT INTO `sys_apis` VALUES ('289', '2022-01-13 15:18:22', '2022-01-13 15:18:22', null, '/bbPiContacts/quickEdit', '快速编辑bbPiContacts表', 'bbPiContacts', 'POST');
INSERT INTO `sys_apis` VALUES ('290', '2022-01-13 15:18:22', '2022-01-13 15:18:22', null, '/bbPiContacts/excelList', '分页excel导出bbPiContacts表列表', 'bbPiContacts', 'GET');
INSERT INTO `sys_apis` VALUES ('291', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/createBbPiDepartment', '新增bbPiDepartment表', 'bbPiDepartment', 'POST');
INSERT INTO `sys_apis` VALUES ('292', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/deleteBbPiDepartment', '删除bbPiDepartment表', 'bbPiDepartment', 'DELETE');
INSERT INTO `sys_apis` VALUES ('293', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/deleteBbPiDepartmentByIds', '批量删除bbPiDepartment表', 'bbPiDepartment', 'DELETE');
INSERT INTO `sys_apis` VALUES ('294', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/updateBbPiDepartment', '更新bbPiDepartment表', 'bbPiDepartment', 'PUT');
INSERT INTO `sys_apis` VALUES ('295', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/findBbPiDepartment', '根据ID获取bbPiDepartment表', 'bbPiDepartment', 'GET');
INSERT INTO `sys_apis` VALUES ('296', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/getBbPiDepartmentList', '获取bbPiDepartment表列表', 'bbPiDepartment', 'GET');
INSERT INTO `sys_apis` VALUES ('297', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/quickEdit', '快速编辑bbPiDepartment表', 'bbPiDepartment', 'POST');
INSERT INTO `sys_apis` VALUES ('298', '2022-01-13 15:19:37', '2022-01-13 15:19:37', null, '/bbPiDepartment/excelList', '分页excel导出bbPiDepartment表列表', 'bbPiDepartment', 'GET');
INSERT INTO `sys_apis` VALUES ('299', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/createBbPiDevice', '新增bbPiDevice表', 'bbPiDevice', 'POST');
INSERT INTO `sys_apis` VALUES ('300', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/deleteBbPiDevice', '删除bbPiDevice表', 'bbPiDevice', 'DELETE');
INSERT INTO `sys_apis` VALUES ('301', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/deleteBbPiDeviceByIds', '批量删除bbPiDevice表', 'bbPiDevice', 'DELETE');
INSERT INTO `sys_apis` VALUES ('302', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/updateBbPiDevice', '更新bbPiDevice表', 'bbPiDevice', 'PUT');
INSERT INTO `sys_apis` VALUES ('303', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/findBbPiDevice', '根据ID获取bbPiDevice表', 'bbPiDevice', 'GET');
INSERT INTO `sys_apis` VALUES ('304', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/getBbPiDeviceList', '获取bbPiDevice表列表', 'bbPiDevice', 'GET');
INSERT INTO `sys_apis` VALUES ('305', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/quickEdit', '快速编辑bbPiDevice表', 'bbPiDevice', 'POST');
INSERT INTO `sys_apis` VALUES ('306', '2022-01-13 16:13:23', '2022-01-13 16:13:23', null, '/bbPiDevice/excelList', '分页excel导出bbPiDevice表列表', 'bbPiDevice', 'GET');
INSERT INTO `sys_apis` VALUES ('307', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/createBbPiInstitution', '新增bbPiInstitution表', 'bbPiInstitution', 'POST');
INSERT INTO `sys_apis` VALUES ('308', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/deleteBbPiInstitution', '删除bbPiInstitution表', 'bbPiInstitution', 'DELETE');
INSERT INTO `sys_apis` VALUES ('309', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/deleteBbPiInstitutionByIds', '批量删除bbPiInstitution表', 'bbPiInstitution', 'DELETE');
INSERT INTO `sys_apis` VALUES ('310', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/updateBbPiInstitution', '更新bbPiInstitution表', 'bbPiInstitution', 'PUT');
INSERT INTO `sys_apis` VALUES ('311', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/findBbPiInstitution', '根据ID获取bbPiInstitution表', 'bbPiInstitution', 'GET');
INSERT INTO `sys_apis` VALUES ('312', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/getBbPiInstitutionList', '获取bbPiInstitution表列表', 'bbPiInstitution', 'GET');
INSERT INTO `sys_apis` VALUES ('313', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/quickEdit', '快速编辑bbPiInstitution表', 'bbPiInstitution', 'POST');
INSERT INTO `sys_apis` VALUES ('314', '2022-01-13 16:13:55', '2022-01-13 16:13:55', null, '/bbPiInstitution/excelList', '分页excel导出bbPiInstitution表列表', 'bbPiInstitution', 'GET');
INSERT INTO `sys_apis` VALUES ('315', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/createBbPiInstitutionBusiness', '新增bbPiInstitutionBusiness表', 'bbPiInstitutionBusiness', 'POST');
INSERT INTO `sys_apis` VALUES ('316', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusiness', '删除bbPiInstitutionBusiness表', 'bbPiInstitutionBusiness', 'DELETE');
INSERT INTO `sys_apis` VALUES ('317', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusinessByIds', '批量删除bbPiInstitutionBusiness表', 'bbPiInstitutionBusiness', 'DELETE');
INSERT INTO `sys_apis` VALUES ('318', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/updateBbPiInstitutionBusiness', '更新bbPiInstitutionBusiness表', 'bbPiInstitutionBusiness', 'PUT');
INSERT INTO `sys_apis` VALUES ('319', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/findBbPiInstitutionBusiness', '根据ID获取bbPiInstitutionBusiness表', 'bbPiInstitutionBusiness', 'GET');
INSERT INTO `sys_apis` VALUES ('320', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/getBbPiInstitutionBusinessList', '获取bbPiInstitutionBusiness表列表', 'bbPiInstitutionBusiness', 'GET');
INSERT INTO `sys_apis` VALUES ('321', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/quickEdit', '快速编辑bbPiInstitutionBusiness表', 'bbPiInstitutionBusiness', 'POST');
INSERT INTO `sys_apis` VALUES ('322', '2022-01-13 16:14:25', '2022-01-13 16:14:25', null, '/bbPiInstitutionBusiness/excelList', '分页excel导出bbPiInstitutionBusiness表列表', 'bbPiInstitutionBusiness', 'GET');
INSERT INTO `sys_apis` VALUES ('323', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/createBbPiPerson', '新增bbPiPerson表', 'bbPiPerson', 'POST');
INSERT INTO `sys_apis` VALUES ('324', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/deleteBbPiPerson', '删除bbPiPerson表', 'bbPiPerson', 'DELETE');
INSERT INTO `sys_apis` VALUES ('325', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/deleteBbPiPersonByIds', '批量删除bbPiPerson表', 'bbPiPerson', 'DELETE');
INSERT INTO `sys_apis` VALUES ('326', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/updateBbPiPerson', '更新bbPiPerson表', 'bbPiPerson', 'PUT');
INSERT INTO `sys_apis` VALUES ('327', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/findBbPiPerson', '根据ID获取bbPiPerson表', 'bbPiPerson', 'GET');
INSERT INTO `sys_apis` VALUES ('328', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/getBbPiPersonList', '获取bbPiPerson表列表', 'bbPiPerson', 'GET');
INSERT INTO `sys_apis` VALUES ('329', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/quickEdit', '快速编辑bbPiPerson表', 'bbPiPerson', 'POST');
INSERT INTO `sys_apis` VALUES ('330', '2022-01-13 16:15:15', '2022-01-13 16:15:15', null, '/bbPiPerson/excelList', '分页excel导出bbPiPerson表列表', 'bbPiPerson', 'GET');
INSERT INTO `sys_apis` VALUES ('331', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/createBbPiServicePoint', '新增bbPiServicePoint表', 'bbPiServicePoint', 'POST');
INSERT INTO `sys_apis` VALUES ('332', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/deleteBbPiServicePoint', '删除bbPiServicePoint表', 'bbPiServicePoint', 'DELETE');
INSERT INTO `sys_apis` VALUES ('333', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/deleteBbPiServicePointByIds', '批量删除bbPiServicePoint表', 'bbPiServicePoint', 'DELETE');
INSERT INTO `sys_apis` VALUES ('334', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/updateBbPiServicePoint', '更新bbPiServicePoint表', 'bbPiServicePoint', 'PUT');
INSERT INTO `sys_apis` VALUES ('335', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/findBbPiServicePoint', '根据ID获取bbPiServicePoint表', 'bbPiServicePoint', 'GET');
INSERT INTO `sys_apis` VALUES ('336', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/getBbPiServicePointList', '获取bbPiServicePoint表列表', 'bbPiServicePoint', 'GET');
INSERT INTO `sys_apis` VALUES ('337', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/quickEdit', '快速编辑bbPiServicePoint表', 'bbPiServicePoint', 'POST');
INSERT INTO `sys_apis` VALUES ('338', '2022-01-13 16:15:41', '2022-01-13 16:15:41', null, '/bbPiServicePoint/excelList', '分页excel导出bbPiServicePoint表列表', 'bbPiServicePoint', 'GET');
INSERT INTO `sys_apis` VALUES ('339', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/createBbPiStaff', '新增bbPiStaff表', 'bbPiStaff', 'POST');
INSERT INTO `sys_apis` VALUES ('340', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/deleteBbPiStaff', '删除bbPiStaff表', 'bbPiStaff', 'DELETE');
INSERT INTO `sys_apis` VALUES ('341', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/deleteBbPiStaffByIds', '批量删除bbPiStaff表', 'bbPiStaff', 'DELETE');
INSERT INTO `sys_apis` VALUES ('342', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/updateBbPiStaff', '更新bbPiStaff表', 'bbPiStaff', 'PUT');
INSERT INTO `sys_apis` VALUES ('343', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/findBbPiStaff', '根据ID获取bbPiStaff表', 'bbPiStaff', 'GET');
INSERT INTO `sys_apis` VALUES ('344', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/getBbPiStaffList', '获取bbPiStaff表列表', 'bbPiStaff', 'GET');
INSERT INTO `sys_apis` VALUES ('345', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/quickEdit', '快速编辑bbPiStaff表', 'bbPiStaff', 'POST');
INSERT INTO `sys_apis` VALUES ('346', '2022-01-13 16:16:12', '2022-01-13 16:16:12', null, '/bbPiStaff/excelList', '分页excel导出bbPiStaff表列表', 'bbPiStaff', 'GET');
INSERT INTO `sys_apis` VALUES ('347', '2022-01-20 14:24:14', '2022-01-20 14:24:14', null, '/bbPiTreatmentDiagnose/createBbPiTreatmentDiagnose', '新增bbPiTreatmentDiagnose表', 'bbPiTreatmentDiagnose', 'POST');
INSERT INTO `sys_apis` VALUES ('348', '2022-01-20 14:24:14', '2022-01-20 14:24:14', null, '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnose', '删除bbPiTreatmentDiagnose表', 'bbPiTreatmentDiagnose', 'DELETE');
INSERT INTO `sys_apis` VALUES ('349', '2022-01-20 14:24:14', '2022-01-20 14:24:14', null, '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnoseByIds', '批量删除bbPiTreatmentDiagnose表', 'bbPiTreatmentDiagnose', 'DELETE');
INSERT INTO `sys_apis` VALUES ('350', '2022-01-20 14:24:14', '2022-01-20 14:24:14', null, '/bbPiTreatmentDiagnose/updateBbPiTreatmentDiagnose', '更新bbPiTreatmentDiagnose表', 'bbPiTreatmentDiagnose', 'PUT');
INSERT INTO `sys_apis` VALUES ('351', '2022-01-20 14:24:14', '2022-01-20 14:24:14', null, '/bbPiTreatmentDiagnose/findBbPiTreatmentDiagnose', '根据ID获取bbPiTreatmentDiagnose表', 'bbPiTreatmentDiagnose', 'GET');
INSERT INTO `sys_apis` VALUES ('352', '2022-01-20 14:24:15', '2022-01-20 14:24:15', null, '/bbPiTreatmentDiagnose/getBbPiTreatmentDiagnoseList', '获取bbPiTreatmentDiagnose表列表', 'bbPiTreatmentDiagnose', 'GET');
INSERT INTO `sys_apis` VALUES ('353', '2022-01-20 14:24:15', '2022-01-20 14:24:15', null, '/bbPiTreatmentDiagnose/quickEdit', '快速编辑bbPiTreatmentDiagnose表', 'bbPiTreatmentDiagnose', 'POST');
INSERT INTO `sys_apis` VALUES ('354', '2022-01-20 14:24:15', '2022-01-20 14:24:15', null, '/bbPiTreatmentDiagnose/excelList', '分页excel导出bbPiTreatmentDiagnose表列表', 'bbPiTreatmentDiagnose', 'GET');
INSERT INTO `sys_apis` VALUES ('355', '2022-01-20 14:30:12', '2022-01-20 14:30:12', null, '/bbPiTreatmentOrder/createBbPiTreatmentOrder', '新增bbPiTreatmentOrder表', 'bbPiTreatmentOrder', 'POST');
INSERT INTO `sys_apis` VALUES ('356', '2022-01-20 14:30:12', '2022-01-20 14:30:12', null, '/bbPiTreatmentOrder/deleteBbPiTreatmentOrder', '删除bbPiTreatmentOrder表', 'bbPiTreatmentOrder', 'DELETE');
INSERT INTO `sys_apis` VALUES ('357', '2022-01-20 14:30:12', '2022-01-20 14:30:12', null, '/bbPiTreatmentOrder/deleteBbPiTreatmentOrderByIds', '批量删除bbPiTreatmentOrder表', 'bbPiTreatmentOrder', 'DELETE');
INSERT INTO `sys_apis` VALUES ('358', '2022-01-20 14:30:13', '2022-01-20 14:30:13', null, '/bbPiTreatmentOrder/updateBbPiTreatmentOrder', '更新bbPiTreatmentOrder表', 'bbPiTreatmentOrder', 'PUT');
INSERT INTO `sys_apis` VALUES ('359', '2022-01-20 14:30:13', '2022-01-20 14:30:13', null, '/bbPiTreatmentOrder/findBbPiTreatmentOrder', '根据ID获取bbPiTreatmentOrder表', 'bbPiTreatmentOrder', 'GET');
INSERT INTO `sys_apis` VALUES ('360', '2022-01-20 14:30:13', '2022-01-20 14:30:13', null, '/bbPiTreatmentOrder/getBbPiTreatmentOrderList', '获取bbPiTreatmentOrder表列表', 'bbPiTreatmentOrder', 'GET');
INSERT INTO `sys_apis` VALUES ('361', '2022-01-20 14:30:13', '2022-01-20 14:30:13', null, '/bbPiTreatmentOrder/quickEdit', '快速编辑bbPiTreatmentOrder表', 'bbPiTreatmentOrder', 'POST');
INSERT INTO `sys_apis` VALUES ('362', '2022-01-20 14:30:14', '2022-01-20 14:30:14', null, '/bbPiTreatmentOrder/excelList', '分页excel导出bbPiTreatmentOrder表列表', 'bbPiTreatmentOrder', 'GET');
INSERT INTO `sys_apis` VALUES ('363', '2022-01-20 14:31:01', '2022-01-20 14:31:01', null, '/bbPiTreatmentRecord/createBbPiTreatmentRecord', '新增bbPiTreatmentRecord表', 'bbPiTreatmentRecord', 'POST');
INSERT INTO `sys_apis` VALUES ('364', '2022-01-20 14:31:02', '2022-01-20 14:31:02', null, '/bbPiTreatmentRecord/deleteBbPiTreatmentRecord', '删除bbPiTreatmentRecord表', 'bbPiTreatmentRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES ('365', '2022-01-20 14:31:02', '2022-01-20 14:31:02', null, '/bbPiTreatmentRecord/deleteBbPiTreatmentRecordByIds', '批量删除bbPiTreatmentRecord表', 'bbPiTreatmentRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES ('366', '2022-01-20 14:31:02', '2022-01-20 14:31:02', null, '/bbPiTreatmentRecord/updateBbPiTreatmentRecord', '更新bbPiTreatmentRecord表', 'bbPiTreatmentRecord', 'PUT');
INSERT INTO `sys_apis` VALUES ('367', '2022-01-20 14:31:02', '2022-01-20 14:31:02', null, '/bbPiTreatmentRecord/findBbPiTreatmentRecord', '根据ID获取bbPiTreatmentRecord表', 'bbPiTreatmentRecord', 'GET');
INSERT INTO `sys_apis` VALUES ('368', '2022-01-20 14:31:03', '2022-01-20 14:31:03', null, '/bbPiTreatmentRecord/getBbPiTreatmentRecordList', '获取bbPiTreatmentRecord表列表', 'bbPiTreatmentRecord', 'GET');
INSERT INTO `sys_apis` VALUES ('369', '2022-01-20 14:31:03', '2022-01-20 14:31:03', null, '/bbPiTreatmentRecord/quickEdit', '快速编辑bbPiTreatmentRecord表', 'bbPiTreatmentRecord', 'POST');
INSERT INTO `sys_apis` VALUES ('370', '2022-01-20 14:31:03', '2022-01-20 14:31:03', null, '/bbPiTreatmentRecord/excelList', '分页excel导出bbPiTreatmentRecord表列表', 'bbPiTreatmentRecord', 'GET');
INSERT INTO `sys_apis` VALUES ('371', '2022-01-20 14:32:08', '2022-01-20 14:32:08', null, '/bbPiTreatmentReferral/createBbPiTreatmentReferral', '新增bbPiTreatmentReferral表', 'bbPiTreatmentReferral', 'POST');
INSERT INTO `sys_apis` VALUES ('372', '2022-01-20 14:32:08', '2022-01-20 14:32:08', null, '/bbPiTreatmentReferral/deleteBbPiTreatmentReferral', '删除bbPiTreatmentReferral表', 'bbPiTreatmentReferral', 'DELETE');
INSERT INTO `sys_apis` VALUES ('373', '2022-01-20 14:32:08', '2022-01-20 14:32:08', null, '/bbPiTreatmentReferral/deleteBbPiTreatmentReferralByIds', '批量删除bbPiTreatmentReferral表', 'bbPiTreatmentReferral', 'DELETE');
INSERT INTO `sys_apis` VALUES ('374', '2022-01-20 14:32:09', '2022-01-20 14:32:09', null, '/bbPiTreatmentReferral/updateBbPiTreatmentReferral', '更新bbPiTreatmentReferral表', 'bbPiTreatmentReferral', 'PUT');
INSERT INTO `sys_apis` VALUES ('375', '2022-01-20 14:32:09', '2022-01-20 14:32:09', null, '/bbPiTreatmentReferral/findBbPiTreatmentReferral', '根据ID获取bbPiTreatmentReferral表', 'bbPiTreatmentReferral', 'GET');
INSERT INTO `sys_apis` VALUES ('376', '2022-01-20 14:32:09', '2022-01-20 14:32:09', null, '/bbPiTreatmentReferral/getBbPiTreatmentReferralList', '获取bbPiTreatmentReferral表列表', 'bbPiTreatmentReferral', 'GET');
INSERT INTO `sys_apis` VALUES ('377', '2022-01-20 14:32:09', '2022-01-20 14:32:09', null, '/bbPiTreatmentReferral/quickEdit', '快速编辑bbPiTreatmentReferral表', 'bbPiTreatmentReferral', 'POST');
INSERT INTO `sys_apis` VALUES ('378', '2022-01-20 14:32:09', '2022-01-20 14:32:09', null, '/bbPiTreatmentReferral/excelList', '分页excel导出bbPiTreatmentReferral表列表', 'bbPiTreatmentReferral', 'GET');
INSERT INTO `sys_apis` VALUES ('379', '2022-01-20 14:33:38', '2022-01-20 14:33:38', null, '/bbPiUpField/createBbPiUpField', '新增bbPiUpField表', 'bbPiUpField', 'POST');
INSERT INTO `sys_apis` VALUES ('380', '2022-01-20 14:33:39', '2022-01-20 14:33:39', null, '/bbPiUpField/deleteBbPiUpField', '删除bbPiUpField表', 'bbPiUpField', 'DELETE');
INSERT INTO `sys_apis` VALUES ('381', '2022-01-20 14:33:39', '2022-01-20 14:33:39', null, '/bbPiUpField/deleteBbPiUpFieldByIds', '批量删除bbPiUpField表', 'bbPiUpField', 'DELETE');
INSERT INTO `sys_apis` VALUES ('382', '2022-01-20 14:33:39', '2022-01-20 14:33:39', null, '/bbPiUpField/updateBbPiUpField', '更新bbPiUpField表', 'bbPiUpField', 'PUT');
INSERT INTO `sys_apis` VALUES ('383', '2022-01-20 14:33:39', '2022-01-20 14:33:39', null, '/bbPiUpField/findBbPiUpField', '根据ID获取bbPiUpField表', 'bbPiUpField', 'GET');
INSERT INTO `sys_apis` VALUES ('384', '2022-01-20 14:33:39', '2022-01-20 14:33:39', null, '/bbPiUpField/getBbPiUpFieldList', '获取bbPiUpField表列表', 'bbPiUpField', 'GET');
INSERT INTO `sys_apis` VALUES ('385', '2022-01-20 14:33:39', '2022-01-20 14:33:39', null, '/bbPiUpField/quickEdit', '快速编辑bbPiUpField表', 'bbPiUpField', 'POST');
INSERT INTO `sys_apis` VALUES ('386', '2022-01-20 14:33:40', '2022-01-20 14:33:40', null, '/bbPiUpField/excelList', '分页excel导出bbPiUpField表列表', 'bbPiUpField', 'GET');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2021-09-10 14:08:19', '2022-01-20 15:43:32', null, '888', '普通用户', '0', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2021-09-10 14:08:19', '2022-01-20 15:44:35', null, '8881', '普通用户子角色', '888', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '9528', '测试角色', '0', 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES ('1', '888');
INSERT INTO `sys_authority_menus` VALUES ('1', '8881');
INSERT INTO `sys_authority_menus` VALUES ('1', '9528');
INSERT INTO `sys_authority_menus` VALUES ('2', '888');
INSERT INTO `sys_authority_menus` VALUES ('2', '8881');
INSERT INTO `sys_authority_menus` VALUES ('2', '9528');
INSERT INTO `sys_authority_menus` VALUES ('3', '888');
INSERT INTO `sys_authority_menus` VALUES ('3', '8881');
INSERT INTO `sys_authority_menus` VALUES ('3', '9528');
INSERT INTO `sys_authority_menus` VALUES ('4', '888');
INSERT INTO `sys_authority_menus` VALUES ('4', '8881');
INSERT INTO `sys_authority_menus` VALUES ('4', '9528');
INSERT INTO `sys_authority_menus` VALUES ('5', '888');
INSERT INTO `sys_authority_menus` VALUES ('5', '8881');
INSERT INTO `sys_authority_menus` VALUES ('5', '9528');
INSERT INTO `sys_authority_menus` VALUES ('6', '888');
INSERT INTO `sys_authority_menus` VALUES ('6', '8881');
INSERT INTO `sys_authority_menus` VALUES ('6', '9528');
INSERT INTO `sys_authority_menus` VALUES ('7', '888');
INSERT INTO `sys_authority_menus` VALUES ('7', '8881');
INSERT INTO `sys_authority_menus` VALUES ('7', '9528');
INSERT INTO `sys_authority_menus` VALUES ('8', '888');
INSERT INTO `sys_authority_menus` VALUES ('8', '8881');
INSERT INTO `sys_authority_menus` VALUES ('8', '9528');
INSERT INTO `sys_authority_menus` VALUES ('9', '888');
INSERT INTO `sys_authority_menus` VALUES ('9', '8881');
INSERT INTO `sys_authority_menus` VALUES ('9', '9528');
INSERT INTO `sys_authority_menus` VALUES ('10', '888');
INSERT INTO `sys_authority_menus` VALUES ('10', '8881');
INSERT INTO `sys_authority_menus` VALUES ('10', '9528');
INSERT INTO `sys_authority_menus` VALUES ('11', '888');
INSERT INTO `sys_authority_menus` VALUES ('11', '8881');
INSERT INTO `sys_authority_menus` VALUES ('11', '9528');
INSERT INTO `sys_authority_menus` VALUES ('12', '888');
INSERT INTO `sys_authority_menus` VALUES ('12', '8881');
INSERT INTO `sys_authority_menus` VALUES ('12', '9528');
INSERT INTO `sys_authority_menus` VALUES ('14', '888');
INSERT INTO `sys_authority_menus` VALUES ('14', '8881');
INSERT INTO `sys_authority_menus` VALUES ('14', '9528');
INSERT INTO `sys_authority_menus` VALUES ('15', '888');
INSERT INTO `sys_authority_menus` VALUES ('15', '8881');
INSERT INTO `sys_authority_menus` VALUES ('15', '9528');
INSERT INTO `sys_authority_menus` VALUES ('16', '888');
INSERT INTO `sys_authority_menus` VALUES ('16', '8881');
INSERT INTO `sys_authority_menus` VALUES ('16', '9528');
INSERT INTO `sys_authority_menus` VALUES ('17', '888');
INSERT INTO `sys_authority_menus` VALUES ('17', '8881');
INSERT INTO `sys_authority_menus` VALUES ('17', '9528');
INSERT INTO `sys_authority_menus` VALUES ('18', '888');
INSERT INTO `sys_authority_menus` VALUES ('18', '8881');
INSERT INTO `sys_authority_menus` VALUES ('19', '888');
INSERT INTO `sys_authority_menus` VALUES ('19', '8881');
INSERT INTO `sys_authority_menus` VALUES ('20', '888');
INSERT INTO `sys_authority_menus` VALUES ('23', '888');
INSERT INTO `sys_authority_menus` VALUES ('23', '8881');
INSERT INTO `sys_authority_menus` VALUES ('24', '888');
INSERT INTO `sys_authority_menus` VALUES ('24', '8881');
INSERT INTO `sys_authority_menus` VALUES ('25', '888');
INSERT INTO `sys_authority_menus` VALUES ('25', '8881');
INSERT INTO `sys_authority_menus` VALUES ('26', '888');
INSERT INTO `sys_authority_menus` VALUES ('27', '888');
INSERT INTO `sys_authority_menus` VALUES ('28', '888');
INSERT INTO `sys_authority_menus` VALUES ('29', '888');
INSERT INTO `sys_authority_menus` VALUES ('30', '888');
INSERT INTO `sys_authority_menus` VALUES ('31', '888');
INSERT INTO `sys_authority_menus` VALUES ('32', '888');
INSERT INTO `sys_authority_menus` VALUES ('33', '888');
INSERT INTO `sys_authority_menus` VALUES ('34', '888');
INSERT INTO `sys_authority_menus` VALUES ('35', '888');
INSERT INTO `sys_authority_menus` VALUES ('36', '888');
INSERT INTO `sys_authority_menus` VALUES ('37', '888');
INSERT INTO `sys_authority_menus` VALUES ('38', '888');
INSERT INTO `sys_authority_menus` VALUES ('39', '888');
INSERT INTO `sys_authority_menus` VALUES ('40', '888');
INSERT INTO `sys_authority_menus` VALUES ('41', '888');
INSERT INTO `sys_authority_menus` VALUES ('42', '888');
INSERT INTO `sys_authority_menus` VALUES ('43', '888');
INSERT INTO `sys_authority_menus` VALUES ('44', '888');
INSERT INTO `sys_authority_menus` VALUES ('45', '888');
INSERT INTO `sys_authority_menus` VALUES ('45', '8881');
INSERT INTO `sys_authority_menus` VALUES ('46', '888');
INSERT INTO `sys_authority_menus` VALUES ('47', '888');
INSERT INTO `sys_authority_menus` VALUES ('48', '888');
INSERT INTO `sys_authority_menus` VALUES ('49', '888');
INSERT INTO `sys_authority_menus` VALUES ('50', '888');
INSERT INTO `sys_authority_menus` VALUES ('50', '8881');
INSERT INTO `sys_authority_menus` VALUES ('51', '888');
INSERT INTO `sys_authority_menus` VALUES ('52', '888');
INSERT INTO `sys_authority_menus` VALUES ('53', '888');
INSERT INTO `sys_authority_menus` VALUES ('54', '888');
INSERT INTO `sys_authority_menus` VALUES ('54', '8881');
INSERT INTO `sys_authority_menus` VALUES ('55', '888');
INSERT INTO `sys_authority_menus` VALUES ('55', '8881');
INSERT INTO `sys_authority_menus` VALUES ('56', '888');
INSERT INTO `sys_authority_menus` VALUES ('56', '8881');
INSERT INTO `sys_authority_menus` VALUES ('57', '888');
INSERT INTO `sys_authority_menus` VALUES ('57', '8881');
INSERT INTO `sys_authority_menus` VALUES ('58', '888');
INSERT INTO `sys_authority_menus` VALUES ('58', '8881');
INSERT INTO `sys_authority_menus` VALUES ('59', '888');
INSERT INTO `sys_authority_menus` VALUES ('59', '8881');
INSERT INTO `sys_authority_menus` VALUES ('60', '888');
INSERT INTO `sys_authority_menus` VALUES ('60', '8881');
INSERT INTO `sys_authority_menus` VALUES ('61', '888');
INSERT INTO `sys_authority_menus` VALUES ('61', '8881');
INSERT INTO `sys_authority_menus` VALUES ('62', '888');
INSERT INTO `sys_authority_menus` VALUES ('62', '8881');
INSERT INTO `sys_authority_menus` VALUES ('63', '888');
INSERT INTO `sys_authority_menus` VALUES ('63', '8881');
INSERT INTO `sys_authority_menus` VALUES ('64', '888');
INSERT INTO `sys_authority_menus` VALUES ('64', '8881');
INSERT INTO `sys_authority_menus` VALUES ('65', '888');
INSERT INTO `sys_authority_menus` VALUES ('65', '8881');
INSERT INTO `sys_authority_menus` VALUES ('66', '888');
INSERT INTO `sys_authority_menus` VALUES ('66', '8881');
INSERT INTO `sys_authority_menus` VALUES ('67', '888');
INSERT INTO `sys_authority_menus` VALUES ('67', '8881');
INSERT INTO `sys_authority_menus` VALUES ('68', '888');
INSERT INTO `sys_authority_menus` VALUES ('68', '8881');
INSERT INTO `sys_authority_menus` VALUES ('69', '888');
INSERT INTO `sys_authority_menus` VALUES ('69', '8881');
INSERT INTO `sys_authority_menus` VALUES ('70', '888');
INSERT INTO `sys_authority_menus` VALUES ('70', '8881');
INSERT INTO `sys_authority_menus` VALUES ('71', '888');
INSERT INTO `sys_authority_menus` VALUES ('71', '8881');
INSERT INTO `sys_authority_menus` VALUES ('72', '888');
INSERT INTO `sys_authority_menus` VALUES ('72', '8881');
INSERT INTO `sys_authority_menus` VALUES ('73', '888');
INSERT INTO `sys_authority_menus` VALUES ('73', '8881');
INSERT INTO `sys_authority_menus` VALUES ('74', '888');
INSERT INTO `sys_authority_menus` VALUES ('74', '8881');
INSERT INTO `sys_authority_menus` VALUES ('75', '888');
INSERT INTO `sys_authority_menus` VALUES ('75', '8881');
INSERT INTO `sys_authority_menus` VALUES ('76', '888');
INSERT INTO `sys_authority_menus` VALUES ('76', '8881');
INSERT INTO `sys_authority_menus` VALUES ('77', '888');
INSERT INTO `sys_authority_menus` VALUES ('77', '8881');
INSERT INTO `sys_authority_menus` VALUES ('78', '888');
INSERT INTO `sys_authority_menus` VALUES ('78', '8881');
INSERT INTO `sys_authority_menus` VALUES ('79', '888');
INSERT INTO `sys_authority_menus` VALUES ('79', '8881');
INSERT INTO `sys_authority_menus` VALUES ('80', '888');
INSERT INTO `sys_authority_menus` VALUES ('80', '8881');

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES ('1', '2021-09-10 14:08:19', '2021-12-22 14:18:39', null, '0', '0', 'dashboard', 'dashboard', '0', 'view/dashboard/index.vue', '0', '0', '0', '仪表盘', 'share', '0');
INSERT INTO `sys_base_menus` VALUES ('2', '2021-09-10 14:08:19', '2021-12-20 17:29:16', null, '0', '14', 'about', 'about', '0', 'view/about/index.vue', '7', '0', '0', '关于我们', 'info', '0');
INSERT INTO `sys_base_menus` VALUES ('3', '2021-09-10 14:08:19', '2021-12-22 14:18:16', null, '0', '0', 'admin', 'superAdmin', '0', 'view/superAdmin/index.vue', '100', '0', '0', '系统管理', 's-tools', '0');
INSERT INTO `sys_base_menus` VALUES ('4', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '3', 'authority', 'authority', '0', 'view/superAdmin/authority/authority.vue', '1', '0', '0', '角色管理', 's-custom', '0');
INSERT INTO `sys_base_menus` VALUES ('5', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '3', 'menu', 'menu', '0', 'view/superAdmin/menu/menu.vue', '2', '1', '0', '菜单管理', 's-order', '0');
INSERT INTO `sys_base_menus` VALUES ('6', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '3', 'api', 'api', '0', 'view/superAdmin/api/api.vue', '3', '1', '0', 'api管理', 's-platform', '0');
INSERT INTO `sys_base_menus` VALUES ('7', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '3', 'user', 'user', '0', 'view/superAdmin/user/user.vue', '4', '0', '0', '用户管理', 'coordinate', '0');
INSERT INTO `sys_base_menus` VALUES ('8', '2021-09-10 14:08:19', '2022-01-04 15:59:21', null, '0', '0', 'person', 'person', '1', 'view/person/person.vue', '999', '0', '0', '个人信息', 'message-solid', '0');
INSERT INTO `sys_base_menus` VALUES ('9', '2021-09-10 14:08:19', '2022-01-13 17:01:48', null, '0', '0', 'im', 'im', '0', 'view/im/index.vue', '6', '0', '0', '即时通信', 'chat-line-square', '0');
INSERT INTO `sys_base_menus` VALUES ('10', '2021-09-10 14:08:19', '2021-12-20 14:46:32', null, '0', '9', 'imTxMsg', 'imTxMsg', '0', 'view/im/imTxMsg/imTxMsg.vue', '4', '0', '0', '腾讯IM消息', 's-marketing', '0');
INSERT INTO `sys_base_menus` VALUES ('11', '2021-09-10 14:08:19', '2021-12-20 15:02:30', null, '0', '9', 'imTxMsgForm/:id', 'imTxMsgForm', '1', 'view/im/imTxMsg/imTxMsgForm.vue', '5', '0', '0', '腾讯IM消息详情', 'upload', '0');
INSERT INTO `sys_base_menus` VALUES ('12', '2021-09-10 14:08:19', '2021-12-20 14:50:11', null, '0', '9', 'imTxMsgFile', 'imTxMsgFile', '0', 'view/im/imTxMsgFile/imTxMsgFile.vue', '6', '0', '0', '腾讯消息文件', 'upload', '0');
INSERT INTO `sys_base_menus` VALUES ('13', '2021-09-10 14:08:19', '2021-09-10 14:08:19', '2021-12-20 14:50:19', '0', '9', 'customer', 'customer', '0', 'view/example/customer/customer.vue', '7', '0', '0', '客户列表（资源示例）', 's-custom', '0');
INSERT INTO `sys_base_menus` VALUES ('14', '2021-09-10 14:08:19', '2021-12-21 13:34:17', null, '0', '0', 'systemTools', 'systemTools', '0', 'view/systemTools/index.vue', '101', '0', '0', '系统工具', 's-cooperation', '0');
INSERT INTO `sys_base_menus` VALUES ('15', '2021-09-10 14:08:19', '2021-12-21 10:52:48', null, '0', '14', 'superBuilder', 'superBuilder', '0', 'view/systemTools/superBuilder/index.vue', '1', '0', '0', '代码生成器', 'cpu', '0');
INSERT INTO `sys_base_menus` VALUES ('16', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '14', 'formCreate', 'formCreate', '0', 'view/systemTools/formCreate/index.vue', '2', '1', '0', '表单生成器', 'magic-stick', '0');
INSERT INTO `sys_base_menus` VALUES ('17', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '14', 'system', 'system', '0', 'view/systemTools/system/system.vue', '3', '0', '0', '系统配置', 's-operation', '0');
INSERT INTO `sys_base_menus` VALUES ('18', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '3', 'dictionary', 'dictionary', '0', 'view/superAdmin/dictionary/sysDictionary.vue', '5', '0', '0', '字典管理', 'notebook-2', '0');
INSERT INTO `sys_base_menus` VALUES ('19', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '3', 'dictionaryDetail/:id', 'dictionaryDetail', '1', 'view/superAdmin/dictionary/sysDictionaryDetail.vue', '1', '0', '0', '字典详情', 's-order', '0');
INSERT INTO `sys_base_menus` VALUES ('20', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '3', 'operation', 'operation', '0', 'view/superAdmin/operation/sysOperationRecord.vue', '6', '0', '0', '操作历史', 'time', '0');
INSERT INTO `sys_base_menus` VALUES ('22', '2021-09-10 14:08:19', '2021-09-24 11:11:40', '2021-11-15 15:42:38', '0', '0', 'https://www.88act.com', 'https://www.88act.com', '0', '/', '0', '0', '0', '官方网站', 's-home', '0');
INSERT INTO `sys_base_menus` VALUES ('23', '2021-09-10 14:08:19', '2021-12-20 17:29:00', null, '0', '14', 'state', 'state', '0', 'view/system/state.vue', '6', '0', '0', '服务器状态', 'cloudy', '0');
INSERT INTO `sys_base_menus` VALUES ('24', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '14', 'superBuilderAdmin', 'superBuilderAdmin', '0', 'view/systemTools/superBuilderAdmin/index.vue', '1', '0', '0', '代码生成历史', 's-finance', '0');
INSERT INTO `sys_base_menus` VALUES ('25', '2021-09-10 14:08:19', '2021-09-10 14:08:19', null, '0', '14', 'superBuilderEdit/:id', 'superBuilderEdit', '1', 'view/systemTools/superBuilder/index.vue', '0', '0', '0', '代码生成[复用]', 's-finance', '0');
INSERT INTO `sys_base_menus` VALUES ('26', '2021-07-16 15:10:11', '2021-12-21 13:34:37', null, '0', '0', 'cms', 'cms', '0', 'view/cms/index.vue', '1', '0', '0', '内容管理系统', 's-goods', '0');
INSERT INTO `sys_base_menus` VALUES ('27', '2021-07-16 15:10:58', '2021-08-20 14:14:46', null, '0', '26', 'cmsAdSeat', 'cmsAdSeat', '0', 'view/cms/cmsAdSeat/cmsAdSeat.vue', '0', '0', '0', '广告位管理', 'setting', '0');
INSERT INTO `sys_base_menus` VALUES ('28', '2021-07-16 15:11:38', '2021-08-20 14:14:57', null, '0', '26', 'cmsAd', 'cmsAd', '0', 'view/cms/cmsAd/cmsAd.vue', '0', '0', '0', '广告管理', 'circle-plus', '0');
INSERT INTO `sys_base_menus` VALUES ('29', '2021-07-19 14:16:27', '2021-08-20 14:15:22', null, '0', '26', 'cmsCat', 'cmsCat', '0', 'view/cms/cmsCat/cmsCat.vue', '0', '0', '0', '文章栏目', 's-tools', '0');
INSERT INTO `sys_base_menus` VALUES ('30', '2021-07-19 15:34:26', '2021-07-19 15:40:34', null, '0', '26', 'cmsArticle', 'cmsArticle', '0', 'view/cms/cmsArticle/cmsArticle.vue', '0', '0', '0', '文章管理', 'video-camera-solid', '0');
INSERT INTO `sys_base_menus` VALUES ('31', '2021-09-24 11:11:07', '2021-12-21 13:34:44', null, '0', '0', 'basic', 'basic', '0', 'view/basic/index.vue', '5', '0', '0', '基础数据', 's-help', '0');
INSERT INTO `sys_base_menus` VALUES ('32', '2021-09-24 11:14:47', '2021-09-24 11:14:47', null, '0', '31', 'basicFile', 'basicFile', '0', 'view/basic/basicFile/basicFile.vue', '0', '0', '0', '文件管理', 'picture-outline', '0');
INSERT INTO `sys_base_menus` VALUES ('33', '2021-10-09 11:55:50', '2021-10-14 17:47:17', null, '0', '26', 'cmsCatForm/:id', 'cmsCatForm', '1', 'view/cms/cmsCat/cmsCatForm.vue', '0', '0', '0', '编辑cmsCat', 'info', '0');
INSERT INTO `sys_base_menus` VALUES ('34', '2021-10-09 11:55:50', '2021-12-20 14:48:17', null, '0', '26', 'cmsArticleForm/:id', 'cmsArticleForm', '1', 'view/cms/cmsArticle/cmsArticleForm.vue', '0', '0', '0', '编辑cmsArticle', 'info', '0');
INSERT INTO `sys_base_menus` VALUES ('35', '2021-11-15 15:17:06', '2022-01-04 15:59:00', null, '0', '0', 'col', 'col', '0', 'view/col/index.vue', '9', '0', '0', '数据采集', 'guide', '0');
INSERT INTO `sys_base_menus` VALUES ('36', '2021-11-15 15:20:15', '2021-11-15 15:20:15', null, '0', '35', 'collect', 'collect', '0', 'view/col/colCollect/colCollect.vue', '0', '0', '0', '采集列表', 's-unfold', '0');
INSERT INTO `sys_base_menus` VALUES ('37', '2021-11-15 15:22:09', '2021-11-15 15:22:09', null, '0', '35', 'colKeyField', 'colKeyField', '0', 'view/col/colKeyField/colKeyField.vue', '0', '0', '0', '采集字段', 'video-camera', '0');
INSERT INTO `sys_base_menus` VALUES ('38', '2021-11-15 17:12:34', '2021-11-15 17:12:34', null, '0', '35', 'colHsj', 'colHsj', '0', 'view/col/colHsj/colHsj.vue', '0', '0', '0', '海事信息', 's-order', '0');
INSERT INTO `sys_base_menus` VALUES ('39', '2021-12-11 13:32:58', '2022-01-04 15:58:50', null, '0', '0', 'k8s', 'k8s', '0', 'view/k8s/index.vue', '10', '0', '0', 'K8S管理', 'cloudy', '0');
INSERT INTO `sys_base_menus` VALUES ('40', '2021-12-11 13:32:58', '2021-12-21 13:53:52', null, '0', '39', 'k8sClusters', 'k8sClusters', '0', 'view/k8s/k8sClusters/k8sClusters.vue', '0', '0', '0', '集群管理', 's-grid', '0');
INSERT INTO `sys_base_menus` VALUES ('41', '2021-12-11 13:32:58', '2021-12-21 13:54:31', null, '0', '39', 'k8sNamespaces', 'k8sNamespaces', '0', 'view/k8s/k8sNamespaces/k8sNamespaces.vue', '2', '0', '0', '命名空间管理', 'menu', '0');
INSERT INTO `sys_base_menus` VALUES ('42', '2021-12-11 13:32:58', '2021-12-21 13:54:45', null, '0', '39', 'k8sDeployments', 'k8sDeployments', '0', 'view/k8s/k8sDeployments/k8sDeployments.vue', '3', '0', '0', '应用管理', 's-grid', '0');
INSERT INTO `sys_base_menus` VALUES ('43', '2021-12-11 13:32:58', '2021-12-21 13:55:02', null, '0', '39', 'k8sPods', 'k8sPods', '0', 'view/k8s/k8sPods/k8sPods.vue', '4', '0', '0', '容器管理', 's-grid', '0');
INSERT INTO `sys_base_menus` VALUES ('44', '2021-12-11 13:32:58', '2021-12-21 13:54:17', null, '0', '39', 'k8sNodes', 'k8sNodes', '0', 'view/k8s/k8sNodes/k8sNodes.vue', '1', '0', '0', '节点管理', 's-grid', '0');
INSERT INTO `sys_base_menus` VALUES ('45', '2021-12-21 10:31:37', '2021-12-21 11:00:35', null, '0', '9', 'imTxim', 'imTxim', '0', 'view/im/imTxim/imTxim.vue', '0', '0', '0', 'IM账户管理', 's-order', '0');
INSERT INTO `sys_base_menus` VALUES ('46', '2021-12-22 14:09:02', '2021-12-22 14:20:16', null, '0', '0', 'mem', 'mem', '0', 'view/mem/index.vue', '4', '0', '0', '用户管理', 'user-solid', '0');
INSERT INTO `sys_base_menus` VALUES ('47', '2021-12-22 14:11:33', '2021-12-22 14:11:33', null, '0', '46', 'memUser', 'memUser', '0', 'view/mem/memUser/memUser.vue', '0', '0', '0', '用户列表', 'user', '0');
INSERT INTO `sys_base_menus` VALUES ('48', '2021-12-22 14:12:16', '2021-12-22 14:12:16', null, '0', '46', 'memUserForm', 'memUserForm', '1', 'view/mem/memUser/memUserForm.vue', '0', '0', '0', 'memUserForm', '', '0');
INSERT INTO `sys_base_menus` VALUES ('49', '2021-12-22 14:13:57', '2022-01-04 16:00:15', null, '0', '46', 'memUserSafe', 'memUserSafe', '0', 'view/mem/memUserSafe/memUserSafe.vue', '2', '0', '0', '用户安全', 'question', '0');
INSERT INTO `sys_base_menus` VALUES ('50', '2021-12-27 10:48:34', '2021-12-27 17:39:17', null, '0', '9', 'imTxFile', 'imTxFile', '0', 'view/im/imTxFile/imTxFile.vue', '7', '0', '0', '聊天文件', 'film', '0');
INSERT INTO `sys_base_menus` VALUES ('51', '2022-01-04 14:15:42', '2022-01-04 16:00:03', null, '0', '46', 'memAddress', 'memAddress', '0', 'view/mem/memAddress/memAddress.vue', '9', '0', '0', '会员地址', 's-grid', '0');
INSERT INTO `sys_base_menus` VALUES ('52', '2022-01-04 14:16:49', '2022-01-04 15:59:57', null, '0', '46', 'memUserLog', 'memUserLog', '0', 'view/mem/memUserLog/memUserLog.vue', '9', '0', '0', '用户日志', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('53', '2022-01-04 14:44:22', '2022-01-04 14:44:22', null, '0', '31', 'basicArea', 'basicArea', '0', 'view/basic/basicArea/basicArea.vue', '0', '0', '0', '行政地区', 'rank', '0');
INSERT INTO `sys_base_menus` VALUES ('54', '2022-01-13 15:40:02', '2022-01-13 16:38:03', null, '0', '0', 'bbpi', 'bbpi', '0', 'view/bbpi/index.vue', '3', '0', '0', '数据上报', 'ice-drink', '0');
INSERT INTO `sys_base_menus` VALUES ('55', null, '2022-01-13 15:56:26', null, '0', '54', 'bbPiContacts', 'bbPiContacts', '0', 'view/bbpi/bbPiContacts/bbPiContacts.vue', '9', '0', '0', '行政通讯录', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('56', null, '2022-01-13 15:57:31', null, '0', '54', 'bbPiContactsForm', 'bbPiContactsForm', '1', 'view/bbpi/bbPiContacts/bbPiContactsForm.vue', '9', '0', '0', '行政通讯录编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('57', null, null, null, '0', '54', 'bbPiDepartment', 'bbPiDepartment', '0', 'view/bbpi/bbPiDepartment/bbPiDepartment.vue', '9', '0', '0', '科室字典', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('58', null, null, null, '0', '54', 'bbPiDepartmentForm', 'bbPiDepartmentForm', '1', 'view/bbpi/bbPiDepartment/bbPiDepartmentForm.vue', '9', '0', '0', '科室字典编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('59', null, null, null, '0', '54', 'bbPiDevice', 'bbPiDevice', '0', 'view/bbpi/bbPiDevice/bbPiDevice.vue', '9', '0', '0', '医疗设备', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('60', null, null, null, '0', '54', 'bbPiDeviceForm', 'bbPiDeviceForm', '1', 'view/bbpi/bbPiDevice/bbPiDeviceForm.vue', '9', '0', '0', '医疗设备编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('61', null, null, null, '0', '54', 'bbPiInstitution', 'bbPiInstitution', '0', 'view/bbpi/bbPiInstitution/bbPiInstitution.vue', '9', '0', '0', '机构信息', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('62', null, null, null, '0', '54', 'bbPiInstitutionForm', 'bbPiInstitutionForm', '1', 'view/bbpi/bbPiInstitution/bbPiInstitutionForm.vue', '9', '0', '0', '机构信息编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('63', null, null, null, '0', '54', 'bbPiInstitutionBusiness', 'bbPiInstitutionBusiness', '0', 'view/bbpi/bbPiInstitutionBusiness/bbPiInstitutionBusiness.vue', '9', '0', '0', '机构业务', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('64', null, null, null, '0', '54', 'bbPiInstitutionBusinessForm', 'bbPiInstitutionBusinessForm', '1', 'view/bbpi/bbPiInstitutionBusiness/bbPiInstitutionBusinessForm.vue', '9', '0', '0', '机构业务编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('65', null, null, null, '0', '54', 'bbPiPerson', 'bbPiPerson', '0', 'view/bbpi/bbPiPerson/bbPiPerson.vue', '9', '0', '0', '患者信息', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('66', null, null, null, '0', '54', 'bbPiPersonForm', 'bbPiPersonForm', '1', 'view/bbpi/bbPiPerson/bbPiPersonForm.vue', '9', '0', '0', '患者信息编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('67', null, null, null, '0', '54', 'bbPiServicePoint', 'bbPiServicePoint', '0', 'view/bbpi/bbPiServicePoint/bbPiServicePoint.vue', '9', '0', '0', '服务网点', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('68', null, null, null, '0', '54', 'bbPiServicePointff', 'bbPiServicePointForm', '1', 'view/bbpi/bbPiServicePoint/bbPiServicePointForm.vue', '9', '0', '0', '服务网点编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('69', null, null, null, '0', '54', 'bbPiStaff', 'bbPiStaff', '0', 'view/bbpi/bbPiStaff/bbPiStaff.vue', '9', '0', '0', '人力资源', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('70', null, null, null, '0', '54', 'bbPiStaffForm', 'bbPiStaffForm', '1', 'view/bbpi/bbPiStaff/bbPiStaffForm.vue', '9', '0', '0', '人力资源编辑', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('71', null, null, null, '0', '54', 'bbPiTreatmentDiagnose', 'bbPiTreatmentDiagnose', '0', 'view/bbpi/bbPiTreatmentDiagnose/bbPiTreatmentDiagnose.vue', '100', '0', '0', '医学诊断', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('72', null, null, null, '0', '54', 'bbPiTreatmentDiagnoseForm', 'bbPiTreatmentDiagnoseForm', '1', 'view/bbpi/bbPiTreatmentDiagnose/bbPiTreatmentDiagnoseForm.vue', '100', '0', '0', '医学诊断Form', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('73', null, null, null, '0', '54', 'bbPiTreatmentOrder', 'bbPiTreatmentOrder', '0', 'view/bbpi/bbPiTreatmentOrder/bbPiTreatmentOrder.vue', '100', '0', '0', '诊疗处方', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('74', null, null, null, '0', '54', 'bbPiTreatmentOrderForm', 'bbPiTreatmentOrderForm', '1', 'view/bbpi/bbPiTreatmentOrder/bbPiTreatmentOrderForm.vue', '100', '0', '0', '诊疗处方Form', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('75', null, null, null, '0', '54', 'bbPiTreatmentRecord', 'bbPiTreatmentRecord', '0', 'view/bbpi/bbPiTreatmentRecord/bbPiTreatmentRecord.vue', '100', '0', '0', '诊疗病历', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('76', null, null, null, '0', '54', 'bbPiTreatmentRecordForm', 'bbPiTreatmentRecordForm', '1', 'view/bbpi/bbPiTreatmentRecord/bbPiTreatmentRecordForm.vue', '100', '0', '0', '诊疗病历From', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('77', null, null, null, '0', '54', 'bbPiTreatmentrReferral', 'bbPiTreatmentReferral', '0', 'view/bbpi/bbPiTreatmentReferral/bbPiTreatmentReferral.vue', '100', '0', '0', '转诊记录', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('78', null, null, null, '0', '54', 'bbPiTreatmentrReferralForm', 'bbPiTreatmentReferralForm', '1', 'view/bbpi/bbPiTreatmentReferral/bbPiTreatmentReferralForm.vue', '100', '0', '0', '转诊记录Form', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('79', null, null, null, '0', '54', 'bbPiUpField', 'bbPiUpField', '0', 'view/bbpi/bbPiUpField/bbPiUpField.vue', '110', '0', '0', '上传字段设置', 's-fold', '0');
INSERT INTO `sys_base_menus` VALUES ('80', null, null, null, '0', '54', 'bbPiUpFieldForm', 'bbPiUpFieldForm', '1', 'view/bbpi/bbPiUpField/bbPiUpFieldForm.vue', '110', '0', '0', '上传字段设置Form', 's-fold', '0');

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES ('888', '888');
INSERT INTO `sys_data_authority_id` VALUES ('888', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('888', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '9528');

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES ('1', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, '性别', 'sex', '1', '性别字典');
INSERT INTO `sys_dictionaries` VALUES ('2', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, '数据库int类型', 'int', '1', 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES ('3', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, '数据库时间日期类型', 'time.Time', '1', '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES ('4', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, '数据库浮点型', 'float64', '1', '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES ('5', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, '数据库字符串', 'string', '1', '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES ('6', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, '数据库bool类型', 'bool', '1', '数据库bool类型');
INSERT INTO `sys_dictionaries` VALUES ('7', '2021-07-07 16:03:04', '2021-07-07 16:03:04', null, '平台用户', 'plat', '1', '平台用户类别');
INSERT INTO `sys_dictionaries` VALUES ('8', '2021-07-07 16:07:47', '2021-07-07 16:07:47', null, '记录状态', 'status', '1', '记录状态');
INSERT INTO `sys_dictionaries` VALUES ('9', '2021-07-07 16:27:25', '2021-07-07 16:27:25', null, '是或否', 'truefalse', '1', '是或否');
INSERT INTO `sys_dictionaries` VALUES ('10', '2021-07-07 16:46:24', '2021-07-07 17:25:02', null, '文章媒体类型', 'media_type', '1', '0文章, 1图片 2视频, 3音频，4链接 5小游戏');
INSERT INTO `sys_dictionaries` VALUES ('11', '2021-07-07 17:04:38', '2021-07-07 17:04:38', null, '兴趣圈子分类', 'group_type', '1', '兴趣圈子分类');
INSERT INTO `sys_dictionaries` VALUES ('12', '2021-08-17 14:04:00', '2021-08-17 14:04:00', null, '数据库图片类型', 'image', '1', '数据库图片类型');
INSERT INTO `sys_dictionaries` VALUES ('13', '2021-09-24 10:05:36', '2021-09-24 10:05:36', null, '系统模块', 'module', '1', '系统模块名称');
INSERT INTO `sys_dictionaries` VALUES ('14', '2021-09-24 16:01:45', '2021-09-24 16:03:19', null, '文件存储', 'driver', '1', '文件存储设备');
INSERT INTO `sys_dictionaries` VALUES ('15', '2021-11-22 14:58:52', '2021-11-22 14:58:52', null, '运行状态', 'status_run', '1', '运行状态');
INSERT INTO `sys_dictionaries` VALUES ('16', '2021-12-20 14:13:06', '2021-12-20 14:13:06', '2021-12-20 14:14:13', 'IM消息类型', 'im_chat_type', '1', 'IM消息类型');
INSERT INTO `sys_dictionaries` VALUES ('17', '2021-12-20 14:31:03', '2021-12-20 14:31:03', null, '下载状态', 'status_download', '1', '下载状态');
INSERT INTO `sys_dictionaries` VALUES ('18', '2021-12-20 14:32:24', '2021-12-20 14:32:24', null, '上传状态', 'status_up', '1', '上传状态');
INSERT INTO `sys_dictionaries` VALUES ('19', '2022-01-04 15:38:06', '2022-01-04 15:38:06', null, '用户安全操作类型', 'usersafe_type', '1', '用户安全-操作类型');
INSERT INTO `sys_dictionaries` VALUES ('20', '2022-01-04 15:40:29', '2022-01-04 15:40:29', null, '用户日志类型', 'userlog_type', '1', '用户日志类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=83 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES ('1', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'smallint', '1', '1', '1', '2');
INSERT INTO `sys_dictionary_details` VALUES ('2', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'mediumint', '2', '1', '2', '2');
INSERT INTO `sys_dictionary_details` VALUES ('3', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'int', '3', '1', '3', '2');
INSERT INTO `sys_dictionary_details` VALUES ('4', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'bigint', '4', '1', '4', '2');
INSERT INTO `sys_dictionary_details` VALUES ('5', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'date', '0', '1', '0', '3');
INSERT INTO `sys_dictionary_details` VALUES ('6', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'time', '1', '1', '1', '3');
INSERT INTO `sys_dictionary_details` VALUES ('7', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'year', '2', '1', '2', '3');
INSERT INTO `sys_dictionary_details` VALUES ('8', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'datetime', '3', '1', '3', '3');
INSERT INTO `sys_dictionary_details` VALUES ('9', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'timestamp', '5', '1', '5', '3');
INSERT INTO `sys_dictionary_details` VALUES ('10', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'float', '0', '1', '0', '4');
INSERT INTO `sys_dictionary_details` VALUES ('11', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'double', '1', '1', '1', '4');
INSERT INTO `sys_dictionary_details` VALUES ('12', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'decimal', '2', '1', '2', '4');
INSERT INTO `sys_dictionary_details` VALUES ('13', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'char', '0', '1', '0', '5');
INSERT INTO `sys_dictionary_details` VALUES ('14', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'varchar', '1', '1', '1', '5');
INSERT INTO `sys_dictionary_details` VALUES ('15', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'tinyblob', '2', '1', '2', '5');
INSERT INTO `sys_dictionary_details` VALUES ('16', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'tinytext', '3', '1', '3', '5');
INSERT INTO `sys_dictionary_details` VALUES ('17', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'text', '4', '1', '4', '5');
INSERT INTO `sys_dictionary_details` VALUES ('18', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'blob', '5', '1', '5', '5');
INSERT INTO `sys_dictionary_details` VALUES ('19', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'mediumblob', '6', '1', '6', '5');
INSERT INTO `sys_dictionary_details` VALUES ('20', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'mediumtext', '7', '1', '7', '5');
INSERT INTO `sys_dictionary_details` VALUES ('21', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'longblob', '8', '1', '8', '5');
INSERT INTO `sys_dictionary_details` VALUES ('22', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'longtext', '9', '1', '9', '5');
INSERT INTO `sys_dictionary_details` VALUES ('23', '2021-09-10 14:11:33', '2021-09-10 14:11:33', null, 'tinyint', '0', '1', '0', '6');
INSERT INTO `sys_dictionary_details` VALUES ('24', '2021-07-07 16:00:37', '2021-10-27 17:04:18', null, '男', '1', '1', '1', '1');
INSERT INTO `sys_dictionary_details` VALUES ('25', '2021-07-07 16:00:51', '2021-07-07 16:00:51', null, '女', '2', '1', '1', '1');
INSERT INTO `sys_dictionary_details` VALUES ('26', '2021-07-07 16:01:12', '2021-07-07 16:01:12', null, '未知', '0', '1', '3', '1');
INSERT INTO `sys_dictionary_details` VALUES ('27', '2021-07-07 16:08:04', '2021-07-07 16:08:04', null, '未审核', '0', '1', '0', '8');
INSERT INTO `sys_dictionary_details` VALUES ('28', '2021-07-07 16:08:33', '2021-07-07 16:08:33', null, '已审核(有效)', '1', '1', '1', '8');
INSERT INTO `sys_dictionary_details` VALUES ('29', '2021-07-07 16:09:03', '2021-07-07 16:09:03', null, '审核未通过', '2', '1', '2', '8');
INSERT INTO `sys_dictionary_details` VALUES ('30', '2021-07-07 16:27:43', '2021-07-07 16:27:43', null, '是', '1', '1', '1', '9');
INSERT INTO `sys_dictionary_details` VALUES ('31', '2021-07-07 16:27:54', '2021-07-07 16:27:54', null, '否', '0', '1', '2', '9');
INSERT INTO `sys_dictionary_details` VALUES ('32', '2021-07-07 16:44:04', '2021-07-07 16:44:04', null, '存草稿', '3', '1', '3', '8');
INSERT INTO `sys_dictionary_details` VALUES ('33', '2021-07-07 16:46:41', '2021-09-24 10:29:20', null, '文章', '1', '1', '1', '10');
INSERT INTO `sys_dictionary_details` VALUES ('34', '2021-07-07 16:46:51', '2021-09-24 10:29:24', null, '图片', '2', '1', '2', '10');
INSERT INTO `sys_dictionary_details` VALUES ('35', '2021-07-07 16:47:06', '2021-07-07 16:47:06', null, '视频', '3', '1', '3', '10');
INSERT INTO `sys_dictionary_details` VALUES ('36', '2021-07-07 16:47:21', '2021-07-07 16:47:21', null, '音频', '4', '1', '4', '10');
INSERT INTO `sys_dictionary_details` VALUES ('37', '2021-07-07 16:47:31', '2021-07-07 16:47:31', null, '链接', '5', '1', '5', '10');
INSERT INTO `sys_dictionary_details` VALUES ('38', '2021-07-07 17:08:04', '2021-07-07 17:11:09', null, '民族音乐', '1', '1', '1', '11');
INSERT INTO `sys_dictionary_details` VALUES ('39', '2021-07-07 17:08:40', '2021-07-07 17:08:40', null, '书法美术', '2', '1', '2', '11');
INSERT INTO `sys_dictionary_details` VALUES ('40', '2021-07-07 17:09:27', '2021-07-07 17:09:27', null, '古文诗词', '3', '1', '3', '11');
INSERT INTO `sys_dictionary_details` VALUES ('41', '2021-07-07 17:10:29', '2021-07-07 17:10:29', null, '曲艺戏剧', '4', '1', '4', '11');
INSERT INTO `sys_dictionary_details` VALUES ('42', '2021-07-07 17:11:48', '2021-07-07 17:11:48', null, '园林建筑', '5', '1', '5', '11');
INSERT INTO `sys_dictionary_details` VALUES ('43', '2021-07-07 17:12:22', '2021-07-07 17:12:22', null, '对联灯谜酒令', '5', '1', '5', '11');
INSERT INTO `sys_dictionary_details` VALUES ('44', '2021-07-07 17:13:37', '2021-07-07 17:13:37', null, '手工艺品', '6', '1', '6', '11');
INSERT INTO `sys_dictionary_details` VALUES ('45', '2021-07-07 17:14:11', '2021-07-07 17:14:11', null, '摄影旅行', '1', '1', '1', '11');
INSERT INTO `sys_dictionary_details` VALUES ('46', '2021-07-07 17:26:18', '2021-07-07 17:26:18', null, '小游戏', '6', '1', '6', '10');
INSERT INTO `sys_dictionary_details` VALUES ('47', '2021-08-17 14:07:49', '2021-08-17 14:07:58', null, '单图片', '1', '1', '0', '12');
INSERT INTO `sys_dictionary_details` VALUES ('48', '2021-08-17 14:08:19', '2021-08-17 14:14:52', null, '多图片', '2', '1', '0', '12');
INSERT INTO `sys_dictionary_details` VALUES ('49', '2021-09-24 10:06:10', '2021-09-24 10:06:10', null, 'admin超管', '1', '1', '1', '13');
INSERT INTO `sys_dictionary_details` VALUES ('50', '2021-09-24 10:06:37', '2021-09-24 10:06:37', null, 'basic基础', '2', '1', '2', '13');
INSERT INTO `sys_dictionary_details` VALUES ('51', '2021-09-24 10:07:03', '2021-09-24 10:07:15', null, 'system系统', '3', '1', '3', '13');
INSERT INTO `sys_dictionary_details` VALUES ('52', '2021-09-24 10:07:53', '2021-09-24 10:07:53', null, 'common通用', '4', '1', '4', '13');
INSERT INTO `sys_dictionary_details` VALUES ('53', '2021-09-24 10:08:08', '2021-09-24 10:08:08', null, 'cms新闻内容', '5', '1', '5', '13');
INSERT INTO `sys_dictionary_details` VALUES ('54', '2021-09-24 10:08:08', '2021-09-24 10:08:08', null, 'game游戏圈', '6', '1', '6', '13');
INSERT INTO `sys_dictionary_details` VALUES ('55', '2021-09-24 10:08:08', '2021-09-24 10:08:08', null, 'act活动圈', '7', '1', '6', '13');
INSERT INTO `sys_dictionary_details` VALUES ('56', '2021-09-24 10:08:08', '2021-09-24 10:08:08', null, 'it技术圈', '8', '1', '8', '13');
INSERT INTO `sys_dictionary_details` VALUES ('57', '2021-09-24 10:08:08', '2021-09-24 10:08:08', null, 'award抽奖圈', '9', '1', '9', '13');
INSERT INTO `sys_dictionary_details` VALUES ('58', '2021-09-24 10:29:10', '2021-09-27 17:18:41', null, '默认未知', '0', '1', '0', '10');
INSERT INTO `sys_dictionary_details` VALUES ('59', '2021-09-24 10:32:20', '2021-09-24 10:32:20', null, '直播', '7', '1', '7', '10');
INSERT INTO `sys_dictionary_details` VALUES ('60', '2021-09-24 16:02:07', '2021-09-24 16:02:07', null, '本地', '0', '1', '0', '14');
INSERT INTO `sys_dictionary_details` VALUES ('61', '2021-09-24 16:02:17', '2021-09-24 16:02:17', null, '阿里云', '1', '1', '1', '14');
INSERT INTO `sys_dictionary_details` VALUES ('62', '2021-09-24 16:02:31', '2021-09-24 16:02:31', null, '腾讯云', '2', '1', '2', '14');
INSERT INTO `sys_dictionary_details` VALUES ('63', '2021-09-24 16:02:47', '2021-09-24 16:02:47', null, '七牛', '3', '1', '3', '14');
INSERT INTO `sys_dictionary_details` VALUES ('64', '2021-09-27 17:18:15', '2021-09-27 17:18:15', null, 'default默认', '0', '1', '0', '13');
INSERT INTO `sys_dictionary_details` VALUES ('65', '2021-10-27 23:30:00', '2021-10-27 23:30:00', '2021-10-27 23:30:17', '3333', '3333', '1', '3333', '10');
INSERT INTO `sys_dictionary_details` VALUES ('66', '2021-11-22 14:59:20', '2022-03-29 15:42:05', null, '开始(运行中)', '1', '1', '1', '15');
INSERT INTO `sys_dictionary_details` VALUES ('67', '2021-11-22 14:59:58', '2021-11-22 14:59:58', null, '停止', '0', '1', '0', '15');
INSERT INTO `sys_dictionary_details` VALUES ('68', '2021-11-22 15:00:16', '2022-03-29 15:41:19', null, '结束', '2', '1', '2', '15');
INSERT INTO `sys_dictionary_details` VALUES ('69', '2021-12-20 14:13:23', '2021-12-20 14:13:23', '2021-12-20 14:14:05', 'C2C', '1', '1', '1', '16');
INSERT INTO `sys_dictionary_details` VALUES ('70', '2021-12-20 14:31:26', '2021-12-20 14:31:26', null, '未下载', '0', '1', '0', '17');
INSERT INTO `sys_dictionary_details` VALUES ('71', '2021-12-20 14:31:41', '2021-12-20 14:31:41', null, '已下载', '1', '1', '1', '17');
INSERT INTO `sys_dictionary_details` VALUES ('72', '2021-12-20 14:31:51', '2021-12-20 14:31:51', null, '下载失败', '2', '1', '2', '17');
INSERT INTO `sys_dictionary_details` VALUES ('73', '2021-12-20 14:32:49', '2021-12-20 14:32:49', null, '未上传', '0', '1', '0', '18');
INSERT INTO `sys_dictionary_details` VALUES ('74', '2021-12-20 14:33:03', '2021-12-20 14:33:03', null, '已上传', '1', '1', '1', '18');
INSERT INTO `sys_dictionary_details` VALUES ('75', '2021-12-20 14:33:14', '2021-12-20 14:33:14', null, '上传失败', '2', '1', '2', '18');
INSERT INTO `sys_dictionary_details` VALUES ('76', '2022-01-04 15:38:20', '2022-01-04 15:38:20', null, '改密码', '1', '1', '1', '19');
INSERT INTO `sys_dictionary_details` VALUES ('77', '2022-01-04 15:38:38', '2022-01-04 15:38:38', null, '改手机号', '2', '1', '2', '19');
INSERT INTO `sys_dictionary_details` VALUES ('78', '2022-01-04 15:38:53', '2022-01-04 15:38:53', null, '改用户名', '3', '1', '3', '19');
INSERT INTO `sys_dictionary_details` VALUES ('79', '2022-01-04 15:39:05', '2022-01-04 15:39:05', null, '实名认证', '4', '1', '4', '19');
INSERT INTO `sys_dictionary_details` VALUES ('80', '2022-01-04 15:39:17', '2022-01-04 15:39:17', null, '修改支付密码', '5', '1', '5', '19');
INSERT INTO `sys_dictionary_details` VALUES ('81', '2022-01-04 15:41:04', '2022-01-04 15:41:04', null, '用户登录', '1', '1', '1', '20');
INSERT INTO `sys_dictionary_details` VALUES ('82', '2022-01-04 15:41:20', '2022-01-04 15:41:20', null, '用户退出', '2', '1', '2', '20');

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4849 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
 

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES ('1', '888');
INSERT INTO `sys_user_authority` VALUES ('1', '8881');
INSERT INTO `sys_user_authority` VALUES ('1', '9528');
INSERT INTO `sys_user_authority` VALUES ('2', '888');
INSERT INTO `sys_user_authority` VALUES ('2', '8881');
INSERT INTO `sys_user_authority` VALUES ('3', '888');
INSERT INTO `sys_user_authority` VALUES ('3', '8881');
INSERT INTO `sys_user_authority` VALUES ('4', '888');
INSERT INTO `sys_user_authority` VALUES ('4', '8881');
INSERT INTO `sys_user_authority` VALUES ('4', '9528');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '888' COMMENT '用户角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_users
-- ---------------------------- 
 
INSERT INTO `sys_users` VALUES ('5', '2021-12-09 14:19:33', '2021-12-09 14:19:33', null, '9e0e3942-7262-453f-832e-9d7a81bb9525', 'test123', 'cc03e747a6afbbcbf8be7668acfebee5', 'test123', 'dark', 'https://cms.88act.com/res/sys/0/avatar.jpg', '#fff', '#1890ff', '888');


CREATE  VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`close_tab` AS `close_tab`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));