package dto

type Student struct {
	Id   uint   `json:"id" validate:"required"`
	NIS  string `json:"nis" validate:"required"`
	Name string `json:"name" validate:"required"`
}
