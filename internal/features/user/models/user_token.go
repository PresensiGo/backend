package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model

	UserId       uint
	User         User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RefreshToken string    `gorm:"unique"`
	TTL          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
