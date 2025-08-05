package repositories

import (
	"api/internal/features/major/domains"
	"api/internal/features/major/models"
	"gorm.io/gorm"
)

type Major struct {
	db *gorm.DB
}

func NewMajor(db *gorm.DB) *Major {
	return &Major{db}
}

func (r *Major) CreateInTx(tx *gorm.DB, data domains.Major) (*uint, error) {
	major := models.Major{
		Name:    data.Name,
		BatchId: data.BatchId,
	}
	if err := tx.Create(&major).Error; err != nil {
		return nil, err
	}

	return &major.ID, nil
}

func (r *Major) GetAllByBatchId(batchId uint) ([]domains.Major, error) {
	var majors []models.Major
	if err := r.db.Model(&models.Major{}).
		Where("batch_id = ?", batchId).
		Order("name asc").
		Find(&majors).Error; err != nil {
		return nil, err
	}

	var mappedMajors []domains.Major
	for _, major := range majors {
		mappedMajors = append(
			mappedMajors, domains.Major{
				Id:   major.ID,
				Name: major.Name,
			},
		)
	}

	return mappedMajors, nil
}

func (r *Major) GetManyByIds(majorIds []uint) (*[]domains.Major, error) {
	var majors []models.Major
	if err := r.db.Model(&models.Major{}).
		Where("id in ?", majorIds).
		Find(&majors).
		Error; err != nil {
		return nil, err
	}

	mappedMajors := make([]domains.Major, len(majors))
	for i, item := range majors {
		mappedMajors[i] = *domains.FromMajorModel(&item)
	}

	return &mappedMajors, nil
}

func (r *Major) GetManyByBatchIds(batchIds []uint) (*[]domains.Major, error) {
	var majors []models.Major
	if err := r.db.Model(&models.Major{}).
		Where("batch_id in ?", batchIds).
		Find(&majors).
		Error; err != nil {
		return nil, err
	}

	mappedMajors := make([]domains.Major, len(majors))
	for i, item := range majors {
		mappedMajors[i] = *domains.FromMajorModel(&item)
	}

	return &mappedMajors, nil
}
