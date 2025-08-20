package requests

import "api/pkg/constants"

type CreateSubjectAttendance struct {
	SubjectId uint   `json:"subject_id" validate:"required"`
	DateTime  string `json:"datetime" validate:"required"`
	Note      string `json:"note" validate:"required"`
} // @name CreateSubjectAttendanceReq

type CreateSubjectAttendanceRecord struct {
	StudentId uint                       `json:"student_id" validate:"required"`
	DateTime  string                     `json:"datetime" validate:"required"`
	Status    constants.AttendanceStatus `json:"status" validate:"required"`
} // @name CreateSubjectAttendanceRecordReq

type CreateSubjectAttendanceRecordStudent struct {
	Code string `json:"code" validate:"required"`
} // @name CreateSubjectAttendanceRecordStudentReq
