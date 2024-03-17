package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"uniqueIndex"`
	Description string
	Rating      float32
	Image       string
}
