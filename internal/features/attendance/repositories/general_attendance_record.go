package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
	"gorm.io/gorm"
)

type GeneralAttendanceRecord struct {
	db *gorm.DB
}

func NewGeneralAttendanceRecord(db *gorm.DB) *GeneralAttendanceRecord {
	return &GeneralAttendanceRecord{
		db: db,
	}
}

func (r *GeneralAttendanceRecord) CreateInTx(tx *gorm.DB, data domains.GeneralAttendanceRecord) (
	*domains.GeneralAttendanceRecord, error,
) {
	generalAttendanceRecord := data.ToModel()
	if err := tx.Create(&generalAttendanceRecord).Error; err != nil {
		return nil, err
	} else {
		return domains.FromGeneralAttendanceRecordModel(generalAttendanceRecord), nil
	}
}

func (r *GeneralAttendanceRecord) DeleteByAttendanceIdStudentIdInTx(
	tx *gorm.DB, attendanceId uint, studentId uint,
) error {
	return tx.Where(
		"general_attendance_id = ? AND student_id = ?", attendanceId, studentId,
	).Unscoped().Delete(&models.GeneralAttendanceRecord{}).Error
}
