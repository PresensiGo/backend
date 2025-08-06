package responses

import "api/internal/features/attendance/domains"

type CreateGeneralAttendance struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
}

type GetAllGeneralAttendances struct {
	GeneralAttendances []domains.GeneralAttendance `json:"general_attendances" validate:"required"`
}

type GetGeneralAttendance struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
}

type UpdateGeneralAttendance struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
}

type DeleteGeneralAttendance struct {
	Message string `json:"message"`
}
