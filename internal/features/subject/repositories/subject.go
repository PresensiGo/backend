package repositories

import (
	"api/internal/features/subject/domains"
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
