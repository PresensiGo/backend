package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
	"gorm.io/gorm"
)

type SubjectAttendanceRecord struct {
	db *gorm.DB
}

func NewSubjectAttendanceRecord(db *gorm.DB) *SubjectAttendanceRecord {
	return &SubjectAttendanceRecord{
		db: db,
	}
}

func (r *SubjectAttendanceRecord) CreateInTx(
	tx *gorm.DB, data domains.SubjectAttendanceRecord,
) (*domains.SubjectAttendanceRecord, error) {
	subjectAttendanceRecord := data.ToModel()
	if err := tx.Create(&subjectAttendanceRecord).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSubjectAttendanceRecordModel(subjectAttendanceRecord), nil
	}
}

func (r *SubjectAttendanceRecord) DeleteByAttendanceIdStudentIdInTx(
	tx *gorm.DB, attendanceId, studentId uint,
) error {
	return tx.Where(
		"subject_attendance_id = ? and student_id = ?", attendanceId, studentId,
	).Unscoped().Delete(&models.SubjectAttendanceRecord{}).Error
}
