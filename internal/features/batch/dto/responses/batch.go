package responses

import (
	"api/internal/shared/domains"
)

type CreateBatch struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetAllBatches struct {
	Batches []domains.BatchInfo `json:"batches" validate:"required"`
}
