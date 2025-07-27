package repositories

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type Classroom struct {
	db *gorm.DB
}

func NewClassroom(db *gorm.DB) *Classroom {
	return &Classroom{db}
}

func (r *Classroom) CreateInTx(tx *gorm.DB, data dto.Classroom) (*uint, error) {
	classroom := models.Classroom{
		Name:    data.Name,
		MajorId: data.MajorId,
	}
	if err := tx.Create(&classroom).Error; err != nil {
		return nil, err
	}

	return &classroom.ID, nil
}

func (r *Classroom) GetManyByMajorId(majorIds []uint) ([]dto.Classroom, error) {
	var classes []models.Classroom
	if err := r.db.Where("major_id in ?", majorIds).
		Order("name asc").
		Find(&classes).
		Error; err != nil {
		return nil, err
	}

	var mappedClasses []dto.Classroom
	for _, class := range classes {
		mappedClasses = append(mappedClasses, dto.Classroom{
			Id:      class.ID,
			Name:    class.Name,
			MajorId: class.MajorId,
		})
	}

	return mappedClasses, nil
}
