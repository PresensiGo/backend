package responses

import (
	"api/internal/dto/combined"
)

type CreateBatch struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetAllBatches struct {
	Batches []combined.BatchInfo `json:"batches" validate:"required"`
}
