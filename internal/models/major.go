package models

import "gorm.io/gorm"

type Major struct {
	gorm.Model

	Name    string `gorm:"not null"`
	BatchId uint   `gorm:"not null"`
	Batch   Batch
}
