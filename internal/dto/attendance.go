package dto

import (
	"time"
)

type Attendance struct {
	ID          uint      `json:"id" validate:"required"`
	ClassroomID uint      `json:"classroom_id" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}
