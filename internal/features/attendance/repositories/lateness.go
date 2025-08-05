package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
	"gorm.io/gorm"
)

type Lateness struct {
	db *gorm.DB
}

func NewLateness(db *gorm.DB) *Lateness {
	return &Lateness{db}
}

func (r *Lateness) Create(dto *domains.Lateness) (*uint, error) {
	lateness := dto.ToModel()
	if err := r.db.Create(&lateness).Error; err != nil {
		return nil, err
	}

	return &lateness.ID, nil
}

func (r *Lateness) GetAllBySchoolId(schoolId uint) (*[]domains.Lateness, error) {
	var latenesses []models.Lateness
	if err := r.db.Where("school_id = ?", schoolId).
		Order("date desc").
		Find(&latenesses).
		Error; err != nil {
		return nil, err
	}

	mappedLatenesses := make([]domains.Lateness, len(latenesses))
	for i, lateness := range latenesses {
		mappedLatenesses[i] = *domains.FromLatenessModel(&lateness)
	}

	return &mappedLatenesses, nil
}

func (r *Lateness) GetById(latenessId uint) (*domains.Lateness, error) {
	var lateness models.Lateness
	if err := r.db.Where("id = ?", latenessId).
		First(&lateness).
		Error; err != nil {
		return nil, err
	}

	return domains.FromLatenessModel(&lateness), nil
}

func (r *Lateness) DeleteBySchoolIdInTx(tx *gorm.DB, schoolId uint) error {
	return tx.Where("school_id = ?", schoolId).
		Unscoped().
		Delete(&models.Lateness{}).Error

}
