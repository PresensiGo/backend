package dto

import (
	"api/internal/features/batch/domains"
)

type BatchInfo struct {
	Batch           domains.Batch `json:"batch" validate:"required"`
	MajorsCount     int           `json:"majors_count" validate:"required"`
	ClassroomsCount int           `json:"classrooms_count" validate:"required"`
}
