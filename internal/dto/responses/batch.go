package responses

import "api/internal/dto"

type CreateBatch struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetAllBatches struct {
	Batches []dto.Batch `json:"batches"`
}
