/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.82(主)
Source Server Version : 50619
Source Host           : 192.168.1.82:3306
Source Database       : vragon_debug

Target Server Type    : MYSQL
Target Server Version : 50619
File Encoding         : 65001

Date: 2015-08-20 14:48:50
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_sys_dialog_stat
-- ----------------------------
DROP TABLE IF EXISTS `tb_sys_dialog_stat`;
CREATE TABLE `tb_sys_dialog_stat` (
  `stat_id` int(10) NOT NULL,
  `dialog_id` int(10) NOT NULL DEFAULT '0',
  `abs_path` char(15) COLLATE utf8_unicode_ci NOT NULL,
  `up_worth` mediumint(8) NOT NULL,
  `up_total_time` mediumint(8) NOT NULL,
  `up_end_time` int(10) NOT NULL,
  `up_remain_time` mediumint(5) NOT NULL,
  `up_flag` tinyint(1) NOT NULL,
  `follow_num` mediumint(8) NOT NULL,
  `hot_num` smallint(5) NOT NULL,
  `dialog_type` tinyint(1) NOT NULL,
  `create_time` int(10) NOT NULL,
  PRIMARY KEY (`stat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of tb_sys_dialog_stat
-- ----------------------------
INSERT INTO `tb_sys_dialog_stat` VALUES ('77', '2', '1_1_2_1_7', '172', '0', '0', '0', '0', '0', '0', '2', '1432636419');
INSERT INTO `tb_sys_dialog_stat` VALUES ('78', '3', '1_1_2_0_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432636454');
INSERT INTO `tb_sys_dialog_stat` VALUES ('79', '4', '1_1_2_2_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432636831');
INSERT INTO `tb_sys_dialog_stat` VALUES ('80', '5', '1_1_2_2_4', '80', '0', '0', '0', '0', '0', '0', '2', '1432636846');
INSERT INTO `tb_sys_dialog_stat` VALUES ('81', '6', '1_1_2_2_2', '112', '0', '0', '0', '0', '0', '0', '2', '1432636861');
INSERT INTO `tb_sys_dialog_stat` VALUES ('82', '7', '1_1_2_0_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432636874');
INSERT INTO `tb_sys_dialog_stat` VALUES ('83', '8', '1_1_2_0_6', '99', '0', '0', '0', '0', '0', '0', '2', '1432636888');
INSERT INTO `tb_sys_dialog_stat` VALUES ('84', '9', '1_2_1_8_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432636889');
INSERT INTO `tb_sys_dialog_stat` VALUES ('85', '10', '1_1_2_1_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432636918');
INSERT INTO `tb_sys_dialog_stat` VALUES ('86', '11', '1_2_1_6_4', '123', '0', '0', '0', '0', '0', '0', '2', '1432636918');
INSERT INTO `tb_sys_dialog_stat` VALUES ('87', '12', '1_1_2_1_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432636935');
INSERT INTO `tb_sys_dialog_stat` VALUES ('88', '13', '1_1_1_8_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432636949');
INSERT INTO `tb_sys_dialog_stat` VALUES ('89', '14', '1_2_4_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432636959');
INSERT INTO `tb_sys_dialog_stat` VALUES ('90', '15', '1_1_1_9_7', '101', '0', '0', '0', '0', '0', '0', '2', '1432636967');
INSERT INTO `tb_sys_dialog_stat` VALUES ('91', '16', '1_1_5_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432636983');
INSERT INTO `tb_sys_dialog_stat` VALUES ('92', '17', '1_2_1_7_7', '101', '0', '0', '0', '0', '0', '0', '2', '1432636987');
INSERT INTO `tb_sys_dialog_stat` VALUES ('93', '18', '1_3_1_5_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432636999');
INSERT INTO `tb_sys_dialog_stat` VALUES ('94', '19', '1_1_5_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432637001');
INSERT INTO `tb_sys_dialog_stat` VALUES ('95', '20', '1_1_4_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637015');
INSERT INTO `tb_sys_dialog_stat` VALUES ('96', '21', '1_1_5_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637029');
INSERT INTO `tb_sys_dialog_stat` VALUES ('97', '22', '1_2_1_7_8', '101', '0', '0', '0', '0', '0', '0', '2', '1432637037');
INSERT INTO `tb_sys_dialog_stat` VALUES ('98', '23', '1_1_1_0_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637043');
INSERT INTO `tb_sys_dialog_stat` VALUES ('99', '24', '1_1_1_0_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637060');
INSERT INTO `tb_sys_dialog_stat` VALUES ('100', '25', '1_2_1_6_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637069');
INSERT INTO `tb_sys_dialog_stat` VALUES ('101', '26', '1_1_1_3_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432637073');
INSERT INTO `tb_sys_dialog_stat` VALUES ('102', '27', '1_3_1_5_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637081');
INSERT INTO `tb_sys_dialog_stat` VALUES ('103', '28', '1_1_1_0_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637086');
INSERT INTO `tb_sys_dialog_stat` VALUES ('104', '29', '1_1_5_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637101');
INSERT INTO `tb_sys_dialog_stat` VALUES ('105', '30', '1_2_4_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637107');
INSERT INTO `tb_sys_dialog_stat` VALUES ('106', '31', '1_3_1_6_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637113');
INSERT INTO `tb_sys_dialog_stat` VALUES ('107', '32', '1_1_1_3_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637117');
INSERT INTO `tb_sys_dialog_stat` VALUES ('108', '33', '1_1_1_0_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637128');
INSERT INTO `tb_sys_dialog_stat` VALUES ('109', '34', '1_1_1_2_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637141');
INSERT INTO `tb_sys_dialog_stat` VALUES ('110', '35', '1_1_1_2_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432637153');
INSERT INTO `tb_sys_dialog_stat` VALUES ('111', '36', '1_2_8_9', '101', '0', '0', '0', '0', '0', '0', '2', '1432637165');
INSERT INTO `tb_sys_dialog_stat` VALUES ('112', '37', '1_1_4_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432637166');
INSERT INTO `tb_sys_dialog_stat` VALUES ('113', '38', '1_3_1_4_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637184');
INSERT INTO `tb_sys_dialog_stat` VALUES ('114', '39', '1_1_1_7_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637185');
INSERT INTO `tb_sys_dialog_stat` VALUES ('115', '40', '1_1_1_7_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637199');
INSERT INTO `tb_sys_dialog_stat` VALUES ('116', '41', '1_2_4_3', '31', '0', '0', '0', '0', '0', '0', '2', '1432637206');
INSERT INTO `tb_sys_dialog_stat` VALUES ('117', '42', '1_1_1_7_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637213');
INSERT INTO `tb_sys_dialog_stat` VALUES ('118', '43', '1_3_1_4_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637230');
INSERT INTO `tb_sys_dialog_stat` VALUES ('119', '44', '1_1_3_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432637231');
INSERT INTO `tb_sys_dialog_stat` VALUES ('120', '45', '1_2_4_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432637237');
INSERT INTO `tb_sys_dialog_stat` VALUES ('121', '46', '1_1_8_7', '99', '0', '0', '0', '0', '0', '0', '2', '1432637245');
INSERT INTO `tb_sys_dialog_stat` VALUES ('122', '47', '1_1_9_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637260');
INSERT INTO `tb_sys_dialog_stat` VALUES ('123', '48', '1_3_1_6_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432637267');
INSERT INTO `tb_sys_dialog_stat` VALUES ('124', '49', '1_1_8_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637271');
INSERT INTO `tb_sys_dialog_stat` VALUES ('125', '50', '1_2_4_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432637282');
INSERT INTO `tb_sys_dialog_stat` VALUES ('126', '51', '1_1_8_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637289');
INSERT INTO `tb_sys_dialog_stat` VALUES ('127', '52', '1_3_1_6_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637295');
INSERT INTO `tb_sys_dialog_stat` VALUES ('128', '53', '1_1_7_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432637306');
INSERT INTO `tb_sys_dialog_stat` VALUES ('129', '54', '1_1_7_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637322');
INSERT INTO `tb_sys_dialog_stat` VALUES ('130', '55', '1_2_1_5_4', '101', '0', '0', '0', '0', '0', '0', '2', '1432637326');
INSERT INTO `tb_sys_dialog_stat` VALUES ('131', '56', '1_3_8_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637332');
INSERT INTO `tb_sys_dialog_stat` VALUES ('132', '57', '1_1_7_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637338');
INSERT INTO `tb_sys_dialog_stat` VALUES ('133', '58', '1_2_1_5_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637354');
INSERT INTO `tb_sys_dialog_stat` VALUES ('134', '59', '1_1_7_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637357');
INSERT INTO `tb_sys_dialog_stat` VALUES ('135', '60', '1_3_8_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637363');
INSERT INTO `tb_sys_dialog_stat` VALUES ('136', '61', '1_1_8_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637374');
INSERT INTO `tb_sys_dialog_stat` VALUES ('137', '62', '1_2_2_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637384');
INSERT INTO `tb_sys_dialog_stat` VALUES ('138', '63', '1_1_8_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432637400');
INSERT INTO `tb_sys_dialog_stat` VALUES ('139', '64', '1_1_7_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637416');
INSERT INTO `tb_sys_dialog_stat` VALUES ('140', '65', '1_1_7_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432637432');
INSERT INTO `tb_sys_dialog_stat` VALUES ('141', '66', '1_3_5_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637435');
INSERT INTO `tb_sys_dialog_stat` VALUES ('142', '67', '1_1_7_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637445');
INSERT INTO `tb_sys_dialog_stat` VALUES ('143', '68', '1_1_6_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637458');
INSERT INTO `tb_sys_dialog_stat` VALUES ('144', '69', '1_1_7_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432637477');
INSERT INTO `tb_sys_dialog_stat` VALUES ('145', '70', '1_3_6_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432637486');
INSERT INTO `tb_sys_dialog_stat` VALUES ('146', '71', '1_1_7_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637490');
INSERT INTO `tb_sys_dialog_stat` VALUES ('147', '72', '1_1_6_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637506');
INSERT INTO `tb_sys_dialog_stat` VALUES ('148', '73', '1_2_1_2_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637512');
INSERT INTO `tb_sys_dialog_stat` VALUES ('149', '74', '1_1_6_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637518');
INSERT INTO `tb_sys_dialog_stat` VALUES ('150', '75', '1_1_6_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637536');
INSERT INTO `tb_sys_dialog_stat` VALUES ('151', '76', '1_3_6_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637542');
INSERT INTO `tb_sys_dialog_stat` VALUES ('152', '77', '1_2_1_2_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432637543');
INSERT INTO `tb_sys_dialog_stat` VALUES ('153', '78', '1_1_2_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637551');
INSERT INTO `tb_sys_dialog_stat` VALUES ('154', '79', '1_1_2_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637565');
INSERT INTO `tb_sys_dialog_stat` VALUES ('155', '80', '1_3_8_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637578');
INSERT INTO `tb_sys_dialog_stat` VALUES ('156', '81', '1_2_1_2_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637579');
INSERT INTO `tb_sys_dialog_stat` VALUES ('157', '82', '1_1_2_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432637581');
INSERT INTO `tb_sys_dialog_stat` VALUES ('158', '83', '1_1_3_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637595');
INSERT INTO `tb_sys_dialog_stat` VALUES ('159', '84', '1_3_9_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432637606');
INSERT INTO `tb_sys_dialog_stat` VALUES ('160', '85', '1_1_1_7_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432637610');
INSERT INTO `tb_sys_dialog_stat` VALUES ('161', '86', '1_2_1_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637626');
INSERT INTO `tb_sys_dialog_stat` VALUES ('162', '87', '1_1_1_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637627');
INSERT INTO `tb_sys_dialog_stat` VALUES ('163', '88', '1_1_1_4_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637641');
INSERT INTO `tb_sys_dialog_stat` VALUES ('164', '89', '1_1_1_4_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637653');
INSERT INTO `tb_sys_dialog_stat` VALUES ('165', '90', '1_1_1_4_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637667');
INSERT INTO `tb_sys_dialog_stat` VALUES ('166', '91', '1_2_1_2_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637672');
INSERT INTO `tb_sys_dialog_stat` VALUES ('167', '92', '1_1_1_4_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432637680');
INSERT INTO `tb_sys_dialog_stat` VALUES ('168', '93', '1_1_6_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432637694');
INSERT INTO `tb_sys_dialog_stat` VALUES ('169', '94', '1_1_1_6_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637707');
INSERT INTO `tb_sys_dialog_stat` VALUES ('170', '95', '1_1_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432637725');
INSERT INTO `tb_sys_dialog_stat` VALUES ('171', '96', '1_1_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637741');
INSERT INTO `tb_sys_dialog_stat` VALUES ('172', '97', '1_1_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637755');
INSERT INTO `tb_sys_dialog_stat` VALUES ('173', '98', '1_2_1_3_6', '102', '0', '0', '0', '0', '0', '0', '2', '1432637766');
INSERT INTO `tb_sys_dialog_stat` VALUES ('174', '99', '1_3_8_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637768');
INSERT INTO `tb_sys_dialog_stat` VALUES ('175', '100', '1_1_1_5_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637783');
INSERT INTO `tb_sys_dialog_stat` VALUES ('176', '101', '1_1_1_5_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637796');
INSERT INTO `tb_sys_dialog_stat` VALUES ('177', '102', '1_2_1_3_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432637801');
INSERT INTO `tb_sys_dialog_stat` VALUES ('178', '103', '1_1_1_5_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637821');
INSERT INTO `tb_sys_dialog_stat` VALUES ('179', '104', '1_1_1_4_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637881');
INSERT INTO `tb_sys_dialog_stat` VALUES ('180', '105', '1_3_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637888');
INSERT INTO `tb_sys_dialog_stat` VALUES ('181', '106', '1_2_1_3_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637899');
INSERT INTO `tb_sys_dialog_stat` VALUES ('182', '107', '1_1_1_4_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432637901');
INSERT INTO `tb_sys_dialog_stat` VALUES ('183', '108', '1_1_1_5_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432637925');
INSERT INTO `tb_sys_dialog_stat` VALUES ('184', '109', '1_3_1_1_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432637940');
INSERT INTO `tb_sys_dialog_stat` VALUES ('185', '110', '1_3_5_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432637967');
INSERT INTO `tb_sys_dialog_stat` VALUES ('186', '111', '1_2_1_3_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432637971');
INSERT INTO `tb_sys_dialog_stat` VALUES ('187', '112', '1_3_1_1_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432637993');
INSERT INTO `tb_sys_dialog_stat` VALUES ('188', '113', '1_3_5_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432638044');
INSERT INTO `tb_sys_dialog_stat` VALUES ('189', '114', '1_3_3_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432638073');
INSERT INTO `tb_sys_dialog_stat` VALUES ('190', '115', '1_3_4_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432638104');
INSERT INTO `tb_sys_dialog_stat` VALUES ('191', '116', '1_3_4_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432638135');
INSERT INTO `tb_sys_dialog_stat` VALUES ('192', '117', '1_2_1_3_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432638165');
INSERT INTO `tb_sys_dialog_stat` VALUES ('193', '118', '1_3_3_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432638165');
INSERT INTO `tb_sys_dialog_stat` VALUES ('194', '119', '1_3_3_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432638199');
INSERT INTO `tb_sys_dialog_stat` VALUES ('195', '120', '1_3_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432638263');
INSERT INTO `tb_sys_dialog_stat` VALUES ('196', '121', '1_3_2_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432638317');
INSERT INTO `tb_sys_dialog_stat` VALUES ('197', '122', '1_3_1_3_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432638347');
INSERT INTO `tb_sys_dialog_stat` VALUES ('198', '123', '1_2_1_4_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432638545');
INSERT INTO `tb_sys_dialog_stat` VALUES ('199', '124', '1_3_2_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432638604');
INSERT INTO `tb_sys_dialog_stat` VALUES ('200', '125', '1_2_1_3_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432638634');
INSERT INTO `tb_sys_dialog_stat` VALUES ('201', '126', '1_2_1_3_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432638715');
INSERT INTO `tb_sys_dialog_stat` VALUES ('202', '127', '1_3_1_4_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432638732');
INSERT INTO `tb_sys_dialog_stat` VALUES ('203', '128', '1_2_1_2_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432638770');
INSERT INTO `tb_sys_dialog_stat` VALUES ('204', '129', '1_3_9_8', '101', '0', '0', '0', '0', '0', '0', '2', '1432638791');
INSERT INTO `tb_sys_dialog_stat` VALUES ('205', '130', '1_2_1_3_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432638801');
INSERT INTO `tb_sys_dialog_stat` VALUES ('206', '131', '1_2_1_3_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432638834');
INSERT INTO `tb_sys_dialog_stat` VALUES ('207', '132', '1_3_1_0_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432638849');
INSERT INTO `tb_sys_dialog_stat` VALUES ('208', '133', '1_3_1_0_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432638875');
INSERT INTO `tb_sys_dialog_stat` VALUES ('209', '134', '1_2_9_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432638880');
INSERT INTO `tb_sys_dialog_stat` VALUES ('210', '135', '1_2_9_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432638911');
INSERT INTO `tb_sys_dialog_stat` VALUES ('211', '136', '1_3_1_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432638947');
INSERT INTO `tb_sys_dialog_stat` VALUES ('212', '137', '1_2_9_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432638950');
INSERT INTO `tb_sys_dialog_stat` VALUES ('213', '138', '1_2_9_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432638976');
INSERT INTO `tb_sys_dialog_stat` VALUES ('214', '139', '1_3_9_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432638988');
INSERT INTO `tb_sys_dialog_stat` VALUES ('215', '140', '1_2_9_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432639011');
INSERT INTO `tb_sys_dialog_stat` VALUES ('216', '141', '1_3_9_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432639012');
INSERT INTO `tb_sys_dialog_stat` VALUES ('217', '142', '1_3_9_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432639036');
INSERT INTO `tb_sys_dialog_stat` VALUES ('218', '143', '1_2_1_0_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432639046');
INSERT INTO `tb_sys_dialog_stat` VALUES ('219', '144', '1_3_1_1_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432639064');
INSERT INTO `tb_sys_dialog_stat` VALUES ('220', '145', '1_2_1_0_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432639077');
INSERT INTO `tb_sys_dialog_stat` VALUES ('221', '146', '1_3_1_1_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432639083');
INSERT INTO `tb_sys_dialog_stat` VALUES ('222', '147', '1_3_1_1_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432639103');
INSERT INTO `tb_sys_dialog_stat` VALUES ('223', '148', '1_2_1_0_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432639117');
INSERT INTO `tb_sys_dialog_stat` VALUES ('224', '149', '1_2_1_1_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432639297');
INSERT INTO `tb_sys_dialog_stat` VALUES ('225', '150', '1_2_1_1_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432639329');
INSERT INTO `tb_sys_dialog_stat` VALUES ('226', '151', '1_2_1_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432639356');
INSERT INTO `tb_sys_dialog_stat` VALUES ('227', '152', '1_2_1_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432639384');
INSERT INTO `tb_sys_dialog_stat` VALUES ('228', '153', '1_2_1_4_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432639414');
INSERT INTO `tb_sys_dialog_stat` VALUES ('229', '154', '1_2_1_5_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432639442');
INSERT INTO `tb_sys_dialog_stat` VALUES ('230', '155', '1_2_1_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432639472');
INSERT INTO `tb_sys_dialog_stat` VALUES ('231', '156', '1_2_9', '100', '0', '0', '0', '0', '0', '0', '2', '1432639512');
INSERT INTO `tb_sys_dialog_stat` VALUES ('232', '157', '1_2_1_4_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432639546');
INSERT INTO `tb_sys_dialog_stat` VALUES ('233', '158', '1_2_1_1_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432639578');
INSERT INTO `tb_sys_dialog_stat` VALUES ('234', '159', '1_2_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432639607');
INSERT INTO `tb_sys_dialog_stat` VALUES ('235', '160', '1_2_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432639632');
INSERT INTO `tb_sys_dialog_stat` VALUES ('236', '161', '1_2_8_4', '100', '0', '0', '0', '0', '0', '0', '2', '1432639664');
INSERT INTO `tb_sys_dialog_stat` VALUES ('237', '162', '1_2_8_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432639689');
INSERT INTO `tb_sys_dialog_stat` VALUES ('238', '163', '1_2_8_3', '100', '0', '0', '0', '0', '0', '0', '2', '1432639717');
INSERT INTO `tb_sys_dialog_stat` VALUES ('239', '164', '1_2_8_2', '100', '0', '0', '0', '0', '0', '0', '2', '1432639744');
INSERT INTO `tb_sys_dialog_stat` VALUES ('240', '165', '1_2_8_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432639772');
INSERT INTO `tb_sys_dialog_stat` VALUES ('241', '166', '1_2_8_0', '100', '0', '0', '0', '0', '0', '0', '2', '1432639799');
INSERT INTO `tb_sys_dialog_stat` VALUES ('242', '167', '1_2_6_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432639829');
INSERT INTO `tb_sys_dialog_stat` VALUES ('243', '168', '1_2_6_6', '100', '0', '0', '0', '0', '0', '0', '2', '1432639864');
INSERT INTO `tb_sys_dialog_stat` VALUES ('244', '169', '1_2_7_5', '100', '0', '0', '0', '0', '0', '0', '2', '1432639881');
INSERT INTO `tb_sys_dialog_stat` VALUES ('245', '170', '1_2_2_7', '100', '0', '0', '0', '0', '0', '0', '2', '1432639985');
INSERT INTO `tb_sys_dialog_stat` VALUES ('246', '171', '1_2_3_1', '100', '0', '0', '0', '0', '0', '0', '2', '1432640047');
INSERT INTO `tb_sys_dialog_stat` VALUES ('247', '172', '1_2_2_8', '100', '0', '0', '0', '0', '0', '0', '2', '1432640075');
