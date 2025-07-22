package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type Batch struct {
	db *gorm.DB
}

func NewBatch(db *gorm.DB) *Batch {
	return &Batch{db}
}

func (s *Batch) Create(name string) (*responses.CreateBatch, error) {
	batch := models.Batch{
		Name: name,
	}

	if err := s.db.Create(&batch).Error; err != nil {
		return nil, err
	}

	return &responses.CreateBatch{
		Id:   batch.ID,
		Name: batch.Name,
	}, nil
}

func (s *Batch) GetAll() (*responses.GetAllBatches, error) {
	var batches []models.Batch
	if err := s.db.Find(&batches).Error; err != nil {
		return nil, err
	}

	var mappedBatches []dto.Batch
	for _, batch := range batches {
		mappedBatches = append(
			mappedBatches,
			dto.Batch{
				Id:   batch.ID,
				Name: batch.Name,
			},
		)
	}

	return &responses.GetAllBatches{
		Batches: mappedBatches,
	}, nil
}
