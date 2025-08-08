package requests

type CreateSubjectAttendance struct {
	SubjectId uint   `json:"subject_id" validate:"required"`
	DateTime  string `json:"datetime" validate:"required"`
	Note      string `json:"note" validate:"required"`
}
