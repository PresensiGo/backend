package responses

import "api/internal/dto"

type GetAllClassrooms struct {
	Classrooms []dto.Classroom `json:"classrooms" validate:"required"`
}
