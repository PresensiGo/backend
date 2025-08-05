package responses

import (
	"api/internal/features/student/domains"
	domains2 "api/internal/shared/domains"
)

type GetAllStudentsByClassroomId struct {
	Students []domains.Student `json:"students" validate:"required"`
} // @name GetAllStudentsByClassroomIdRes

type GetAllStudents struct {
	Students []domains2.StudentMajorClassroom `json:"students" validate:"required"`
} // @name GetAllStudentsRes
