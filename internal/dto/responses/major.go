package responses

import "api/internal/dto"

type GetAllMajors struct {
	Majors []dto.Major `json:"majors"`
}
