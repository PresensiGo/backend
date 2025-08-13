package requests

type CreateGeneralAttendance struct {
	DateTime string `json:"datetime"`
	Note     string `json:"note"`
} // @name CreateGeneralAttendanceReq

type UpdateGeneralAttendance struct {
	DateTime string `json:"datetime"`
	Note     string `json:"note"`
}

type CreateGeneralAttendanceRecordStudent struct {
	Code string `json:"code" validate:"required"`
} // @name CreateGeneralAttendanceRecordStudentReq
