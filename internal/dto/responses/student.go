package responses

import "api/internal/dto"

type GetAllStudentsByClassroomId struct {
	Students []dto.Student `json:"students" validate:"required"`
} // @name GetAllStudentsByClassroomIdRes

type GetAllStudents struct {
	Students []dto.Student `json:"students" validate:"required"`
} // @name GetAllStudentsRes
