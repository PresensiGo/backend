package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
	"gorm.io/gorm"
)

type LatenessDetail struct {
	db *gorm.DB
}

func NewLatenessDetail(db *gorm.DB) *LatenessDetail {
	return &LatenessDetail{db}
}

func (r *LatenessDetail) CreateBatch(data *[]domains.LatenessDetail) error {
	latenessDetails := make([]models.LatenessDetail, len(*data))
	for i, item := range *data {
		latenessDetails[i] = *item.ToModel()
	}

	return r.db.Create(&latenessDetails).Error
}

func (r *LatenessDetail) GetAllByLatenessId(latenessId uint) (*[]domains.LatenessDetail, error) {
	var latenessDetails []models.LatenessDetail
	if err := r.db.Where("lateness_id = ?", latenessId).
		Find(&latenessDetails).
		Error; err != nil {
		return nil, err
	}

	mappedLatenessDetails := make([]domains.LatenessDetail, len(latenessDetails))
	for i, item := range latenessDetails {
		mappedLatenessDetails[i] = *domains.FromLatenessDetailModel(&item)
	}

	return &mappedLatenessDetails, nil
}
