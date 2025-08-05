package responses

import (
	"api/internal/features/major/domains"
)

type GetAllMajors struct {
	Majors []domains.Major `json:"majors" validate:"required"`
}
