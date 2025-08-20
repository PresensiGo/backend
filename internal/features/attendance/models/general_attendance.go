package models

import (
	"time"

	"api/internal/features/school/models"
	userModel "api/internal/features/user/models"
	"gorm.io/gorm"
)

type GeneralAttendance struct {
	gorm.Model

	DateTime  time.Time
	Note      string
	SchoolId  uint
	School    models.School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Code      string        `gorm:"unique"`
	CreatorId uint
	Creator   userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
