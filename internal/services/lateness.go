package services

import (
	"api/internal/dto"
	"api/internal/dto/requests"
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
