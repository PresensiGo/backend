package repositories

import (
	"api/internal/features/subject/domains"
	"api/internal/features/subject/models"
	"gorm.io/gorm"
)

type Subject struct {
	db *gorm.DB
}

func NewSubject(db *gorm.DB) *Subject {
	return &Subject{db: db}
}

func (r *Subject) Create(data domains.Subject) (*domains.Subject, error) {
	subject := data.ToModel()
	if err := r.db.Create(&subject).Error; err != nil {
		return nil, err
	}

	return domains.FromSubjectModel(subject), nil
}

func (r *Subject) GetAll(schoolId uint) (*[]domains.Subject, error) {
	var subjects []models.Subject
	if err := r.db.Where(
		"school_id = ?", schoolId,
	).Order("name asc").Find(&subjects).Error; err != nil {
		return nil, err
	}

	result := make([]domains.Subject, len(subjects))
	for i, v := range subjects {
		result[i] = *domains.FromSubjectModel(&v)
	}

	return &result, nil
}

func (r *Subject) GetMany(subjectIds []uint) (*[]domains.Subject, error) {
	var subjects []models.Subject
	if err := r.db.Where("id in ?", subjectIds).Find(&subjects).Error; err != nil {
		return nil, err
	}

	result := make([]domains.Subject, len(subjects))
	for i, v := range subjects {
		result[i] = *domains.FromSubjectModel(&v)
	}

	return &result, nil
}

func (r *Subject) Get(subjectId uint) (*domains.Subject, error) {
	var subject models.Subject
	if err := r.db.Where("id = ?", subjectId).First(&subject).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSubjectModel(&subject), nil
	}
}

func (r *Subject) Update(subjectId uint, data domains.Subject) (*domains.Subject, error) {
	subject := data.ToModel()
	if err := r.db.Model(&subject).Where("id = ?", subjectId).Updates(&subject).Error; err != nil {
		return nil, err
	}

	return domains.FromSubjectModel(subject), nil
}

func (r *Subject) Delete(subjectId uint) error {
	return r.db.Where("id = ?", subjectId).Unscoped().Delete(&domains.Subject{}).Error
}
