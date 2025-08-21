package domains

import (
	"time"

	"api/internal/features/attendance/models"
	"api/pkg/constants"
)

type GeneralAttendanceRecord struct {
	Id                  uint                       `json:"id" validate:"required"`
	GeneralAttendanceId uint                       `json:"general_attendance_id" validate:"required"`
	StudentId           uint                       `json:"student_id" validate:"required"`
	DateTime            time.Time                  `json:"date_time" validate:"required"`
	Status              constants.AttendanceStatus `json:"status" validate:"required"`
}

func FromGeneralAttendanceRecordModel(m *models.GeneralAttendanceRecord) *GeneralAttendanceRecord {
	return &GeneralAttendanceRecord{
		Id:                  m.ID,
		GeneralAttendanceId: m.GeneralAttendanceId,
		StudentId:           m.StudentId,
		DateTime:            m.DateTime,
		Status:              m.Status,
	}
}

func (g *GeneralAttendanceRecord) ToModel() *models.GeneralAttendanceRecord {
	return &models.GeneralAttendanceRecord{
		GeneralAttendanceId: g.GeneralAttendanceId,
		StudentId:           g.StudentId,
		DateTime:            g.DateTime,
		Status:              g.Status,
	}
}
