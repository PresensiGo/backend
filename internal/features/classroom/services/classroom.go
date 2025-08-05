package services

import (
	"api/internal/features/classroom/dto/responses"
	"api/internal/features/classroom/repositories"
	"api/internal/features/major/domains"
	repositories2 "api/internal/features/major/repositories"
)

type Classroom struct {
	classroomRepo *repositories.Classroom
	majorRepo     *repositories2.Major
}

func NewClassroom(
	classroomRepo *repositories.Classroom,
	majorRepo *repositories2.Major,
) *Classroom {
	return &Classroom{
		classroomRepo,
		majorRepo,
	}
}

func (s *Classroom) GetAllWithMajor(batchId uint) (*responses.GetAllClassroomWithMajors, error) {
	majors, err := s.majorRepo.GetAllByBatchId(batchId)
	if err != nil {
		return nil, err
	}

	majorMap := make(map[uint]domains.Major)
	for _, major := range majors {
		majorMap[major.Id] = major
	}

	var majorIds []uint
	for _, major := range majors {
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
