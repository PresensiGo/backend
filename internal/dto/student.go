package dto

type Student struct {
	ID      uint   `json:"id" validate:"required"`
	NIS     string `json:"nis" validate:"required"`
	Name    string `json:"name" validate:"required"`
	ClassID uint   `json:"class_id" validate:"required"`
}
