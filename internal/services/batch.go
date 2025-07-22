package services

import (
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db}
}

func (s *Service) Create(name string) (*responses.CreateBatchResponse, error) {
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
