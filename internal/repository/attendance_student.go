package repository

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type AttendanceStudent struct {
	db *gorm.DB
}

func NewAttendanceStudent(
	db *gorm.DB,
) *AttendanceStudent {
	return &AttendanceStudent{db}
}

func (r *AttendanceStudent) CreateBatch(
	tx *gorm.DB,
	attendanceStudents *[]dto.AttendanceDetail,
) error {
	return tx.Create(attendanceStudents).Error
}

func (r *AttendanceStudent) GetAllByAttendanceId(attendanceId uint) (*[]dto.AttendanceDetail, error) {
	var attendanceStudents []models.AttendanceDetail
	if err := r.db.Model(&models.AttendanceDetail{}).
		Where("attendance_id = ?", attendanceId).
		Find(&attendanceStudents).Error; err != nil {
		return nil, err
	}

	mappedAttendanceStudents := make([]dto.AttendanceDetail, len(attendanceStudents))
	for i, item := range attendanceStudents {
		mappedAttendanceStudents[i] = dto.AttendanceDetail{
			ID:           item.ID,
			AttendanceID: item.AttendanceId,
			StudentID:    item.StudentId,
			Status:       item.Status,
			Note:         item.Note,
		}
	}

	return &mappedAttendanceStudents, nil
}

func (r *AttendanceStudent) DeleteByAttendanceID(tx *gorm.DB, attendanceID uint) error {
	return tx.Where("attendance_id = ?", attendanceID).
		Unscoped().
		Delete(&models.AttendanceDetail{}).
		Error
}
