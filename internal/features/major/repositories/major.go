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

func (r *Major) GetOrCreateInTx(tx *gorm.DB, data domains.Major) (*domains.Major, error) {
	var major models.Major
	if err := tx.FirstOrCreate(&major, data.ToModel()).Error; err != nil {
		return nil, err
	} else {
		return domains.FromMajorModel(&major), nil
	}
}

func (r *Major) Create(data domains.Major) (*domains.Major, error) {
	major := data.ToModel()
	if err := r.db.Create(&major).Error; err != nil {
		return nil, err
	}

	return domains.FromMajorModel(major), nil
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

func (r *Major) CreateInTx2(tx *gorm.DB, data domains.Major) (*domains.Major, error) {
	major := data.ToModel()
	if err := tx.Create(&major).Error; err != nil {
		return nil, err
	} else {
		return domains.FromMajorModel(major), nil
	}
}

func (r *Major) GetAllByBatchId(batchId uint) (*[]domains.Major, error) {
	var majors []models.Major
	if err := r.db.Model(&models.Major{}).
		Where("batch_id = ?", batchId).
		Order("name asc").
		Find(&majors).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.Major, len(majors))
		for i, major := range majors {
			result[i] = *domains.FromMajorModel(&major)
		}

		return &result, nil
	}
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

func (r *Major) GetByBatchIdNameInTx(
	tx *gorm.DB, batchId uint, name string,
) (*domains.Major, error) {
	var major models.Major
	if err := tx.Model(&models.Major{}).Where(
		"batch_id = ? AND name = ?", batchId, name,
	).First(&major).Error; err != nil {
		return nil, err
	} else {
		return domains.FromMajorModel(&major), nil
	}
}

func (r *Major) GetCountByBatchId(batchId uint) (int64, error) {
	var count int64
	if err := r.db.Model(&models.Major{}).
		Where("batch_id = ?", batchId).
		Count(&count).Error; err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func (r *Major) Get(majorId uint) (*domains.Major, error) {
	var major models.Major
	if err := r.db.Model(&models.Major{}).Where(
		"id = ?", majorId,
	).First(&major).Error; err != nil {
		return nil, err
	} else {
		return domains.FromMajorModel(&major), nil
	}
}

func (r *Major) Update(majorId uint, data domains.Major) (*domains.Major, error) {
	major := data.ToModel()
	if err := r.db.Model(&models.Major{}).Where(
		"id = ?", majorId,
	).Updates(&major).Error; err != nil {
		return nil, err
	}

	return domains.FromMajorModel(major), nil
}

func (r *Major) Delete(majorId uint) error {
	return r.db.Where("id = ?", majorId).Unscoped().Delete(&models.Major{}).Error
}
