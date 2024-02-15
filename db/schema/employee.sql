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
