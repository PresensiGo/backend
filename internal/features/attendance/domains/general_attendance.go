package domains

import (
	"time"

	"api/internal/features/attendance/models"
)

type GeneralAttendance struct {
	Id        uint      `json:"id" validate:"required"`
	DateTime  time.Time `json:"datetime" validate:"required"`
	Note      string    `json:"note" validate:"required"`
	SchoolId  uint      `json:"school_id" validate:"required"`
	Code      string    `json:"code" validate:"required"`
	CreatorId uint      `json:"creator_id" validate:"required"`
} // @name GeneralAttendance

func FromGeneralAttendanceModel(m *models.GeneralAttendance) *GeneralAttendance {
	return &GeneralAttendance{
		Id:        m.ID,
		DateTime:  m.DateTime,
		Note:      m.Note,
		SchoolId:  m.SchoolId,
		Code:      m.Code,
		CreatorId: m.CreatorId,
	}
}

func (g *GeneralAttendance) ToModel() *models.GeneralAttendance {
	return &models.GeneralAttendance{
		DateTime:  g.DateTime,
		Note:      g.Note,
		SchoolId:  g.SchoolId,
		Code:      g.Code,
		CreatorId: g.CreatorId,
	}
}
