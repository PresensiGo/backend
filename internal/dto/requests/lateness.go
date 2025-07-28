package requests

type CreateLateness struct {
	Date string `json:"date" validate:"required"`
} // @name CreateLatenessReq

type CreateLatenessDetail struct {
	StudentId uint `json:"student_id" validate:"required"`
} // @name CreateLatenessDetailReq
