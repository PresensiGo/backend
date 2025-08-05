package services

import (
	"api/internal/features/attendance/repositories"
	repositories2 "api/internal/features/batch/repositories"
	"gorm.io/gorm"
)

type Reset struct {
	db           *gorm.DB
	batchRepo    *repositories2.Batch
	latenessRepo *repositories.Lateness
}

func NewReset(
	db *gorm.DB, batchRepo *repositories2.Batch, latenessRepo *repositories.Lateness,
) *Reset {
	return &Reset{db, batchRepo, latenessRepo}
}

func (s *Reset) ResetBySchoolId(schoolId uint) error {
	return s.db.Transaction(
		func(tx *gorm.DB) error {
			// delete batch with its relation
			if err := s.batchRepo.DeleteBySchoolIdInTx(tx, schoolId); err != nil {
				return err
			}

			// delete lateness with its relation
			if err := s.latenessRepo.DeleteBySchoolIdInTx(tx, schoolId); err != nil {
				return err
			}

			return nil
		},
	)
}
