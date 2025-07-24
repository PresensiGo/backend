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
		var major dto.Major
		for _, _major := range majors {
			if _major.ID == class.MajorID {
				major = _major
				break
			}
		}

		result = append(result, responses.ClassMajor{
			Class: class,
			Major: major,
		})
	}

	return &responses.GetAllClassMajors{
		Data: result,
	}, nil
}
