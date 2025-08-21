package domains

import "api/internal/features/school/models"

type School struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
} // @name School

func FromSchoolModel(m *models.School) *School {
	return &School{
		Id:   m.ID,
		Name: m.Name,
		Code: m.Code,
	}
}

func (s *School) ToModel() *models.School {
	return &models.School{
		Name: s.Name,
		Code: s.Code,
	}
}
