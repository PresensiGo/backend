package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type Class struct {
	db *gorm.DB
}

func NewClass(db *gorm.DB) *Class {
	return &Class{db}
}

func (s *Class) GetAllClasses(majorId uint64) (*responses.GetAllClasses, error) {
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

	return &responses.GetAllClasses{
		Classes: mappedClasses,
	}, nil
}
