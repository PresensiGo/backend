package models

import (
	"gorm.io/gorm"
	"time"
)

type UserToken struct {
	gorm.Model

	UserId       uint
	User         User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RefreshToken string    `gorm:"unique"`
	LastLogin    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
