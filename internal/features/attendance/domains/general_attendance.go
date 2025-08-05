package domains

import (
	"time"

	"api/internal/features/attendance/models"
)

type GeneralAttendance struct {
	Id       uint      `json:"id" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
	DueTime  time.Time `json:"due_time" validate:"required"`
	Note     string    `json:"note" validate:"required"`
	SchoolId uint      `json:"school_id" validate:"required"`
}

func FromGeneralAttendanceModel(m *models.GeneralAttendance) *GeneralAttendance {
	return &GeneralAttendance{
		Id:       m.ID,
		Date:     m.Date,
		DueTime:  m.DueTime,
		Note:     m.Note,
		SchoolId: m.SchoolId,
	}
}

func (g *GeneralAttendance) ToModel() *models.GeneralAttendance {
	return &models.GeneralAttendance{
		Date:     g.Date,
		DueTime:  g.DueTime,
		Note:     g.Note,
		SchoolId: g.SchoolId,
	}
}
