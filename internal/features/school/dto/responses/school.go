package responses

import "api/internal/features/school/domains"

type GetSchool struct {
	School domains.School `json:"school" validate:"required"`
} // @name GetSchoolRes
