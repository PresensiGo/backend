package responses

import "api/internal/dto"

type GetAllClassesResponse struct {
	Classes []dto.Class `json:"classes"`
}
