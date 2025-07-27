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

func (r *Batch) DeleteBySchoolId(schoolId uint) error {
	return r.db.Where("school_id = ?", schoolId).
		Unscoped().
		Delete(&models.Batch{}).
		Error
}
