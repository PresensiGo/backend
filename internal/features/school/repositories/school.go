package repositories

import (
	"api/internal/features/school/domains"
	"api/internal/features/school/models"
	"gorm.io/gorm"
)

type School struct {
	db *gorm.DB
}

func NewSchool(db *gorm.DB) *School {
	return &School{db}
}

func (r *School) CreateInTx(tx *gorm.DB, data domains.School) (*domains.School, error) {
	school := data.ToModel()

	if err := tx.Create(&school).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSchoolModel(school), nil
	}
}

func (r *School) GetById(id uint) (*domains.School, error) {
	var school models.School
	if err := r.db.Where("id = ?", id).First(&school).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSchoolModel(&school), nil
	}
}

func (r *School) GetByCode(code string) (*domains.School, error) {
	var school models.School
	if err := r.db.Where("code = ?", code).First(&school).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSchoolModel(&school), nil
	}
}

func (r *School) GetByCodeInTx(tx *gorm.DB, code string) (*domains.School, error) {
	var school models.School
	if err := tx.Where("code = ?", code).First(&school).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSchoolModel(&school), nil
	}
}
