package domains

import (
	"time"

	"api/internal/features/attendance/models"
	"api/pkg/constants"
)

type SubjectAttendanceRecord struct {
	Id                  uint                       `json:"id" validate:"required"`
	SubjectAttendanceId uint                       `json:"subject_attendance_id" validate:"required"`
	StudentId           uint                       `json:"student_id" validate:"required"`
	CreatedAt           time.Time                  `json:"created_at" validate:"required"`
	Status              constants.AttendanceStatus `json:"status" validate:"required"`
}

func FromSubjectAttendanceRecordModel(m *models.SubjectAttendanceRecord) *SubjectAttendanceRecord {
	return &SubjectAttendanceRecord{
		Id:                  m.ID,
		SubjectAttendanceId: m.SubjectAttendanceId,
		StudentId:           m.StudentId,
		CreatedAt:           m.CreatedAt,
		Status:              m.Status,
	}
}

func (s *SubjectAttendanceRecord) ToModel() *models.SubjectAttendanceRecord {
	return &models.SubjectAttendanceRecord{
		SubjectAttendanceId: s.SubjectAttendanceId,
		StudentId:           s.StudentId,
		Status:              s.Status,
	}
}
