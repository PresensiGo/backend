package combined

import "api/internal/dto"

type BatchInfo struct {
	Batch           dto.Batch `json:"batch" validate:"required"`
	MajorsCount     int       `json:"majors_count" validate:"required"`
	ClassroomsCount int       `json:"classrooms_count" validate:"required"`
}
