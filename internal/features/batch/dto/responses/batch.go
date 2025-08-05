package responses

import (
	"api/internal/shared/domains"
)

type GetAllBatches struct {
	Batches []domains.BatchInfo `json:"batches" validate:"required"`
}
