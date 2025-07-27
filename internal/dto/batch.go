package dto

type Batch struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	SchoolId uint   `json:"school_id" validate:"required"`
}
