-- 若不存在则创建数据库 blog, 字符编码为 utf8, 校验规则为 utf8_general_ci 不区分大小写
CREATE DATABASE IF NOT EXISTS blog CHARSET utf8 COLLATE utf8_general_ci;
-- 进入数据库 blog
USE `blog`;
-- 标签表
CREATE TABLE IF NOT EXISTS `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 禁用，1 启用',
    `created_by` varchar(100) DEFAULT  '' COMMENT '创建人',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,
    PRIMARY KEY (id)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '文章标签管理';
-- 文章表
CREATE TABLE IF NOT EXISTS `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id` int(10) unsigned DEFAULT '0' COMMENT 'tag id',
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `description` varchar(255) DEFAULT '' COMMENT '文章描述',
    `content` text COMMENT '文章内容',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 禁用，1 启用',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,
    PRIMARY KEY (id)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '文章管理';
-- 认证表
CREATE TABLE IF NOT EXISTS `blog_auth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) DEFAULT '' COMMENT '用户名',
    `password` varchar(50) DEFAULT '' COMMENT '密码',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,
    PRIMARY KEY (id)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '认证管理';
-- 插入 admin
INSERT INTO `blog`.`blog_auth` (`username`,`password`,`created_at`) VALUES ('admin','admin',NOW());