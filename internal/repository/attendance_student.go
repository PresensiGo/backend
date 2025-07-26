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
	attendanceStudents *[]dto.AttendanceStudent,
) error {
	return tx.Create(attendanceStudents).Error
}

func (r *AttendanceStudent) GetAllByAttendanceId(attendanceId uint) (*[]dto.AttendanceStudent, error) {
	var attendanceStudents []models.AttendanceStudent
	if err := r.db.Model(&models.AttendanceStudent{}).
		Where("attendance_id = ?", attendanceId).
		Find(&attendanceStudents).Error; err != nil {
		return nil, err
	}

	mappedAttendanceStudents := make([]dto.AttendanceStudent, len(attendanceStudents))
	for i, item := range attendanceStudents {
		mappedAttendanceStudents[i] = dto.AttendanceStudent{
			ID:           item.ID,
			AttendanceID: item.AttendanceID,
			StudentID:    item.StudentID,
			Status:       item.Status,
			Note:         item.Note,
		}
	}

	return &mappedAttendanceStudents, nil
}

func (r *AttendanceStudent) DeleteByAttendanceID(tx *gorm.DB, attendanceID uint) error {
	return tx.Where("attendance_id = ?", attendanceID).
		Unscoped().
		Delete(&models.AttendanceStudent{}).
		Error
}
