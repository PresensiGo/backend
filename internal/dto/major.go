package dto

import (
	"api/internal/models"
)

type Major struct {
	Id      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	BatchId uint   `json:"batch_id" validate:"required"`
}

func FromMajorModel(model *models.Major) *Major {
	return &Major{
		Id:      model.ID,
		Name:    model.Name,
		BatchId: model.BatchId,
	}
}

func (m *Major) ToModel() *models.Major {
	return &models.Major{
		Name:    m.Name,
		BatchId: m.BatchId,
	}
}
