DROP DATABASE IF EXISTS cross_field_shop;

CREATE DATABASE cross_field_shop;

use cross_field_shop;

-- ----------------------------

-- Table structure for users

-- ----------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `email` varchar(64) NOT NULL COMMENT '邮箱',
    `password` varchar(20) DEFAULT '' COMMENT '密码',
    `name` varchar(20) DEFAULT '',
    `avatar` varchar(100) DEFAULT '',
    `role` int(8) DEFAULT 0 COMMENT '权限 0-无权限用户 1-管理员',
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- --------------------------------

-- Table structure for commodities

-- --------------------------------

DROP TABLE IF EXISTS `commodities`;

CREATE TABLE `commodities` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '',
    `info` varchar(100) DEFAULT '',
    `price` decimal(10,2) DEFAULT 0,
    `pictures` TEXT,
    `video` varchar(100) DEFAULT '',
    `remain` int(11) DEFAULT  0,
    `sale` int(11) DEFAULT 0,
    `tag` varchar(50) DEFAULT '',
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`),
    FULLTEXT KEY (`name`, `info`) WITH PARSER ngram
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- -----------------------------

-- Table structure for comments

-- -----------------------------

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `info` varchar(100) DEFAULT '',
    `pictures` TEXT,
    `tag` varchar(50) DEFAULT '' COMMENT 'tag id list',
    `time` DATETIME DEFAULT NULL,
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ---------------------------

-- Table structure for tags

-- ---------------------------

DROP TABLE IF EXISTS `tags`;

CREATE TABLE `tags` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `tag` varchar(50) NOT NULL,
    `type` int(8) UNSIGNED DEFAULT 0,
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`),
    UNIQUE KEY tag_name (`tag`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- --------------------------------

-- Table structure for purchase

-- --------------------------------

DROP TABLE IF EXISTS `purchase`;

CREATE TABLE `purchase` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(11) UNSIGNED NOT NULL,
    `commodity_id` int(11) UNSIGNED NOT NULL,
    `number` int(11) UNSIGNED NOT NULL DEFAULT 1,
    `price` decimal(10,2) NOT NULL DEFAULT 0,
    `status` int(8) UNSIGNED DEFAULT 0,
    `logistics` TEXT,
    `time` DATETIME DEFAULT NULL,
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------

-- Table structure for history

-- ----------------------------

DROP TABLE IF EXISTS `history`;

CREATE TABLE `history` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(11) UNSIGNED NOT NULL,
    `commodity_id` int(11) UNSIGNED NOT NULL,
    `time` DATETIME DEFAULT NULL,
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------------

-- Table structure for trend_back_up

-- ----------------------------------

DROP TABLE IF EXISTS `trend_back_up`;

CREATE TABLE `trend_back_up` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `list` TEXT,
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- --------------------------------

-- Table structure for consignee

-- --------------------------------

DROP TABLE IF EXISTS `consignee`;

CREATE TABLE `consignee` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(11) UNSIGNED NOT NULL,
    `address` TEXT NOT NULL,
    `name` varchar(100) NOT NULL,
    `tel` varchar(20) NOT NULL DEFAULT 1,
    `tag` int(11) UNSIGNED,
    `re` int(8) DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;
