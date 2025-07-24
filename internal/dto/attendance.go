package dto

import (
	"time"
)

type Attendance struct {
	ID      uint      `json:"id" validate:"required"`
	ClassID uint      `json:"class_id" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
}
