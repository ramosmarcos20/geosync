package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name; size:50"`
}
