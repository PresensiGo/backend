package repositories

import (
	"api/internal/features/school/domains"
	"api/internal/features/school/models"
	"gorm.io/gorm"
)

type School struct {
	db *gorm.DB
}

func NewSchool(db *gorm.DB) *School {
	return &School{db}
}

func (s *School) GetById(id uint) (*domains.School, error) {
	var school models.School
	if err := s.db.Where("id = ?", id).First(&school).Error; err != nil {
		return nil, err
	}

	return &domains.School{
		Id:   school.ID,
		Name: school.Name,
		Code: school.Code,
	}, nil
}

func (s *School) GetByCode(code string) (*domains.School, error) {
	var school models.School
	if err := s.db.Where("code = ?", code).First(&school).Error; err != nil {
		return nil, err
	}

	return &domains.School{
		Id:   school.ID,
		Name: school.Name,
		Code: school.Code,
	}, nil
}
