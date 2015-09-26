/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-13 15:49:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_karma_relation
-- ----------------------------
DROP TABLE IF EXISTS `tb_karma_relation`;
CREATE TABLE `tb_karma_relation` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `account_num` int(8) NOT NULL,
  `friend_id` int(8) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `create_time` int(8) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_friend` (`account_num`,`friend_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
