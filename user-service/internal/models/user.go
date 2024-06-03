package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	UserName string `json:"user_name" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
	RoleID   uint   `json:"role_id"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
