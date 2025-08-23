package services

import (
	"net/http"

	batchRepo "api/internal/features/batch/repositories"
	classroomRepo "api/internal/features/classroom/repositories"
	majorRepo "api/internal/features/major/repositories"
	schoolRepo "api/internal/features/school/repositories"
	domains4 "api/internal/features/student/domains"
	"api/internal/features/student/dto"
	"api/internal/features/student/dto/responses"
	"api/internal/features/student/repositories"
	"api/pkg/authentication"
	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
)

type Student struct {
	schoolRepo       *schoolRepo.School
	batchRepo        *batchRepo.Batch
	majorRepo        *majorRepo.Major
	classroomRepo    *classroomRepo.Classroom
	studentRepo      *repositories.Student
	studentTokenRepo *repositories.StudentToken
}

func NewStudent(
	schoolRepo *schoolRepo.School,
	batchRepo *batchRepo.Batch,
	majorRepo *majorRepo.Major,
	classroomRepo *classroomRepo.Classroom,
	studentRepo *repositories.Student,
	studentTokenRepo *repositories.StudentToken,
) *Student {
	return &Student{
		schoolRepo:       schoolRepo,
		batchRepo:        batchRepo,
		majorRepo:        majorRepo,
		classroomRepo:    classroomRepo,
		studentRepo:      studentRepo,
		studentTokenRepo: studentTokenRepo,
	}
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

func (s *Student) GetAllAccountsByClassroomId(classroomId uint) (
	*responses.GetAllStudentAccountsByClassroomId, error,
) {
	students, err := s.studentRepo.GetAllByClassroomId(classroomId)
	if err != nil {
		return nil, err
	}

	studentIds := make([]uint, len(students))
	for i, v := range students {
		studentIds[i] = v.Id
	}

	studentTokens, err := s.studentTokenRepo.GetManyByStudentIds(studentIds)
	if err != nil {
		return nil, err
	}

	mapStudentTokens := make(map[uint]*domains4.StudentToken)
	for _, v := range *studentTokens {
		mapStudentTokens[v.StudentId] = &v
	}

	result := make([]dto.StudentAccount, len(students))
	for i, v := range students {
		studentToken, ok := mapStudentTokens[v.Id]
		if !ok {
			studentToken = &domains4.StudentToken{}
		}

		result[i] = dto.StudentAccount{
			Student:      v,
			StudentToken: *studentToken,
		}
	}

	return &responses.GetAllStudentAccountsByClassroomId{
		Items: result,
	}, nil
}

// deprecated
// func (s *Student) GetAll(keyword string) (*responses.GetAllStudents, error) {
// 	students, err := s.studentRepo.GetAll(keyword)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	mapMajors := make(map[uint]*domains.Major)
// 	mapClassrooms := make(map[uint]*domains2.Classroom)
//
// 	classroomIds := make([]uint, len(*students))
// 	for i, v := range *students {
// 		classroomIds[i] = v.ClassroomId
// 	}
// 	classrooms, err := s.classroomRepo.GetManyByIds(classroomIds)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	majorIds := make([]uint, len(*classrooms))
// 	for i, v := range *classrooms {
// 		majorIds[i] = v.MajorId
// 		mapClassrooms[v.Id] = &v
// 	}
// 	majors, err := s.majorRepo.GetManyByIds(majorIds)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	for _, v := range *majors {
// 		mapMajors[v.Id] = &v
// 	}
//
// 	mappedStudents := make([]domains3.StudentMajorClassroom, len(*students))
// 	for i, v := range *students {
// 		classroom := mapClassrooms[v.ClassroomId]
// 		major := mapMajors[classroom.MajorId]
// 		mappedStudents[i] = domains3.StudentMajorClassroom{
// 			Student:   v,
// 			Major:     *major,
// 			Classroom: *classroom,
// 		}
// 	}
//
// 	return &responses.GetAllStudents{
// 		Students: mappedStudents,
// 	}, nil
// }

func (s *Student) GetProfileStudent(c *gin.Context) (*responses.GetProfileStudent, *failure.App) {
	auth := authentication.GetAuthenticatedStudent(c)
	if auth.Id == 0 {
		return nil, failure.NewApp(
			http.StatusForbidden,
			"Anda tidak memiliki akses!",
			nil,
		)
	}

	student, err := s.studentRepo.Get(auth.Id)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	school, err := s.schoolRepo.Get(student.SchoolId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	classroom, err := s.classroomRepo.Get(student.ClassroomId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	major, err := s.majorRepo.Get(classroom.MajorId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	batch, err := s.batchRepo.Get(major.BatchId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	return &responses.GetProfileStudent{
		Student:   *student,
		School:    *school,
		Classroom: *classroom,
		Major:     *major,
		Batch:     *batch,
	}, nil
}
