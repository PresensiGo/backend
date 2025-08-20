package responses

import (
	"api/internal/features/batch/domains"
	"api/internal/features/batch/dto"
)

type GetAllBatches struct {
	Items []dto.GetAllBatchesItem `json:"items" validate:"required"`
} // @name GetAllBatchesRes

type GetBatch struct {
	Batch domains.Batch `json:"batch" validate:"required"`
}
