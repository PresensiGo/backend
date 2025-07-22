package responses

import "api/internal/dto"

type GetAllClasses struct {
	Classes []dto.Class `json:"classes"`
}
