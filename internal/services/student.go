package services

import (
	"api/internal/dto/responses"
	"api/internal/repository"
)

type Student struct {
	student *repository.Student
}

func NewStudent(student *repository.Student) *Student {
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
