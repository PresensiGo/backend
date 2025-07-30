package repositories

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type Batch struct {
	db *gorm.DB
}

func NewBatch(db *gorm.DB) *Batch {
	return &Batch{db}
}

func (r *Batch) CreateInTx(tx *gorm.DB, data dto.Batch) (*uint, error) {
	batch := &models.Batch{
		Name:     data.Name,
		SchoolId: data.SchoolId,
	}

	if err := tx.Create(&batch).Error; err != nil {
		return nil, err
	}

	return &batch.ID, nil
}

func (r *Batch) GetAllBySchoolId(schoolId uint) (*[]dto.Batch, error) {
	var batches []models.Batch
	if err := r.db.Model(&models.Batch{}).
		Where("school_id = ?", schoolId).
		Order("name asc").
		Find(&batches).
		Error; err != nil {
		return nil, err
	}

	mappedBatches := make([]dto.Batch, len(batches))
	for i, v := range batches {
		mappedBatches[i] = *dto.FromBatchModel(&v)
	}

	return &mappedBatches, nil
}

func (r *Batch) DeleteBySchoolIdInTx(tx *gorm.DB, schoolId uint) error {
	return tx.Where("school_id = ?", schoolId).
		Unscoped().
		Delete(&models.Batch{}).
		Error
}
