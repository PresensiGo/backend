package responses

import "api/internal/dto"

type GetAllAttendances struct {
	Attendances []dto.Attendance `json:"attendances" validate:"required"`
} // @name GetAllAttendancesRes

type GetAttendanceItem struct {
	Student           dto.Student           `json:"student" validate:"required"`
	AttendanceStudent dto.AttendanceStudent `json:"attendanceStudent" validate:"required"`
} // @name GetAttendanceItemRes
type GetAttendance struct {
	Attendance dto.Attendance      `json:"attendance" validate:"required"`
	Items      []GetAttendanceItem `json:"items" validate:"required"`
} // @name GetAttendanceRes
