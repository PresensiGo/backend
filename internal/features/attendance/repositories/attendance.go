package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
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
	attendance *domains.Attendance,
) error {
	return tx.Create(&attendance).Error
}

func (r *Attendance) GetAll(classroomID uint) (*[]domains.Attendance, error) {
	var attendances []models.Attendance
	if err := r.db.
		Where("classroom_id = ?", classroomID).
		Order("date asc").
		Find(&attendances).Error; err != nil {
		return nil, err
	}

	mappedAttendances := make([]domains.Attendance, len(attendances))
	for i, attendance := range attendances {
		mappedAttendances[i] = domains.Attendance{
			Id:          attendance.ID,
			ClassroomId: attendance.ClassroomId,
			Date:        attendance.Date,
		}
	}

	return &mappedAttendances, nil
}

func (r *Attendance) GetById(attendanceId uint) (*domains.Attendance, error) {
	var attendance models.Attendance
	if err := r.db.Where("id = ?", attendanceId).
		First(&attendance).
		Error; err != nil {
		return nil, err
	}

	return &domains.Attendance{
		Id:          attendance.ID,
		ClassroomId: attendance.ClassroomId,
		Date:        attendance.Date,
	}, nil
}

func (r *Attendance) DeleteByID(tx *gorm.DB, attendanceID uint) error {
	return tx.Where("id = ?", attendanceID).
		Unscoped().
		Delete(&models.Attendance{}).
		Error
}
