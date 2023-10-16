package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	Articles []Article
}
