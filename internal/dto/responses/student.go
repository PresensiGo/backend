package responses

import "api/internal/dto"

type GetAllStudents struct {
	Students []dto.Student `json:"students" validate:"required"`
}
