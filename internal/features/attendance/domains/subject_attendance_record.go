package domains

import (
	"time"

	"api/internal/features/attendance/models"
	"api/pkg/constants"
)

type SubjectAttendanceRecord struct {
	Id                  uint                       `json:"id" validate:"required"`
	DateTime            time.Time                  `json:"date_time" validate:"required"`
	SubjectAttendanceId uint                       `json:"subject_attendance_id" validate:"required"`
	StudentId           uint                       `json:"student_id" validate:"required"`
	Status              constants.AttendanceStatus `json:"status" validate:"required"`
}

func FromSubjectAttendanceRecordModel(m *models.SubjectAttendanceRecord) *SubjectAttendanceRecord {
	return &SubjectAttendanceRecord{
		Id:                  m.ID,
		DateTime:            m.DateTime,
		SubjectAttendanceId: m.SubjectAttendanceId,
		StudentId:           m.StudentId,
		Status:              m.Status,
	}
}

func (s *SubjectAttendanceRecord) ToModel() *models.SubjectAttendanceRecord {
	return &models.SubjectAttendanceRecord{
		DateTime:            s.DateTime,
		SubjectAttendanceId: s.SubjectAttendanceId,
		StudentId:           s.StudentId,
		Status:              s.Status,
	}
}
