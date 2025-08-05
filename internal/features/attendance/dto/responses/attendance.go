package responses

import (
	"api/internal/features/attendance/domains"
	domains2 "api/internal/features/student/domains"
)

type GetAllAttendances struct {
	Attendances []domains.Attendance `json:"attendances" validate:"required"`
} // @name GetAllAttendancesRes

type GetAttendanceItem struct {
	Student           domains2.Student         `json:"student" validate:"required"`
	AttendanceStudent domains.AttendanceDetail `json:"attendanceStudent" validate:"required"`
} // @name GetAttendanceItemRes
type GetAttendance struct {
	Attendance domains.Attendance  `json:"attendance" validate:"required"`
	Items      []GetAttendanceItem `json:"items" validate:"required"`
} // @name GetAttendanceRes
