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
