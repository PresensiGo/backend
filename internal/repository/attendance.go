package repository

import (
	"api/internal/dto"
	"gorm.io/gorm"
)

type Attendance struct {
	db *gorm.DB
}

func NewAttendance(db *gorm.DB) *Attendance {
	return &Attendance{db}
}

func (r *Attendance) Create(
	tx *gorm.DB,
	attendance *dto.Attendance,
) error {
	return tx.Create(&attendance).Error
}
