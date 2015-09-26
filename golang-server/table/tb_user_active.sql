/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-09-18 17:09:03
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_user_active
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_active`;
CREATE TABLE `tb_user_active` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `account_num` int(10) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `last_time` int(10) NOT NULL DEFAULT '0' COMMENT '上一次请求时间',
  `times` tinyint(1) NOT NULL DEFAULT '0' COMMENT '当天获得的活跃天数次数',
  `is_online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否在线',
  `login_ip` int(10) NOT NULL DEFAULT '0' COMMENT '登陆IP',
  `day_num` int(10) NOT NULL DEFAULT '0' COMMENT '连续登陆天数',
  `active_num` char(10) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0' COMMENT '累计活跃天数',
  `login_time` int(10) NOT NULL DEFAULT '0',
  `logout_time` int(10) NOT NULL DEFAULT '0',
  `online_time` int(10) NOT NULL DEFAULT '0' COMMENT '当天累计在线时长(单位秒)',
  PRIMARY KEY (`id`,`account_num`),
  UNIQUE KEY `account_num` (`account_num`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
