package services

import (
	"api/internal/dto/responses"
	"api/internal/repositories"
)

type Student struct {
	student *repositories.Student
}

func NewStudent(student *repositories.Student) *Student {
	return &Student{student}
}

func (s *Student) GetAllStudents(classId uint) (*responses.GetAllStudents, error) {
	students, err := s.student.GetAllByClassId(classId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllStudents{
		Students: students,
	}, nil
}
