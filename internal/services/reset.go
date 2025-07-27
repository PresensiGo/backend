package services

import (
	"api/internal/repositories"
)

type Reset struct {
	batchRepo *repositories.Batch
}

func NewReset(batchRepo *repositories.Batch) *Reset {
	return &Reset{batchRepo}
}

func (s *Reset) ResetBySchoolId(schoolId uint) error {
	return s.batchRepo.DeleteBySchoolId(schoolId)
}
