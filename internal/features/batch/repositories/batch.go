package repositories

import (
	"api/internal/features/batch/domains"
	"api/internal/features/batch/models"
	"gorm.io/gorm"
)

type Batch struct {
	db *gorm.DB
}

func NewBatch(db *gorm.DB) *Batch {
	return &Batch{db}
}

func (r *Batch) GetOrCreateInTx(tx *gorm.DB, data domains.Batch) (*domains.Batch, error) {
	var batch models.Batch
	if err := tx.FirstOrCreate(&batch, data.ToModel()).Error; err != nil {
		return nil, err
	} else {
		return domains.FromBatchModel(&batch), nil
	}
}

func (r *Batch) Create(domain domains.Batch) (*domains.Batch, error) {
	batch := domain.ToModel()
	if err := r.db.Create(&batch).Error; err != nil {
		return nil, err
	}

	return domains.FromBatchModel(batch), nil
}

func (r *Batch) CreateInTx(tx *gorm.DB, data domains.Batch) (*uint, error) {
	batch := &models.Batch{
		Name:     data.Name,
		SchoolId: data.SchoolId,
	}

	if err := tx.Create(&batch).Error; err != nil {
		return nil, err
	}

	return &batch.ID, nil
}

func (r *Batch) CreateInTx2(tx *gorm.DB, data domains.Batch) (*domains.Batch, error) {
	batch := data.ToModel()
	if err := tx.Create(&batch).Error; err != nil {
		return nil, err
	} else {
		return domains.FromBatchModel(batch), nil
	}
}

func (r *Batch) GetAllBySchoolId(schoolId uint) (*[]domains.Batch, error) {
	var batches []models.Batch
	if err := r.db.Model(&models.Batch{}).Where(
		"school_id = ?", schoolId,
	).Order("name asc").Find(&batches).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.Batch, len(batches))
		for i, batch := range batches {
			result[i] = *domains.FromBatchModel(&batch)
		}

		return &result, nil
	}
}

func (r *Batch) Get(batchId uint) (*domains.Batch, error) {
	var batch models.Batch
	if err := r.db.Model(&models.Batch{}).
		Where("id = ?", batchId).
		First(&batch).
		Error; err != nil {
		return nil, err
	}

	return domains.FromBatchModel(&batch), nil
}

func (r *Batch) GetBySchoolIdNameInTx(tx *gorm.DB, schoolId uint, name string) (
	*domains.Batch, error,
) {
	var batch models.Batch
	if err := tx.Model(&models.Batch{}).Where(
		"school_id = ? AND name = ?", schoolId, name,
	).First(&batch).Error; err != nil {
		return nil, err
	} else {
		return domains.FromBatchModel(&batch), nil
	}
}

func (r *Batch) Update(domain domains.Batch) (*domains.Batch, error) {
	model := domain.ToModel()
	if err := r.db.Model(&model).Where("id = ?", domain.Id).Updates(&model).Error; err != nil {
		return nil, err
	}

	return domains.FromBatchModel(model), nil
}

func (r *Batch) Delete(batchId uint) error {
	return r.db.Where("id = ?", batchId).Unscoped().Delete(&models.Batch{}).Error
}

func (r *Batch) DeleteBySchoolIdInTx(tx *gorm.DB, schoolId uint) error {
	return tx.Where("school_id = ?", schoolId).
		Unscoped().
		Delete(&models.Batch{}).
		Error
}
