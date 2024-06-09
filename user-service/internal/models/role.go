package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"unique; not null; size:100"`
	Description string `json:"description; size:150"`
	IsActive    bool   `json:"is_active" gorm:"default:true; not null"`
}
