package responses

import "api/internal/dto"

type ClassMajor struct {
	Class dto.Class `json:"class" validate:"required"`
	Major dto.Major `json:"major" validate:"required"`
}

type GetAllClassMajors struct {
	Data []ClassMajor `json:"data" validate:"required"`
}
