package requests

type Create struct {
	BatchId uint   `json:"batch_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
}
