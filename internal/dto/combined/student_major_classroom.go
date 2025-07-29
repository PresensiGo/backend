package combined

import "api/internal/dto"

type StudentMajorClassroom struct {
	Student   dto.Student   `json:"student" validate:"required"`
	Major     dto.Major     `json:"major" validate:"required"`
	Classroom dto.Classroom `json:"classroom" validate:"required"`
}
