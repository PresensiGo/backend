package repository

import (
	"api/internal/dto"
	"gorm.io/gorm"
)

type AttendanceStudent struct{}

func NewAttendanceStudent() *AttendanceStudent {
	return &AttendanceStudent{}
}

func (r *AttendanceStudent) CreateBatch(
	tx *gorm.DB,
	attendanceStudents *[]dto.AttendanceStudent,
) error {
	return tx.Create(attendanceStudents).Error
}
