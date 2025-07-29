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

func (r *LatenessDetail) GetAllByLatenessId(latenessId uint) (*[]dto.LatenessDetail, error) {
	var latenessDetails []models.LatenessDetail
	if err := r.db.Where("lateness_id = ?", latenessId).
		Find(&latenessDetails).
		Error; err != nil {
		return nil, err
	}

	mappedLatenessDetails := make([]dto.LatenessDetail, len(latenessDetails))
	for i, item := range latenessDetails {
		mappedLatenessDetails[i] = *dto.FromLatenessDetailModel(&item)
	}

	return &mappedLatenessDetails, nil
}
