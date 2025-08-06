package responses

import "api/internal/features/subject/domains"

type CreateSubject struct {
	Subject domains.Subject `json:"subject" validate:"required"`
}
