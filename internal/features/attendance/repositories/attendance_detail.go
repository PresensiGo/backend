package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
	"gorm.io/gorm"
)

type AttendanceDetail struct {
	db *gorm.DB
}

func NewAttendanceStudent(
	db *gorm.DB,
) *AttendanceDetail {
	return &AttendanceDetail{db}
}

func (r *AttendanceDetail) CreateBatch(
	tx *gorm.DB,
	attendanceStudents *[]domains.AttendanceDetail,
) error {
	return tx.Create(attendanceStudents).Error
}

func (r *AttendanceDetail) GetAllByAttendanceId(attendanceId uint) (
	*[]domains.AttendanceDetail, error,
) {
	var attendanceStudents []models.AttendanceDetail
	if err := r.db.Model(&models.AttendanceDetail{}).
		Where("attendance_id = ?", attendanceId).
		Find(&attendanceStudents).Error; err != nil {
		return nil, err
	}

	mappedAttendanceStudents := make([]domains.AttendanceDetail, len(attendanceStudents))
	for i, item := range attendanceStudents {
		mappedAttendanceStudents[i] = domains.AttendanceDetail{
			Id:           item.ID,
			AttendanceId: item.AttendanceId,
			StudentId:    item.StudentId,
			Status:       item.Status,
			Note:         item.Note,
		}
	}

	return &mappedAttendanceStudents, nil
}

func (r *AttendanceDetail) DeleteByAttendanceID(tx *gorm.DB, attendanceID uint) error {
	return tx.Where("attendance_id = ?", attendanceID).
		Unscoped().
		Delete(&models.AttendanceDetail{}).
		Error
}
