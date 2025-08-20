package responses

import (
	"api/internal/features/classroom/domains"
	"api/internal/features/classroom/dto"
	majorDomain "api/internal/features/major/domains"
)

type ClassroomMajor struct {
	Classroom domains.Classroom `json:"classroom" validate:"required"`
	Major     majorDomain.Major `json:"major" validate:"required"`
}

type CreateClassroom struct {
	Classroom domains.Classroom `json:"classroom" validate:"required"`
}

type GetAll struct {
	Classrooms []domains.Classroom `json:"classrooms" validate:"required"`
}

type GetAllClassroomsByMajorId struct {
	Items []dto.GetAllClassroomsByMajorIdItem `json:"items" validate:"required"`
} // @name GetAllClassroomsByMajorIdRes

type GetAllClassroomWithMajors struct {
	Data []ClassroomMajor `json:"data" validate:"required"`
}

type GetClassroom struct {
	Classroom domains.Classroom `json:"classroom" validate:"required"`
} // @name GetClassroomRes

type UpdateClassroom struct {
	Classroom domains.Classroom `json:"classroom" validate:"required"`
}
