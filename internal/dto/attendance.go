package dto

import (
	"time"
)

type Attendance struct {
	Id          uint      `json:"id" validate:"required"`
	ClassroomId uint      `json:"classroom_id" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}
