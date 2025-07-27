package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	NIS         string `gorm:"not null"`
	Name        string `gorm:"not null"`
	ClassroomId uint
	Classroom   Classroom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
