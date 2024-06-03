package models

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleID       uint `gorm:"primary_key"`
	PermissionID uint `gorm:"primary_key"`
}
