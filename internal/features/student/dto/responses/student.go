package responses

import (
	"api/internal/features/student/domains"
	"api/internal/features/student/dto"
	domains2 "api/internal/shared/domains"
)

type GetAllStudentsByClassroomId struct {
	Students []domains.Student `json:"students" validate:"required"`
} // @name GetAllStudentsByClassroomIdRes

type GetAllStudentAccountsByClassroomId struct {
	Items []dto.StudentAccount `json:"items" validate:"required"`
}

type GetAllStudents struct {
	Students []domains2.StudentMajorClassroom `json:"students" validate:"required"`
} // @name GetAllStudentsRes
