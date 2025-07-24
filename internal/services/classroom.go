package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type Classroom struct {
	db *gorm.DB
}

func NewClassroom(db *gorm.DB) *Classroom {
	return &Classroom{db}
}

func (s *Classroom) GetAllClassrooms(majorId uint64) (*responses.GetAllClassrooms, error) {
	var classes []models.Classroom
	if err := s.db.Where("major_id = ?", majorId).
		Find(&classes).Error; err != nil {
		return nil, err
	}

	var mappedClasses []dto.Classroom
	for _, class := range classes {
		mappedClasses = append(
			mappedClasses,
			dto.Classroom{
				ID:   class.ID,
				Name: class.Name,
			},
		)
	}

	return &responses.GetAllClassrooms{
		Classrooms: mappedClasses,
	}, nil
}
