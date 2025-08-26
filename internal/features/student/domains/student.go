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
	Gender      string `json:"gender" validate:"required"`
} // @name Student

func FromStudentModel(m *models.Student) *Student {
	return &Student{
		Id:          m.ID,
		NIS:         m.NIS,
		Name:        m.Name,
		SchoolId:    m.SchoolId,
		ClassroomId: m.ClassroomId,
		Gender:      m.Gender,
	}
}

func (s *Student) ToModel() *models.Student {
	return &models.Student{
		NIS:         s.NIS,
		Name:        s.Name,
		SchoolId:    s.SchoolId,
		ClassroomId: s.ClassroomId,
		Gender:      s.Gender,
	}
}
