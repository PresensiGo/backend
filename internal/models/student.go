package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	NIS     string `gorm:"not null"`
	Name    string `gorm:"not null"`
	ClassId uint   `gorm:"not null"`
	Class   Class
}
