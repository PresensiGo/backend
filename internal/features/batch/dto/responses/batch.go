package responses

import (
	"api/internal/features/batch/domains"
	"api/internal/features/batch/dto"
)

type GetAllBatches struct {
	Batches []dto.BatchInfo `json:"batches" validate:"required"`
}

type GetBatch struct {
	Batch domains.Batch `json:"batch" validate:"required"`
}
