package repositories

import (
	"api/internal/dto"
	"gorm.io/gorm"
)

type LatenessDetail struct {
	db *gorm.DB
}

func NewLatenessDetail(db *gorm.DB) *LatenessDetail {
	return &LatenessDetail{db}
}

func (r *LatenessDetail) Create(data *dto.LatenessDetail) (*uint, error) {
	latenessDetail := data.ToModel()
	if err := r.db.Create(&latenessDetail).Error; err != nil {
		return nil, err
	}

	return &latenessDetail.ID, nil
}
