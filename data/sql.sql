CREATE TABLE `users` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(50) NOT NULL,
	`password` VARCHAR(50) NULL DEFAULT NULL,
	`password_hash` VARCHAR(50) NOT NULL,
	`created_by` INT(11) NULL DEFAULT NULL,
	`created_date` DATETIME NULL DEFAULT NULL,
	`updated_by` INT(11) NULL DEFAULT NULL,
	`updated_date` DATETIME NULL DEFAULT NULL,
	`first_name` VARCHAR(50) NULL DEFAULT NULL,
	`last_name` VARCHAR(50) NULL DEFAULT NULL,
	`email` VARCHAR(128) NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
;

CREATE TABLE `application` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(128) NULL DEFAULT NULL,
	`description` VARCHAR(500) NULL DEFAULT NULL,
	`created_by` INT(11) NULL DEFAULT NULL,
	`created_date` DATETIME NULL DEFAULT NULL,
	`updated_by` INT(11) NULL DEFAULT NULL,
	`updated_date` DATETIME NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
AUTO_INCREMENT=2
;



CREATE TABLE `resource` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`key` VARCHAR(128) NULL DEFAULT NULL,
	`created_by` INT(11) NULL DEFAULT NULL,
	`created_date` DATETIME NULL DEFAULT NULL,
	`updated_by` INT(11) NULL DEFAULT NULL,
	`updated_date` DATETIME NULL DEFAULT NULL,
	`application_id` INT(11) NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `FK__application` (`application_id`),
	CONSTRAINT `FK__application` FOREIGN KEY (`application_id`) REFERENCES `application` (`id`)
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
;