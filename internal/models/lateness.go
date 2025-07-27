package models

import (
	"gorm.io/gorm"
	"time"
)

type Lateness struct {
	gorm.Model

	Date     time.Time
	SchoolId uint
	School   School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
