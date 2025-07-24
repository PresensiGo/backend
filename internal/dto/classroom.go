package dto

type Classroom struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	MajorID uint   `json:"major_id" validate:"required"`
}
