/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-20 14:48:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_user_topstick
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_topstick`;
CREATE TABLE `tb_user_topstick` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `account_num` int(10) NOT NULL,
  `dialog_id` int(10) NOT NULL COMMENT '帖子id',
  `dialog_type` tinyint(1) NOT NULL,
  `abs_path` char(15) COLLATE utf8_unicode_ci NOT NULL,
  `top_value` int(10) NOT NULL DEFAULT '0' COMMENT '顶贴值',
  `down_value` int(10) NOT NULL DEFAULT '0',
  `create_time` int(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_num_dialog_id_abs_path` (`account_num`,`dialog_id`,`abs_path`,`dialog_type`) USING BTREE COMMENT '针对对白id，位置信息建立唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=179 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
