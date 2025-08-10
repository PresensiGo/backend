package domains

import (
	"api/internal/features/attendance/models"
)

type SubjectAttendanceRecord struct {
	Id                  uint `json:"id" validate:"required"`
	SubjectAttendanceId uint `json:"subject_attendance_id" validate:"required"`
	StudentId           uint `json:"student_id" validate:"required"`
}

func FromSubjectAttendanceRecordModel(m *models.SubjectAttendanceRecord) *SubjectAttendanceRecord {
	return &SubjectAttendanceRecord{
		Id:                  m.ID,
		SubjectAttendanceId: m.SubjectAttendanceId,
		StudentId:           m.StudentId,
	}
}

func (s *SubjectAttendanceRecord) ToModel() *models.SubjectAttendanceRecord {
	return &models.SubjectAttendanceRecord{
		SubjectAttendanceId: s.SubjectAttendanceId,
		StudentId:           s.StudentId,
	}
}
