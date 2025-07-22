package responses

import "api/internal/dto"

type CreateBatchResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetAllBatchesResponse struct {
	Batches []dto.Batch `json:"batches"`
}
