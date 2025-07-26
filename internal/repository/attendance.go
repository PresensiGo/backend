package repository

import (
	"api/internal/dto"
	"api/internal/models"
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

func (r *Attendance) GetAll(classroomID uint) (*[]dto.Attendance, error) {
	var attendances []models.Attendance
	if err := r.db.
		Where("classroom_id = ?", classroomID).
		Order("date asc").
		Find(&attendances).Error; err != nil {
		return nil, err
	}

	mappedAttendances := make([]dto.Attendance, len(attendances))
	for i, attendance := range attendances {
		mappedAttendances[i] = dto.Attendance{
			ID:          attendance.ID,
			ClassroomID: attendance.ClassroomID,
			Date:        attendance.Date,
		}
	}

	return &mappedAttendances, nil
}

func (r *Attendance) DeleteByID(tx *gorm.DB, attendanceID uint) error {
	return tx.Where("id = ?", attendanceID).
		Unscoped().
		Delete(&models.Attendance{}).
		Error
}
