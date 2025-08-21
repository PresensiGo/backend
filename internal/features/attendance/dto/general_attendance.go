package dto

import (
	"api/internal/features/attendance/domains"
	student "api/internal/features/student/domains"
	user "api/internal/features/user/domains"
)

type GetAllGeneralAttendanceRecordsItem struct {
	Student student.Student                 `json:"student" validate:"required"`
	Record  domains.GeneralAttendanceRecord `json:"record" validate:"required"`
} // @name GetAllGeneralAttendanceRecordsItem

type GetAllGeneralAttendanceRecordsByClassroomIdItem struct {
	Student student.Student                 `json:"student" validate:"required"`
	Record  domains.GeneralAttendanceRecord `json:"record" validate:"required"`
} // @name GetAllGeneralAttendanceRecordsByClassroomIdItem

type GetAllGeneralAttendancesItem struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
	Creator           user.User                 `json:"creator" validate:"required"`
} // @name GetAllGeneralAttendancesItem
