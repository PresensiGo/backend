package dto

import "api/internal/models"

type AttendanceDetail struct {
	Id           uint                    `json:"id" validate:"required"`
	AttendanceId uint                    `json:"attendance_id" validate:"required"`
	StudentId    uint                    `json:"student_id" validate:"required"`
	Status       models.AttendanceStatus `json:"status" validate:"required"`
	Note         string                  `json:"note" validate:"required"`
}
