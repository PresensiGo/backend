package repository

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type Class struct {
	db *gorm.DB
}

func NewClass(db *gorm.DB) *Class {
	return &Class{db}
}

func (r *Class) GetManyByMajorId(majorIds []uint) ([]dto.Class, error) {
	var classes []models.Class
	if err := r.db.Where("major_id in ?", majorIds).
		Order("name asc").
		Find(&classes).
		Error; err != nil {
		return nil, err
	}

	var mappedClasses []dto.Class
	for _, class := range classes {
		mappedClasses = append(mappedClasses, dto.Class{
			ID:      class.ID,
			Name:    class.Name,
			MajorID: class.MajorId,
		})
	}

	return mappedClasses, nil
}
