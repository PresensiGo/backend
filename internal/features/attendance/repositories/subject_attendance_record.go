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

func (r *SubjectAttendanceRecord) FirstOrCreate(data domains.SubjectAttendanceRecord) (
	*domains.SubjectAttendanceRecord, error,
) {
	var subjectAttendanceRecord models.SubjectAttendanceRecord
	if err := r.db.Where(
		models.SubjectAttendanceRecord{
			SubjectAttendanceId: data.SubjectAttendanceId,
			StudentId:           data.StudentId,
		},
	).Assign(
		models.SubjectAttendanceRecord{
			DateTime: data.DateTime,
			Status:   data.Status,
		},
	).FirstOrCreate(&subjectAttendanceRecord).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSubjectAttendanceRecordModel(&subjectAttendanceRecord), nil
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

func (r *SubjectAttendanceRecord) GetAllByAttendanceId(subjectAttendanceId uint) (
	*[]domains.SubjectAttendanceRecord, error,
) {
	var records []models.SubjectAttendanceRecord
	if err := r.db.Where(
		"subject_attendance_id = ?", subjectAttendanceId,
	).Find(&records).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.SubjectAttendanceRecord, len(records))
		for i, v := range records {
			result[i] = *domains.FromSubjectAttendanceRecordModel(&v)
		}
		return &result, nil
	}
}

func (r *SubjectAttendanceRecord) DeleteByAttendanceIdStudentIdInTx(
	tx *gorm.DB, attendanceId, studentId uint,
) error {
	return tx.Where(
		"subject_attendance_id = ? and student_id = ?", attendanceId, studentId,
	).Unscoped().Delete(&models.SubjectAttendanceRecord{}).Error
}

func (r *SubjectAttendanceRecord) Delete(recordId uint) error {
	return r.db.Where("id = ?", recordId).Unscoped().
		Delete(&models.SubjectAttendanceRecord{}).Error
}
