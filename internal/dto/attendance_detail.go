package dto

import "api/internal/models"

type AttendanceDetail struct {
	ID           uint                    `json:"id" validate:"required"`
	AttendanceID uint                    `json:"attendance_id" validate:"required"`
	StudentID    uint                    `json:"student_id" validate:"required"`
	Status       models.AttendanceStatus `json:"status" validate:"required"`
	Note         string                  `json:"note" validate:"required"`
}
