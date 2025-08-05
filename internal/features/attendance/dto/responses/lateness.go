package responses

import (
	"api/internal/features/attendance/domains"
	domains2 "api/internal/shared/domains"
)

type GetAllLatenesses struct {
	Latenesses []domains.Lateness `json:"latenesses" validate:"required"`
} // @name GetAllLatenessesRes

type GetLateness struct {
	Lateness domains.Lateness                 `json:"lateness" validate:"required"`
	Items    []domains2.StudentMajorClassroom `json:"items" validate:"required"`
} // @name GetLatenessRes
