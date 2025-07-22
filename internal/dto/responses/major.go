package responses

import "api/internal/dto"

type GetAllMajorsResponse struct {
	Majors []dto.Major `json:"majors"`
}
