package dto

import (
	"api/internal/models"
	"time"
)

type Lateness struct {
	Id       uint      `json:"id" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
	SchoolId uint      `json:"school_id" validate:"required"`
}

func FromLatenessModel(model *models.Lateness) *Lateness {
	return &Lateness{
		Id:       model.ID,
		Date:     model.Date,
		SchoolId: model.SchoolId,
	}
}

func (l *Lateness) ToModel() *models.Lateness {
	return &models.Lateness{
		Date:     l.Date,
		SchoolId: l.SchoolId,
	}
}
