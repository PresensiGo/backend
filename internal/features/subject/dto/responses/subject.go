package responses

import "api/internal/features/subject/domains"

type CreateSubject struct {
	Subject domains.Subject `json:"subject" validate:"required"`
}

type GetAllSubjects struct {
	Subjects []domains.Subject `json:"subjects" validate:"required"`
}

type GetSubject struct {
	Subject domains.Subject `json:"subject" validate:"required"`
} // @name GetSubjectRes

type UpdateSubject struct {
	Subject domains.Subject `json:"subject" validate:"required"`
}

type DeleteSubject struct {
	Message string `json:"message" validate:"required"`
}
