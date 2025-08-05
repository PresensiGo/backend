package domains

import (
	"api/internal/features/attendance/models"
)

type LatenessDetail struct {
	Id         uint `json:"id" validate:"required"`
	LatenessId uint `json:"lateness_id" validate:"required"`
	StudentId  uint `json:"student_id" validate:"required"`
}

func FromLatenessDetailModel(model *models.LatenessDetail) *LatenessDetail {
	return &LatenessDetail{
		Id:         model.ID,
		LatenessId: model.LatenessId,
		StudentId:  model.StudentId,
	}
}

func (l *LatenessDetail) ToModel() *models.LatenessDetail {
	return &models.LatenessDetail{
		LatenessId: l.LatenessId,
		StudentId:  l.StudentId,
	}
}
