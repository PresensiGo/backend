package dto

type Major struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
