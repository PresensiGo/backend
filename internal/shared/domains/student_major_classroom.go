package domains

import (
	domains3 "api/internal/features/classroom/domains"
	"api/internal/features/major/domains"
	domains2 "api/internal/features/student/domains"
)

type StudentMajorClassroom struct {
	Student   domains2.Student   `json:"student" validate:"required"`
	Major     domains.Major      `json:"major" validate:"required"`
	Classroom domains3.Classroom `json:"classroom" validate:"required"`
}
