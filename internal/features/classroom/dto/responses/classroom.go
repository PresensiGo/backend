package responses

import (
	domains2 "api/internal/features/classroom/domains"
	"api/internal/features/major/domains"
)

type ClassroomMajor struct {
	Classroom domains2.Classroom `json:"classroom" validate:"required"`
	Major     domains.Major      `json:"major" validate:"required"`
}

type GetAllClassroomWithMajors struct {
	Data []ClassroomMajor `json:"data" validate:"required"`
}
