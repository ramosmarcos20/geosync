package models

import (
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name" gorm:"not null; size:100"`
	IsFree       bool   `json:"is_free" gorm:"not null"`
	Monthly      int    `json:"monthly"`
	Yearly       int    `json:"yearly"`
	NumberUser   int    `json:"number_user" gorm:"not null"`
	NumberDevice int    `json:"number_device" gorm:"not null"`
	Access       string `json:"access"`
	IsActive     bool   `json:"is_active" gorm:"default:true"`
}
