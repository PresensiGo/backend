package responses

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto"
	user "api/internal/features/user/domains"
)

type CreateGeneralAttendance struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
}

type CreateGeneralAttendanceRecordStudent struct {
	Message string `json:"message" validate:"required"`
}

type GetAllGeneralAttendances struct {
	GeneralAttendances []domains.GeneralAttendance `json:"general_attendances" validate:"required"`
}

type GetGeneralAttendance struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
	Creator           user.User                 `json:"creator" validate:"required"`
} // @name GetGeneralAttendanceRes

type GetAllGeneralAttendanceRecords struct {
	Items []dto.GetAllGeneralAttendanceRecordsItem `json:"items" validate:"required"`
}

type GetAllGeneralAttendanceRecordsByClassroomId struct {
	Items []dto.GetAllGeneralAttendanceRecordsByClassroomIdItem `json:"items" validate:"required"`
} // @name GetAllGeneralAttendanceRecordsByClassroomIdRes

type UpdateGeneralAttendance struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
}

type DeleteGeneralAttendance struct {
	Message string `json:"message"`
}
