package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/repository"
)

type Classroom struct {
	classroomRepo *repository.Classroom
	majorRepo     *repository.Major
}

func NewClassroom(
	classroomRepo *repository.Classroom,
	majorRepo *repository.Major,
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

	majorMap := make(map[uint]dto.Major)
	for _, major := range majors {
		majorMap[major.ID] = major
	}

	var majorIds []uint
	for _, major := range majors {
		majorIds = append(majorIds, major.ID)
	}

	classes, err := s.classroomRepo.GetManyByMajorId(majorIds)
	if err != nil {
		return nil, err
	}

	result := make([]responses.ClassroomMajor, len(classes))
	for index, class := range classes {
		result[index] = responses.ClassroomMajor{
			Classroom: class,
			Major:     majorMap[class.MajorID],
		}
	}

	return &responses.GetAllClassroomWithMajors{
		Data: result,
	}, nil
}
