package models

import (
	"gorm.io/gorm"
)

type Category struct {
	// CategoryID uint   `gorm:"primaryKey" json:"category_id"`
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
