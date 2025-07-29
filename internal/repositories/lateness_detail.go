package repositories

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type LatenessDetail struct {
	db *gorm.DB
}

func NewLatenessDetail(db *gorm.DB) *LatenessDetail {
	return &LatenessDetail{db}
}

func (r *LatenessDetail) CreateBatch(data *[]dto.LatenessDetail) error {
	latenessDetails := make([]models.LatenessDetail, len(*data))
	for i, item := range *data {
		latenessDetails[i] = *item.ToModel()
	}

	return r.db.Create(&latenessDetails).Error
}
