package services

import (
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type Reset struct {
	db *gorm.DB
}

func NewReset(db *gorm.DB) *Reset {
	return &Reset{db}
}

func (s *Reset) Reset() (*responses.Reset, error) {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("true").Unscoped().Delete(&models.Student{}).Error; err != nil {
			return err
		}

		if err := tx.Where("true").Unscoped().Delete(&models.Classroom{}).Error; err != nil {
			return err
		}

		if err := tx.Where("true").Unscoped().Delete(&models.Major{}).Error; err != nil {
			return err
		}

		if err := tx.Where("true").Unscoped().Delete(&models.Batch{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &responses.Reset{
		Message: "Reset Success",
	}, nil
}
