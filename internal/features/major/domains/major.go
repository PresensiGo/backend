package domains

import (
	"api/internal/features/major/models"
)

type Major struct {
	Id      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	BatchId uint   `json:"batch_id" validate:"required"`
} // @name major

func FromMajorModel(m *models.Major) *Major {
	return &Major{
		Id:      m.ID,
		Name:    m.Name,
		BatchId: m.BatchId,
	}
}

func (m *Major) ToModel() *models.Major {
	return &models.Major{
		Name:    m.Name,
		BatchId: m.BatchId,
	}
}
