DROP TABLE IF EXISTS `employee`;

CREATE TABLE `employee` (
  `employee_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `employee_name` varchar(255) NOT NULL,
  `employee_age` int unsigned DEFAULT NULL,
  `employee_gender` enum('male', 'female') DEFAULT NULL,
  `employee_role` varchar(255) DEFAULT NULL,
  `employee_mail` varchar(255) UNIQUE NOT NULL,

   PRIMARY KEY (`employee_id`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `project`;

CREATE TABLE `project` (
  `project_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(1000) DEFAULT NULL,
  `status` varchar(255) NOT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `is_public` boolean DEFAULT true,

  PRIMARY KEY (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `todo` (
  `notes` bigint unsigned NOT NULL AUTO_INCREMENT,
  `project_id` bigint unsigned NOT NULL,
  `name` varchar(1000) NOT NULL,
  `status` varchar(255) NOT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,

  PRIMARY KEY (`todo_id`)
  FOREIGN KEY(`project_id `) REFERENCES project (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `comment` (
  `comment_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `todo_id` bigint unsigned NOT NULL,
  `description` varchar(1000) NOT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,

  PRIMARY KEY (`comment_id`)
  FOREIGN KEY(`todo_id `) REFERENCES todo (`todo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
