package models

import "gorm.io/gorm"

type Batch struct {
	gorm.Model

	Name     string `gorm:"not null"`
	SchoolId uint
	School   School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
