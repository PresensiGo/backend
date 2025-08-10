package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentToken struct {
	gorm.Model

	StudentId    uint
	Student      Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeviceId     string
	RefreshToken string    `gorm:"unique"`
	TTL          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
