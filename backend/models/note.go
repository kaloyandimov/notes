package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:255"`
	UserID      uint   `gorm:"not null"`
}
