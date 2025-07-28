package models

import (
	"time"

	"gorm.io/gorm"
)

type Lateness struct {
	gorm.Model

	Date     time.Time
	SchoolId uint
	School   School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
