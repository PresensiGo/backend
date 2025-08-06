package domains

import (
	"time"

	"api/internal/features/attendance/models"
)

type SubjectAttendance struct {
	Id          uint      `json:"id" validate:"required"`
	DateTime    time.Time `json:"date_time" validate:"required"`
	Code        string    `json:"code" validate:"required"`
	Note        string    `json:"note" validate:"required"`
	ClassroomId uint      `json:"classroom_id" validate:"required"`
	SubjectId   uint      `json:"subject_id" validate:"required"`
}

func FromSubjectAttendanceModel(m *models.SubjectAttendance) *SubjectAttendance {
	return &SubjectAttendance{
		Id:          m.ID,
		DateTime:    m.DateTime,
		Code:        m.Code,
		Note:        m.Note,
		ClassroomId: m.ClassroomId,
		SubjectId:   m.SubjectId,
	}
}

func (s *SubjectAttendance) ToModel() *models.SubjectAttendance {
	return &models.SubjectAttendance{
		DateTime:    s.DateTime,
		Code:        s.Code,
		Note:        s.Note,
		ClassroomId: s.ClassroomId,
		SubjectId:   s.SubjectId,
	}
}
