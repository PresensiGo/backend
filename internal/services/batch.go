package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type BatchService struct {
	db *gorm.DB
}

func NewBatchService(db *gorm.DB) *BatchService {
	return &BatchService{db}
}

func (s *BatchService) Create(name string) (*responses.CreateBatchResponse, error) {
	batch := models.Batch{
		Name: name,
	}

	if err := s.db.Create(&batch).Error; err != nil {
		return nil, err
	}

	return &responses.CreateBatchResponse{
		Id:   batch.ID,
		Name: batch.Name,
	}, nil
}

func (s *BatchService) GetAll() (*responses.GetAllBatchesResponse, error) {
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

	return &responses.GetAllBatchesResponse{
		Batches: mappedBatches,
	}, nil
}
