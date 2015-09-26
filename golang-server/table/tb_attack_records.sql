/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-20 15:13:16
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_attack_records
-- ----------------------------
DROP TABLE IF EXISTS `tb_attack_records`;
CREATE TABLE `tb_attack_records` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '记录id',
  `attack_user` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '进攻者ID',
  `defensed_user` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '防御者ID',
  `attack_prop` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '进攻道具ID',
  `defensed_prop` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '防御道具 0:没有防御道具',
  `change_number` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '改变数值',
  `change_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '改变数值类型 0:积分',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '进攻状态 0:成功 1:被防御 2:反弹',
  `msg_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '信息状态 0:未读，1:已读',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=394 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
