package responses

import "api/internal/dto"

type GetAllLatenesses struct {
	Latenesses []dto.Lateness `json:"latenesses" validate:"required"`
} // @name GetAllLatenessesRes
