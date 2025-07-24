package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	NIS         string `gorm:"not null"`
	Name        string `gorm:"not null"`
	ClassroomID uint   `gorm:"not null"`
	Classroom   Classroom
}
