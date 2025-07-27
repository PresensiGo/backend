package repositories

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type Major struct {
	db *gorm.DB
}

func NewMajor(db *gorm.DB) *Major {
	return &Major{db}
}

func (r *Major) CreateInTx(tx *gorm.DB, data dto.Major) (*uint, error) {
	major := models.Major{
		Name:    data.Name,
		BatchId: data.BatchId,
	}
	if err := tx.Create(&major).Error; err != nil {
		return nil, err
	}

	return &major.ID, nil
}

func (r *Major) GetAllByBatchId(batchId uint) ([]dto.Major, error) {
	var majors []models.Major
	if err := r.db.Model(&models.Major{}).
		Where("batch_id = ?", batchId).
		Order("name asc").
		Find(&majors).Error; err != nil {
		return nil, err
	}

	var mappedMajors []dto.Major
	for _, major := range majors {
		mappedMajors = append(mappedMajors, dto.Major{
			ID:   major.ID,
			Name: major.Name,
		})
	}

	return mappedMajors, nil
}
