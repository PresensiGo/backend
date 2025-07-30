package services

import (
	"api/internal/dto"
	"api/internal/dto/combined"
	"api/internal/dto/responses"
	"api/internal/repositories"
)

type Student struct {
	majorRepo     *repositories.Major
	classroomRepo *repositories.Classroom
	studentRepo   *repositories.Student
}

func NewStudent(
	majorRepo *repositories.Major, classroomRepo *repositories.Classroom,
	studentRepo *repositories.Student,
) *Student {
	return &Student{majorRepo, classroomRepo, studentRepo}
}

func (s *Student) GetAllStudentsByClassroomId(classroomId uint) (
	*responses.GetAllStudentsByClassroomId, error,
) {
	students, err := s.studentRepo.GetAllByClassroomId(classroomId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllStudentsByClassroomId{
		Students: students,
	}, nil
}

func (s *Student) GetAll(keyword string) (*responses.GetAllStudents, error) {
	students, err := s.studentRepo.GetAll(keyword)
	if err != nil {
		return nil, err
	}

	mapMajors := make(map[uint]*dto.Major)
	mapClassrooms := make(map[uint]*dto.Classroom)

	classroomIds := make([]uint, len(*students))
	for i, v := range *students {
		classroomIds[i] = v.ClassroomId
	}
	classrooms, err := s.classroomRepo.GetManyByIds(classroomIds)
	if err != nil {
		return nil, err
	}

	majorIds := make([]uint, len(*classrooms))
	for i, v := range *classrooms {
		majorIds[i] = v.MajorId
		mapClassrooms[v.Id] = &v
	}
	majors, err := s.majorRepo.GetManyByIds(majorIds)
	if err != nil {
		return nil, err
	}

	for _, v := range *majors {
		mapMajors[v.Id] = &v
	}

	mappedStudents := make([]combined.StudentMajorClassroom, len(*students))
	for i, v := range *students {
		classroom := mapClassrooms[v.ClassroomId]
		major := mapMajors[classroom.MajorId]
		mappedStudents[i] = combined.StudentMajorClassroom{
			Student:   v,
			Major:     *major,
			Classroom: *classroom,
		}
	}

	return &responses.GetAllStudents{
		Students: mappedStudents,
	}, nil
}
