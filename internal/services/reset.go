package services

import (
	"api/internal/dto/responses"
	"api/internal/models"
	"gorm.io/gorm"
)

type ResetService struct {
	db *gorm.DB
}

func NewResetService(db *gorm.DB) *ResetService {
	return &ResetService{db}
}

func (s *ResetService) Reset() (*responses.ResetResponse, error) {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("true").Unscoped().Delete(&models.Student{}).Error; err != nil {
			return err
		}

		if err := tx.Where("true").Unscoped().Delete(&models.Class{}).Error; err != nil {
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

	return &responses.ResetResponse{
		Message: "Reset Success",
	}, nil
}
