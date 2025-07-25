package requests

import (
	"api/internal/models"
)

type CreateAttendanceItem struct {
	StudentID uint                    `json:"student_id"`
	Status    models.AttendanceStatus `json:"status"`
	Note      string                  `json:"note"`
} // @name CreateAttendanceItemReq

type CreateAttendance struct {
	ClassroomID        uint                   `json:"classroom_id"`
	Date               string                 `json:"date"`
	AttendanceStudents []CreateAttendanceItem `json:"attendance_students"`
} // @name CreateAttendanceReq
