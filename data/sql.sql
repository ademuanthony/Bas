-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.7.17-0ubuntu0.16.04.1 - (Ubuntu)
-- Server OS:                    Linux
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for bas
CREATE DATABASE IF NOT EXISTS `bas` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `bas`;

-- Dumping structure for table bas.application
CREATE TABLE IF NOT EXISTS `application` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- Dumping data for table bas.application: ~0 rows (approximately)
/*!40000 ALTER TABLE `application` DISABLE KEYS */;
INSERT INTO `application` (`id`, `name`, `description`, `created_by`, `created_date`, `updated_by`, `updated_date`) VALUES
	(1, 'CRM', 'Customer relation management system', 1, '2017-04-01 22:04:40', 0, NULL);
/*!40000 ALTER TABLE `application` ENABLE KEYS */;

-- Dumping structure for table bas.resource
CREATE TABLE IF NOT EXISTS `resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(128) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `application_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK__application` (`application_id`),
  CONSTRAINT `FK__application` FOREIGN KEY (`application_id`) REFERENCES `application` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- Dumping data for table bas.resource: ~2 rows (approximately)
/*!40000 ALTER TABLE `resource` DISABLE KEYS */;
INSERT INTO `resource` (`id`, `key`, `created_by`, `created_date`, `updated_by`, `updated_date`, `application_id`) VALUES
	(1, 'crm_foo_bar', 1, '2017-04-03 10:44:29', 1, '2017-04-03 10:44:29', 1),
	(2, 'crm_foo_boo', 1, '2017-04-03 13:53:49', 1, '2017-04-03 13:53:49', 1);
/*!40000 ALTER TABLE `resource` ENABLE KEYS */;

-- Dumping structure for table bas.role
CREATE TABLE IF NOT EXISTS `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_role_user_creator` (`created_by`),
  KEY `FK_role_user_updater` (`updated_by`),
  CONSTRAINT `FK_role_user_creator` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  CONSTRAINT `FK_role_user_updater` FOREIGN KEY (`updated_by`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- Dumping data for table bas.role: ~2 rows (approximately)
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` (`id`, `name`, `created_date`, `created_by`, `updated_date`, `updated_by`) VALUES
	(2, 'Super Admin', '2017-04-07 04:06:27', 1, '2017-04-07 04:06:27', 1),
	(3, 'Admin', '2017-04-07 04:43:59', 1, '2017-04-07 04:43:59', 1),
	(4, 'User', '2017-04-07 04:44:09', 1, '2017-04-07 04:44:09', 1);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;

-- Dumping structure for table bas.role_resource
CREATE TABLE IF NOT EXISTS `role_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  KEY `id` (`id`),
  KEY `FK__role_resource_role` (`role_id`),
  KEY `FK__role_resource_resource` (`resource_id`),
  CONSTRAINT `FK__role_resource_resource` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`),
  CONSTRAINT `FK__role_resource_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Dumping data for table bas.role_resource: ~2 rows (approximately)
/*!40000 ALTER TABLE `role_resource` DISABLE KEYS */;
INSERT INTO `role_resource` (`id`, `role_id`, `resource_id`) VALUES
	(2, 2, 1),
	(3, 2, 2);
/*!40000 ALTER TABLE `role_resource` ENABLE KEYS */;

-- Dumping structure for table bas.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) DEFAULT NULL,
  `password_hash` varchar(256) NOT NULL,
  `created_by` int(11) DEFAULT NULL,
  `created_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_by` int(11) DEFAULT NULL,
  `updated_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Dumping data for table bas.user: ~2 rows (approximately)
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `username`, `password`, `password_hash`, `created_by`, `created_date`, `updated_by`, `updated_date`, `first_name`, `last_name`, `email`) VALUES
	(1, 'tony', 'ojima123', '$2a$10$q8S7t0jabuY3dgZcX60FKuJQxWADrCB91/XbDhBoBQi87qWRUzuai', 0, '2017-03-31 18:41:22', 0, '2017-03-31 18:41:22', 'Anthony', 'Ademu', 'blenyo11@gmail.com'),
	(2, 'tony', 'ojima123', '$2a$10$yXhgn4TZtLV0DLfH1V0.Fe3hhu8LF/bx.bAI4h9Sg9.TLzdzXcc.i', 0, '2017-03-31 18:43:55', 0, '2017-03-31 18:43:55', 'Anthony', 'Ademu', 'blenyo11@gmail.com'),
	(3, 'blenyo', 'ojima123', '$2a$10$/b0LJUhtn2SUPfqdNwDuNev9dk5iaQqw7h2XMuOw2efpwIjhrc3N2', 0, '2017-04-01 04:44:31', 0, '2017-04-01 04:44:31', 'Anthony', 'Ademu', 'ademuanthony@gmail.com');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

-- Dumping structure for table bas.user_role
CREATE TABLE IF NOT EXISTS `user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  `created_by` int(11) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK__user_role_user` (`user_id`),
  KEY `FK__user_role_role` (`role_id`),
  CONSTRAINT `FK__user_role_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
  CONSTRAINT `FK__user_role_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- Dumping data for table bas.user_role: ~3 rows (approximately)
/*!40000 ALTER TABLE `user_role` DISABLE KEYS */;
INSERT INTO `user_role` (`id`, `user_id`, `role_id`, `created_by`, `created_date`) VALUES
	(2, 1, 2, 0, NULL),
	(3, 1, 3, 0, NULL),
	(4, 1, 4, 0, NULL);
/*!40000 ALTER TABLE `user_role` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
