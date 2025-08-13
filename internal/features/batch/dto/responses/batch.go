package responses

import (
	"api/internal/features/batch/domains"
)

type GetAllBatches struct {
	Batches []domains.Batch `json:"batches" validate:"required"`
}

type GetBatch struct {
	Batch domains.Batch `json:"batch" validate:"required"`
}
