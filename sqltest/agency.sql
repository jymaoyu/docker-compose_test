/*
Navicat MySQL Data Transfer

Source Server         : 123
Source Server Version : 50716
Source Host           : localhost:3306
Source Database       : medidb

Target Server Type    : MYSQL
Target Server Version : 50716
File Encoding         : 65001

Date: 2018-05-18 10:33:08
*/
CREATE DATABASE IF NOT EXISTS `medidb`;
USE `medidb`;
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for agency
-- ----------------------------
DROP TABLE IF EXISTS `agency`;
CREATE TABLE `agency` (
  `ano` char(8) NOT NULL,
  `aname` char(8) NOT NULL,
  `asex` char(8) NOT NULL,
  `aphone` char(12) DEFAULT NULL,
  `aremark` varchar(50) DEFAULT NULL COMMENT 'Note',
  PRIMARY KEY (`ano`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of agency
-- ----------------------------
INSERT INTO `agency` VALUES ('1', 'lisi', 'man', '10010', null);
INSERT INTO `agency` VALUES ('2', 'zhenba', 'woman', '95533', null);
INSERT INTO `agency` VALUES ('3', 'zhouba', 'man', '22222', '666');
INSERT INTO `agency` VALUES ('4', 'liuliu', 'woman', '122345', '');
INSERT INTO `agency` VALUES ('5', 'sunwu', 'woman', '23345', '');
INSERT INTO `agency` VALUES ('6', 'huangnan', 'man', '88776', '');
SET FOREIGN_KEY_CHECKS=1;
