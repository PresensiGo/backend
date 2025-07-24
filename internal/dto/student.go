package dto

type Student struct {
	ID          uint   `json:"id" validate:"required"`
	NIS         string `json:"nis" validate:"required"`
	Name        string `json:"name" validate:"required"`
	ClassroomID uint   `json:"classroom_id" validate:"required"`
}
