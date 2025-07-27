package repositories

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type School struct {
	db *gorm.DB
}

func NewSchool(db *gorm.DB) *School {
	return &School{db}
}

func (s *School) GetById(id uint) (*dto.School, error) {
	var school models.School
	if err := s.db.Where("id = ?", id).First(&school).Error; err != nil {
		return nil, err
	}

	return &dto.School{
		Id:   school.ID,
		Name: school.Name,
		Code: school.Code,
	}, nil
}

func (s *School) GetByCode(code string) (*dto.School, error) {
	var school models.School
	if err := s.db.Where("code = ?", code).First(&school).Error; err != nil {
		return nil, err
	}

	return &dto.School{
		Id:   school.ID,
		Name: school.Name,
		Code: school.Code,
	}, nil
}
