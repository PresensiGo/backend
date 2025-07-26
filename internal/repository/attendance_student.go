package repository

import (
	"api/internal/dto"
	"api/internal/models"
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

func (r *AttendanceStudent) DeleteByAttendanceID(tx *gorm.DB, attendanceID uint) error {
	return tx.Where("attendance_id = ?", attendanceID).
		Unscoped().
		Delete(&models.AttendanceStudent{}).
		Error
}
