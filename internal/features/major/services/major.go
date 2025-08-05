package services

import (
	"api/internal/features/major/domains"
	"api/internal/features/major/dto/responses"
	"api/internal/features/major/models"
	"gorm.io/gorm"
)

type Major struct {
	db *gorm.DB
}

func NewMajor(db *gorm.DB) *Major {
	return &Major{db}
}

func (s *Major) GetAllMajors(batchId uint64) (*responses.GetAllMajors, error) {
	var majors []models.Major
	if err := s.db.Where("batch_id = ?", batchId).
		Find(&majors).
		Error; err != nil {
		return nil, err
	}

	var mappedMajors []domains.Major
	for _, major := range majors {
		mappedMajors = append(
			mappedMajors, domains.Major{
				Id:   major.ID,
				Name: major.Name,
			},
		)
	}

	return &responses.GetAllMajors{
		Majors: mappedMajors,
	}, nil
}
