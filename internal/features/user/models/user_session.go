package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model

	UserId    uint
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Token     string `gorm:"uniqueIndex"`
	ExpiresAt time.Time
}
