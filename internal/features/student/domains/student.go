package domains

import (
	"api/internal/features/student/models"
)

type Student struct {
	Id          uint   `json:"id" validate:"required"`
	NIS         string `json:"nis" validate:"required"`
	Name        string `json:"name" validate:"required"`
	SchoolId    uint   `json:"school_id" validate:"required"`
	ClassroomId uint   `json:"classroom_id" validate:"required"`
} // @name Student

func FromStudentModel(model *models.Student) *Student {
	return &Student{
		Id:          model.ID,
		NIS:         model.NIS,
		Name:        model.Name,
		SchoolId:    model.SchoolId,
		ClassroomId: model.ClassroomId,
	}
}

func (s *Student) ToModel() *models.Student {
	return &models.Student{
		NIS:         s.NIS,
		Name:        s.Name,
		SchoolId:    s.SchoolId,
		ClassroomId: s.ClassroomId,
	}
}
