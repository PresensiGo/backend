package responses

import "api/internal/dto"

type ClassroomMajor struct {
	Classroom dto.Classroom `json:"classroom" validate:"required"`
	Major     dto.Major     `json:"major" validate:"required"`
}

type GetAllClassroomWithMajors struct {
	Data []ClassroomMajor `json:"data" validate:"required"`
}
