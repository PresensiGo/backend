package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type ClassService struct {
	db *gorm.DB
}

func NewClassService(db *gorm.DB) *ClassService {
	return &ClassService{db}
}

func (s *ClassService) GetAllClasses(majorId uint64) (*responses.GetAllClassesResponse, error) {
	var classes []models.Class
	if err := s.db.Where("major_id = ?", majorId).
		Find(&classes).Error; err != nil {
		return nil, err
	}

	var mappedClasses []dto.Class
	for _, class := range classes {
		mappedClasses = append(
			mappedClasses,
			dto.Class{
				Id:   class.ID,
				Name: class.Name,
			},
		)
	}

	return &responses.GetAllClassesResponse{
		Classes: mappedClasses,
	}, nil
}
