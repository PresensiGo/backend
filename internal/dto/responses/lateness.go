package responses

import (
	"api/internal/dto"
	"api/internal/dto/combined"
)

type GetAllLatenesses struct {
	Latenesses []dto.Lateness `json:"latenesses" validate:"required"`
} // @name GetAllLatenessesRes

type GetLateness struct {
	Lateness dto.Lateness                     `json:"lateness" validate:"required"`
	Items    []combined.StudentMajorClassroom `json:"items" validate:"required"`
} // @name GetLatenessRes
