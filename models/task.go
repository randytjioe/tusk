package models

import "time"

type Task struct {
	Id int `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	UserId int `gorm:"type:int" json:"user_id"`
	Title string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Status string `gorm:"type:varchar(50)" json:"status"`
	Reason string `gorm:"type:text; default: json:"reason"`
	Revision int8 `gorm:"type:int; default:0" json:"revision"`
	DueDate string `gorm:"type:varchar(50)" json:"dueDate"`
	RejectedDate string `gorm:"type:varchar(50)" json:"rejectedDate"`
	ApprovedDate string `gorm:"type:varchar(50)" json:"approvedDate"`
	Attachment string `gorm:"type:varchar(255)" json:"attachment"`
	CreatedAt time.Time ` json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
	User User `gorm:"foreignKey:UserId" json:"user,omitempty"`
}