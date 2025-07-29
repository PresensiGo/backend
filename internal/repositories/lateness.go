package repositories

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type Lateness struct {
	db *gorm.DB
}

func NewLateness(db *gorm.DB) *Lateness {
	return &Lateness{db}
}

func (r *Lateness) Create(dto *dto.Lateness) (*uint, error) {
	lateness := dto.ToModel()
	if err := r.db.Create(&lateness).Error; err != nil {
		return nil, err
	}

	return &lateness.ID, nil
}

func (r *Lateness) GetAllBySchoolId(schoolId uint) (*[]dto.Lateness, error) {
	var latenesses []models.Lateness
	if err := r.db.Where("school_id = ?", schoolId).
		Order("date desc").
		Find(&latenesses).
		Error; err != nil {
		return nil, err
	}

	mappedLatenesses := make([]dto.Lateness, len(latenesses))
	for i, lateness := range latenesses {
		mappedLatenesses[i] = *dto.FromLatenessModel(&lateness)
	}

	return &mappedLatenesses, nil
}

func (r *Lateness) GetById(latenessId uint) (*dto.Lateness, error) {
	var lateness models.Lateness
	if err := r.db.Where("id = ?", latenessId).
		First(&lateness).
		Error; err != nil {
		return nil, err
	}

	return dto.FromLatenessModel(&lateness), nil
}
