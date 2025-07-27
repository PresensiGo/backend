package repositories

import (
	"api/internal/models"
	"gorm.io/gorm"
)

type Batch struct {
	db *gorm.DB
}

func NewBatch(db *gorm.DB) *Batch {
	return &Batch{db}
}

func (r *Batch) DeleteBySchoolId(schoolId uint) error {
	return r.db.Where("school_id = ?", schoolId).
		Unscoped().
		Delete(&models.Batch{}).
		Error
}
