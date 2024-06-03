package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint         `json:"id" gorm:"primary_key"`
	Name        string       `json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
