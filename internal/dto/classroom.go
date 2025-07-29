package dto

import (
	"api/internal/models"
)

type Classroom struct {
	Id      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	MajorId uint   `json:"major_id" validate:"required"`
}

func FromClassroomModel(model *models.Classroom) *Classroom {
	return &Classroom{
		Id:      model.ID,
		Name:    model.Name,
		MajorId: model.MajorId,
	}
}

func (c *Classroom) ToModel() *models.Classroom {
	return &models.Classroom{
		Name:    c.Name,
		MajorId: c.MajorId,
	}
}
