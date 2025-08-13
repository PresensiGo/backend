package domains

import (
	"api/internal/features/subject/models"
)

type Subject struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	SchoolId uint   `json:"school_id" validate:"required"`
} // @name Subject

func FromSubjectModel(m *models.Subject) *Subject {
	return &Subject{
		Id:       m.ID,
		Name:     m.Name,
		SchoolId: m.SchoolId,
	}
}

func (s *Subject) ToModel() *models.Subject {
	return &models.Subject{
		Name:     s.Name,
		SchoolId: s.SchoolId,
	}
}
