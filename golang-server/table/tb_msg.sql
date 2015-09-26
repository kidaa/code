/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-09-21 21:07:41
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_msg
-- ----------------------------
DROP TABLE IF EXISTS `tb_msg`;
CREATE TABLE `tb_msg` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `send_id` int(10) NOT NULL DEFAULT '0' COMMENT '发送者id，默认0：系统管理员',
  `receive_id` int(10) NOT NULL DEFAULT '0' COMMENT '接受者id',
  `content` varchar(600) NOT NULL DEFAULT '' COMMENT '系统消息',
  `third_id` varchar(20) NOT NULL DEFAULT '' COMMENT '第三方id，如：道具编号',
  `img` varchar(500) DEFAULT '' COMMENT '图片全路径',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '系统消息类型; 默认1:注册通知, 2:邮箱绑定成功, 3:手机绑定成功, 4:站内违规发布内容已被处理, 5:第三方账号绑定成功, 6:顶对白',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否屏蔽; 默认0未读，1已读，2逻辑删除',
  `receive_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认值0:未读, 1:已读, 2删除(逻辑删除)',
  `create_time` int(10) NOT NULL DEFAULT '0' COMMENT '时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=521 DEFAULT CHARSET=utf8 COMMENT='消息';
