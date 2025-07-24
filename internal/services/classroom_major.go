package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/repository"
)

type ClassMajor struct {
	classRepository *repository.Classroom
	majorRepository *repository.Major
}

func NewClassMajor(
	classRepository *repository.Classroom,
	majorRepository *repository.Major,
) *ClassMajor {
	return &ClassMajor{
		classRepository,
		majorRepository,
	}
}

func (s *ClassMajor) GetAll(batchId uint) (*responses.GetAllClassroomMajors, error) {
	majors, err := s.majorRepository.GetAllByBatchId(batchId)
	if err != nil {
		return nil, err
	}

	majorMap := make(map[uint]dto.Major)
	for _, major := range majors {
		majorMap[major.ID] = major
	}

	var majorIds []uint
	for _, major := range majors {
		majorIds = append(majorIds, major.ID)
	}

	classes, err := s.classRepository.GetManyByMajorId(majorIds)
	if err != nil {
		return nil, err
	}

	result := make([]responses.ClassroomMajor, len(classes))
	for index, class := range classes {
		result[index] = responses.ClassroomMajor{
			Classroom: class,
			Major:     majorMap[class.ID],
		}
	}

	return &responses.GetAllClassroomMajors{
		Data: result,
	}, nil
}
