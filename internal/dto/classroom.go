package dto

type Classroom struct {
	Id      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	MajorId uint   `json:"major_id" validate:"required"`
}
