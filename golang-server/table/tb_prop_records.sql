/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-20 15:13:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_prop_records
-- ----------------------------
DROP TABLE IF EXISTS `tb_prop_records`;
CREATE TABLE `tb_prop_records` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '道具记录ID',
  `account_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `prop_number` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '道具ID',
  `count_num` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '数量',
  `record_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '道具记录类型 0:获得，1:消耗',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '记录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=788 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='道具记录表';
