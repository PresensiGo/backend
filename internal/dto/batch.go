package dto

type Batch struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
