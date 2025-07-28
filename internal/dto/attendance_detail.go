package dto

import (
	"api/internal/models"
)

type AttendanceDetail struct {
	Id           uint                    `json:"id" validate:"required"`
	AttendanceId uint                    `json:"attendance_id" validate:"required"`
	StudentId    uint                    `json:"student_id" validate:"required"`
	Status       models.AttendanceStatus `json:"status" validate:"required"`
	Note         string                  `json:"note" validate:"required"`
}

func FromAttendanceDetailModel(model *models.AttendanceDetail) *AttendanceDetail {
	return &AttendanceDetail{
		Id:           model.ID,
		AttendanceId: model.AttendanceId,
		StudentId:    model.StudentId,
		Status:       model.Status,
		Note:         model.Note,
	}
}

func (a *AttendanceDetail) ToModel() *models.AttendanceDetail {
	return &models.AttendanceDetail{
		AttendanceId: a.AttendanceId,
		StudentId:    a.StudentId,
		Status:       a.Status,
		Note:         a.Note,
	}
}
