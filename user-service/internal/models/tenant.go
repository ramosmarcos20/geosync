package models

import (
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null; size:100"`
	TaxNumber      string `json:"tax_number" gorm:"not null; size:13"`
	Email          string `json:"email" gorm:"not null; size:100"`
	Phone          string `json:"phone" gorm:"not null; size:20"`
	Address        string `json:"address" gorm:"size:200"`
	IsActive       bool   `json:"is_active" gorm:"default:true"`
	AdditionalInfo string `json:"additional_info"`
}
