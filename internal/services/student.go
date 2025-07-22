package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type Student struct {
	db *gorm.DB
}

func NewStudent(db *gorm.DB) *Student {
	return &Student{db}
}

func (s *Student) GetAllStudents(classId uint64) (*responses.GetAllStudents, error) {
	var students []models.Student
	if err := s.db.Where("class_id = ?", classId).
		Find(&students).
		Error; err != nil {
		return nil, err
	}

	var mappedStudents []dto.Student
	for _, student := range students {
		mappedStudents = append(
			mappedStudents,
			dto.Student{
				Id:   student.ID,
				Name: student.Name,
			},
		)
	}

	return &responses.GetAllStudents{
		Students: mappedStudents,
	}, nil
}
