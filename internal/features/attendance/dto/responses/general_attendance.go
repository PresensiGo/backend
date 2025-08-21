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

type CreateGeneralAttendanceRecord struct {
	GeneralAttendanceRecord domains.GeneralAttendanceRecord `json:"general_attendance_record" validate:"required"`
}

type GetAllGeneralAttendances struct {
	Items []dto.GetAllGeneralAttendancesItem `json:"items" validate:"required"`
} // @name GetAllGeneralAttendancesRes

type GetAllGeneralAttendancesStudent struct {
	Items []dto.GetAllGeneralAttendancesStudentItem `json:"items" validate:"required"`
} // @name GetAllGeneralAttendancesStudentRes

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

type DeleteGeneralAttendanceRecord struct {
	Message string `json:"message"`
}
