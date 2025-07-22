package models

import "gorm.io/gorm"

type UserToken struct {
	gorm.Model

	UserId       uint
	User         User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RefreshToken string
}
