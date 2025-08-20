package services

import (
	batchRepo "api/internal/features/batch/repositories"
	"api/internal/features/classroom/domains"
	"api/internal/features/classroom/dto"
	"api/internal/features/classroom/dto/requests"
	"api/internal/features/classroom/dto/responses"
	"api/internal/features/classroom/repositories"
	majorDomain "api/internal/features/major/domains"
	majorRepo "api/internal/features/major/repositories"
	studentRepo "api/internal/features/student/repositories"
	"api/pkg/http/failure"
)

type Classroom struct {
	batchRepo     *batchRepo.Batch
	majorRepo     *majorRepo.Major
	classroomRepo *repositories.Classroom
	studentRepo   *studentRepo.Student
}

func NewClassroom(
	batchRepo *batchRepo.Batch,
	majorRepo *majorRepo.Major,
	classroomRepo *repositories.Classroom,
	studentRepo *studentRepo.Student,
) *Classroom {
	return &Classroom{
		batchRepo:     batchRepo,
		majorRepo:     majorRepo,
		classroomRepo: classroomRepo,
		studentRepo:   studentRepo,
	}
}

func (s *Classroom) Create(majorId uint, req requests.CreateClassroom) (
	*responses.CreateClassroom, error,
) {
	classroom := domains.Classroom{
		Name:    req.Name,
		MajorId: majorId,
	}

	result, err := s.classroomRepo.Create(classroom)
	if err != nil {
		return nil, err
	}

	return &responses.CreateClassroom{
		Classroom: *result,
	}, nil
}

func (s *Classroom) GetAll(schoolId uint) (*responses.GetAll, error) {
	batches, err := s.batchRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	batchIds := make([]uint, len(*batches))
	for i, v := range *batches {
		batchIds[i] = v.Id
	}

	majors, err := s.majorRepo.GetManyByBatchIds(batchIds)
	if err != nil {
		return nil, err
	}

	majorIds := make([]uint, len(*majors))
	for i, v := range *majors {
		majorIds[i] = v.Id
	}

	classrooms, err := s.classroomRepo.GetManyByMajorId(majorIds)
	if err != nil {
		return nil, err
	}

	return &responses.GetAll{
		Classrooms: classrooms,
	}, nil
}

func (s *Classroom) GetAllClassroomsByMajorId(majorId uint) (
	*responses.GetAllClassroomsByMajorId, *failure.App,
) {
	if classrooms, err := s.classroomRepo.GetAllByMajorId(majorId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		result := make([]dto.GetAllClassroomsByMajorIdItem, len(*classrooms))

		for i, classroom := range *classrooms {
			if count, err := s.studentRepo.GetCountByClassroomId(classroom.Id); err != nil {
				return nil, failure.NewInternal(err)
			} else {
				result[i] = dto.GetAllClassroomsByMajorIdItem{
					Classroom:    classroom,
					StudentCount: count,
				}
			}
		}

		return &responses.GetAllClassroomsByMajorId{
			Items: result,
		}, nil
	}
}

func (s *Classroom) GetAllWithMajor(batchId uint) (*responses.GetAllClassroomWithMajors, error) {
	majors, err := s.majorRepo.GetAllByBatchId(batchId)
	if err != nil {
		return nil, err
	}

	majorMap := make(map[uint]majorDomain.Major)
	for _, major := range *majors {
		majorMap[major.Id] = major
	}

	var majorIds []uint
	for _, major := range *majors {
		majorIds = append(majorIds, major.Id)
	}

	classes, err := s.classroomRepo.GetManyByMajorId(majorIds)
	if err != nil {
		return nil, err
	}

	result := make([]responses.ClassroomMajor, len(classes))
	for index, class := range classes {
		result[index] = responses.ClassroomMajor{
			Classroom: class,
			Major:     majorMap[class.MajorId],
		}
	}

	return &responses.GetAllClassroomWithMajors{
		Data: result,
	}, nil
}

func (s *Classroom) GetClassroom(classroomId uint) (
	*responses.GetClassroom, *failure.App,
) {
	if classroom, err := s.classroomRepo.Get(classroomId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetClassroom{
			Classroom: *classroom,
		}, nil
	}
}

func (s *Classroom) Update(
	classroomId uint, req requests.UpdateClassroom,
) (*responses.UpdateClassroom, error) {
	classroom := domains.Classroom{
		Name: req.Name,
	}

	result, err := s.classroomRepo.Update(classroomId, classroom)
	if err != nil {
		return nil, err
	}

	return &responses.UpdateClassroom{
		Classroom: *result,
	}, nil
}
