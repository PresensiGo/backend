package services

import (
	"api/internal/dto"
	"api/internal/dto/requests"
	"api/internal/dto/responses"
	"api/internal/repositories"
)

type Lateness struct {
	latenessRepo *repositories.Lateness
}

func NewLateness(latenessRepo *repositories.Lateness) *Lateness {
	return &Lateness{latenessRepo}
}

func (s *Lateness) Create(
	schoolId uint,
	req *requests.CreateLateness,
) error {
	if _, err := s.latenessRepo.Create(&dto.Lateness{
		Date:     req.Date,
		SchoolId: schoolId,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Lateness) GetAllBySchoolId(schoolId uint) (*responses.GetAllLatenesses, error) {
	latenesses, err := s.latenessRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllLatenesses{
		Latenesses: *latenesses,
	}, nil
}
