package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required" gorm:"unique;size:100"`
	UserName string `json:"user_name" binding:"required" gorm:"unique;size:100"`
	Password string `json:"password" binding:"required" gorm:"not null"`
	RoleID   uint   `json:"role_id"`
	PlanId   uint   `json:"plan_id"`
	TenantID uint   `json:"tenant_id" gorm:"not null"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
