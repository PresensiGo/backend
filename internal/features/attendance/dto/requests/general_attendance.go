package requests

type CreateGeneralAttendance struct {
	DateTime string `json:"datetime"`
	Note     string `json:"note"`
}
