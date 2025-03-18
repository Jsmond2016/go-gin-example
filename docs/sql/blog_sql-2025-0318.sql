-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        9.1.0 - MySQL Community Server - GPL
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  12.10.0.7000
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 blog 的数据库结构
CREATE DATABASE IF NOT EXISTS `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `blog`;

-- 导出  表 blog.blog_article 结构
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE IF NOT EXISTS `blog_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` bigint unsigned DEFAULT NULL,
  `title` longtext,
  `desc` longtext,
  `content` longtext,
  `cover_image_url` longtext,
  `created_by` longtext,
  `modified_by` longtext,
  `state` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章管理';

-- 正在导出表  blog.blog_article 的数据：~1 rows (大约)
DELETE FROM `blog_article`;
INSERT INTO `blog_article` (`id`, `tag_id`, `title`, `desc`, `content`, `cover_image_url`, `created_by`, `modified_by`, `state`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 3, '创建文章创建文章创建文章创建文章创建文章', '1111111111111111111', '创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章创建文章', 'http://www.baidu.com', 'test', '', 1, '2025-03-18 21:22:06.451', '2025-03-18 21:22:06.451', NULL);

-- 导出  表 blog.blog_auth 结构
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE IF NOT EXISTS `blog_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  blog.blog_auth 的数据：~1 rows (大约)
DELETE FROM `blog_auth`;
INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES
	(1, 'test', 'test123');

-- 导出  表 blog.blog_tag 结构
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE IF NOT EXISTS `blog_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `created_on` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `created_by` longtext,
  `modified_by` longtext,
  `state` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blog_tag_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章标签管理';

-- 正在导出表  blog.blog_tag 的数据：~3 rows (大约)
DELETE FROM `blog_tag`;
INSERT INTO `blog_tag` (`id`, `name`, `created_on`, `created_by`, `modified_by`, `state`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, '111', 0, 'test', '', 1, '2025-03-18 21:16:18.023', '2025-03-18 21:16:18.023', NULL),
	(2, '2222', 0, 'test', '', 1, '2025-03-18 21:16:26.283', '2025-03-18 21:16:26.283', NULL),
	(3, '333', 0, 'test', '', 1, '2025-03-18 21:16:36.584', '2025-03-18 21:16:36.584', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
