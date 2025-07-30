package dto

import (
	"api/internal/models"
)

type Batch struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	SchoolId uint   `json:"school_id" validate:"required"`
}

func FromBatchModel(model *models.Batch) *Batch {
	return &Batch{
		Id:       model.ID,
		Name:     model.Name,
		SchoolId: model.SchoolId,
	}
}

func (b *Batch) ToModel() *models.Batch {
	return &models.Batch{
		Name:     b.Name,
		SchoolId: b.SchoolId,
	}
}
