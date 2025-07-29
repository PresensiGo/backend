package repositories

import (
	"fmt"
	"strings"

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

func (r *Student) GetAll(keyword string) (*[]dto.Student, error) {
	keyword = fmt.Sprintf("%%%s%%", strings.ToLower(keyword))

	var students []models.Student
	if err := r.db.
		Where(
			"lower(name) like ? or lower(nis) like ?",
			keyword, keyword,
		).
		Order("name asc").
		Order("nis asc").
		Find(&students).
		Error; err != nil {
		return nil, err
	}

	mappedStudents := make([]dto.Student, len(students))
	for i, student := range students {
		mappedStudents[i] = *dto.FromStudentModel(&student)
	}

	return &mappedStudents, nil
}

func (r *Student) GetAllByClassroomId(classroomId uint) ([]dto.Student, error) {
	var students []models.Student
	if err := r.db.Model(&models.Student{}).
		Where("classroom_id = ?", classroomId).
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
	for i, item := range students {
		mappedStudents[i] = *dto.FromStudentModel(&item)
	}

	return &mappedStudents, nil
}
