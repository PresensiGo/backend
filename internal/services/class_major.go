package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/repository"
)

type ClassMajor struct {
	classRepository *repository.Class
	majorRepository *repository.Major
}

func NewClassMajor(
	classRepository *repository.Class,
	majorRepository *repository.Major,
) *ClassMajor {
	return &ClassMajor{
		classRepository,
		majorRepository,
	}
}

func (s *ClassMajor) GetAll(batchId uint) (*responses.GetAllClassMajors, error) {
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

	var result []responses.ClassMajor
	for _, class := range classes {
		result = append(result, responses.ClassMajor{
			Class: class,
			Major: majorMap[class.ID],
		})
	}

	return &responses.GetAllClassMajors{
		Data: result,
	}, nil
}
