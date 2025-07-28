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

func (s *Student) GetAllStudentsByClassroomId(classroomId uint) (
	*responses.GetAllStudentsByClassroomId, error,
) {
	students, err := s.student.GetAllByClassroomId(classroomId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllStudentsByClassroomId{
		Students: students,
	}, nil
}

func (s *Student) GetAll(keyword string) (*responses.GetAllStudents, error) {
	students, err := s.student.GetAll(keyword)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllStudents{
		Students: *students,
	}, nil
}
