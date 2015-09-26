/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-13 15:49:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_karma_msg
-- ----------------------------
DROP TABLE IF EXISTS `tb_karma_msg`;
CREATE TABLE `tb_karma_msg` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `send_id` int(8) DEFAULT '0',
  `receive_id` int(8) NOT NULL DEFAULT '0',
  `content` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `type` tinyint(1) DEFAULT '0',
  `param` tinyint(1) DEFAULT '0',
  `img` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0' COMMENT '0： 未读，1：已读：2：逻辑删除',
  `create_time` int(8) DEFAULT '0',
  PRIMARY KEY (`id`,`receive_id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
