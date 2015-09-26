/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-13 15:49:07
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_karma
-- ----------------------------
DROP TABLE IF EXISTS `tb_karma`;
CREATE TABLE `tb_karma` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `send_id` int(11) NOT NULL DEFAULT '0' COMMENT '发纸条的用户唯一识别码',
  `widget_id` int(11) DEFAULT '0' COMMENT '纸条道具唯一识别码',
  `content` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '纸条携带的消息内容',
  `read_times` tinyint(4) DEFAULT '0' COMMENT '缘分纸条被其他用户获取的次数',
  `sex` tinyint(1) DEFAULT NULL,
  `birth` int(8) DEFAULT NULL,
  `city_id` tinyint(4) DEFAULT NULL,
  `province_id` tinyint(4) DEFAULT NULL,
  `create_time` int(10) DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`,`send_id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
