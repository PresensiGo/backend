package requests

import "time"

type CreateLateness struct {
	Date time.Time `json:"date" validate:"required"`
} // @name CreateLatenessReq
