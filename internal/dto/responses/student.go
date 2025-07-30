package responses

import (
	"api/internal/dto"
	"api/internal/dto/combined"
)

type GetAllStudentsByClassroomId struct {
	Students []dto.Student `json:"students" validate:"required"`
} // @name GetAllStudentsByClassroomIdRes

type GetAllStudents struct {
	Students []combined.StudentMajorClassroom `json:"students" validate:"required"`
} // @name GetAllStudentsRes
