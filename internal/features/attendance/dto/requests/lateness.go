package requests

type CreateLateness struct {
	Date string `json:"date" validate:"required"`
} // @name CreateLatenessReq

type CreateLatenessDetail struct {
	StudentIds []uint `json:"student_ids" validate:"required"`
} // @name CreateLatenessDetailReq
