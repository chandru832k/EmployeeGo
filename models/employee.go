package models

type Employee struct {
	EmployeeId     uint64 `gorm:"primaryKey;column:employee_id;type:bigint(20) unsigned;not null" json:"employee_id"`
	EmployeeName   string `gorm:"column:employee_name;type:varchar(255);default null" json:"employee_name"`
	EmployeeRole   string `gorm:"column:employee_role;type:varchar(255);default null" json:"employee_role"`
	EmployeeMail   string `gorm:"column:employee_mail;type:varchar(255) unique;default null" json:"employee_mail"`
	EmployeeAge    uint   `gorm:"column:employee_age;type:int unsigned;default null" json:"employee_age"`
	EmployeeGender string `gorm:"column:employee_gender;type: enum('male', 'female'); default null" json:"employee_gender"`
}
