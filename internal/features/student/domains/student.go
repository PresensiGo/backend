package domains

import (
	"api/internal/features/student/models"
)

type Student struct {
	Id          uint   `json:"id" validate:"required"`
	NIS         string `json:"nis" validate:"required"`
	Name        string `json:"name" validate:"required"`
	ClassroomId uint   `json:"classroom_id" validate:"required"`
}

func FromStudentModel(model *models.Student) *Student {
	return &Student{
		Id:          model.ID,
		NIS:         model.NIS,
		Name:        model.Name,
		ClassroomId: model.ClassroomId,
	}
}

func (s *Student) ToModel() *models.Student {
	return &models.Student{
		NIS:         s.NIS,
		Name:        s.Name,
		ClassroomId: s.ClassroomId,
	}
}
