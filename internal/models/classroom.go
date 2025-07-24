package models

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model

	Name    string `gorm:"not null"`
	MajorId uint   `gorm:"not null"`
	Major   Major
}
