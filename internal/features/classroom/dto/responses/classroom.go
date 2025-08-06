package responses

import (
	"api/internal/features/classroom/domains"
	majorDomain "api/internal/features/major/domains"
)

type ClassroomMajor struct {
	Classroom domains.Classroom `json:"classroom" validate:"required"`
	Major     majorDomain.Major `json:"major" validate:"required"`
}

type GetAll struct {
	Classrooms []domains.Classroom `json:"classrooms" validate:"required"`
}

type GetAllClassroomsByMajorId struct {
	Classrooms []domains.Classroom `json:"classrooms" validate:"required"`
}

type GetAllClassroomWithMajors struct {
	Data []ClassroomMajor `json:"data" validate:"required"`
}
