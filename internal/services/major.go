package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type MajorService struct {
	db *gorm.DB
}

func NewMajorService(db *gorm.DB) *MajorService {
	return &MajorService{db}
}

func (s *MajorService) GetAllMajors(batchId uint64) (*responses.GetAllMajorsResponse, error) {
	var majors []models.Major
	if err := s.db.Where("batch_id = ?", batchId).
		Find(&majors).
		Error; err != nil {
		return nil, err
	}

	var mappedMajors []dto.Major
	for _, major := range majors {
		mappedMajors = append(mappedMajors, dto.Major{
			Id:   major.ID,
			Name: major.Name,
		})
	}

	return &responses.GetAllMajorsResponse{
		Majors: mappedMajors,
	}, nil
}
