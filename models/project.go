package models

type Project struct {
	ProjectId     	uint64 `gorm:"primaryKey;column:project_id;type:bigint(20) unsigned;not null" json:"project_id"`
	Name   			string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Status   		string `gorm:"column:status;type:varchar(255);not null" json:"status"`
	Description string `gorm:"column:description;type:varchar(1000);default null" json:"description"`
  CreatedBy   string `gorm:"column:created_by;type:varchar(255) unique;default null" json:"created_by"`
	CreatedAt   uint64 `gorm:"column:created_at;type:bigint unsigned;default null" json:"created_at"`
	UpdatedBy 	string `gorm:"column:updated_by;ype:varchar(255) unique;default null" json:"updated_by"`
	UpdatedAt   uint64 `gorm:"column:updated_at;type:bigint unsigned;default null" json:"updated_at"`
}

