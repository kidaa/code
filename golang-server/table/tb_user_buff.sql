/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-20 15:13:46
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_user_buff
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_buff`;
CREATE TABLE `tb_user_buff` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `prop_number` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '道具id',
  `prop_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '道具类型',
  `start_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '结束时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_num_prop_type` (`account_num`,`prop_type`) USING BTREE COMMENT '用户id 道具类型',
  UNIQUE KEY `account_num_prop_number` (`account_num`,`prop_number`) USING BTREE COMMENT '用户id 道具ID'
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户有益状态表';
