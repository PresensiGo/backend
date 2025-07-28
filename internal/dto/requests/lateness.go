package requests

type CreateLateness struct {
	Date string `json:"date" validate:"required"`
} // @name CreateLatenessReq
