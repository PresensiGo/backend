package repositories

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

func (r *Student) CreateBatchInTx(tx *gorm.DB, data []dto.Student) error {
	students := make([]models.Student, len(data))
	for i, student := range data {
		students[i] = models.Student{
			NIS:         student.NIS,
			Name:        student.Name,
			ClassroomId: student.ClassroomId,
		}
	}

	return tx.Create(&students).Error
}

func (r *Student) GetAllByClassId(classId uint) ([]dto.Student, error) {
	var students []models.Student
	if err := r.db.Model(&models.Student{}).
		Where("classroom_id = ?", classId).
		Order("name asc").
		Find(&students).Error; err != nil {
		return nil, err
	}

	mappedStudents := make([]dto.Student, len(students))
	for index, student := range students {
		mappedStudents[index] = dto.Student{
			Id:          student.ID,
			NIS:         student.NIS,
			Name:        student.Name,
			ClassroomId: student.ClassroomId,
		}
	}

	return mappedStudents, nil
}

func (r *Student) GetManyById(studentIds []uint) (*[]dto.Student, error) {
	var students []models.Student
	if err := r.db.Model(&models.Student{}).
		Where("id in ?", studentIds).
		Order("name asc").
		Find(&students).Error; err != nil {
		return nil, err
	}

	mappedStudents := make([]dto.Student, len(students))
	for index, student := range students {
		mappedStudents[index] = dto.Student{
			Id:          student.ID,
			NIS:         student.NIS,
			Name:        student.Name,
			ClassroomId: student.ClassroomId,
		}
	}

	return &mappedStudents, nil
}
