package dto

type Class struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
