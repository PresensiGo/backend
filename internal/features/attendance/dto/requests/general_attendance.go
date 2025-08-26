package requests

import "api/pkg/constants"

type CreateGeneralAttendance struct {
	DateTime string `json:"datetime"`
	Note     string `json:"note"`
} // @name CreateGeneralAttendanceReq

type CreateGeneralAttendanceRecord struct {
	StudentId uint                           `json:"student_id" validate:"required"`
	Status    constants.AttendanceStatusType `json:"status" validate:"required"`
} // @name CreateGeneralAttendanceRecordReq

type UpdateGeneralAttendance struct {
	DateTime string `json:"datetime"`
	Note     string `json:"note"`
}

type CreateGeneralAttendanceRecordStudent struct {
	Code string `json:"code" validate:"required"`
} // @name CreateGeneralAttendanceRecordStudentReq

type ExportGeneralAttendance struct {
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}
