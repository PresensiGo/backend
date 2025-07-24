package requests

import (
	"api/internal/models"
	"time"
)

type CreateAttendance struct {
	ClassID            uint      `json:"class_id"`
	Date               time.Time `json:"date"`
	AttendanceStudents []struct {
		StudentID uint                    `json:"student_id"`
		Status    models.AttendanceStatus `json:"status"`
		Note      string                  `json:"note"`
	} `json:"attendance_students"`
}
