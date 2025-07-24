package repository

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type Student struct {
	db *gorm.DB
}

func NewStudent(db *gorm.DB) *Student {
	return &Student{db}
}

func (r *Student) GetAllByClassId(classId uint) ([]dto.Student, error) {
	var students []models.Student
	if err := r.db.Model(&models.Student{}).
		Where("class_id = ?", classId).
		Order("name asc").
		Find(&students).Error; err != nil {
		return nil, err
	}

	var mappedStudents []dto.Student
	for _, student := range students {
		mappedStudents = append(mappedStudents, dto.Student{
			ID:      student.ID,
			NIS:     student.NIS,
			Name:    student.Name,
			ClassID: student.ClassID,
		})
	}

	return mappedStudents, nil
}
