package services

import (
	batchRepo "api/internal/features/batch/repositories"
	"api/internal/features/major/domains"
	"api/internal/features/major/dto/requests"
	"api/internal/features/major/dto/responses"
	"api/internal/features/major/models"
	"api/internal/features/major/repositories"
	"gorm.io/gorm"
)

type Major struct {
	db        *gorm.DB
	batchRepo *batchRepo.Batch
	majorRepo *repositories.Major
}

func NewMajor(db *gorm.DB, batchRepo *batchRepo.Batch, majorRepo *repositories.Major) *Major {
	return &Major{db, batchRepo, majorRepo}
}

func (s *Major) Create(req requests.Create) (*domains.Major, error) {
	major := domains.Major{
		Name:    req.Name,
		BatchId: req.BatchId,
	}

	return s.majorRepo.Create(major)
}

func (s *Major) GetAllMajors(schoolId uint) (*[]domains.Major, error) {
	batches, err := s.batchRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	batchIds := make([]uint, len(*batches))
	for i, v := range *batches {
		batchIds[i] = v.Id
	}

	majors, err := s.majorRepo.GetManyByBatchIds(batchIds)
	if err != nil {
		return nil, err
	}

	return majors, nil
}

func (s *Major) GetAllMajorsByBatchId(batchId uint64) (*responses.GetAllMajorsByBatchId, error) {
	var majors []models.Major
	if err := s.db.Where("batch_id = ?", batchId).
		Find(&majors).
		Error; err != nil {
		return nil, err
	}

	var mappedMajors []domains.Major
	for _, major := range majors {
		mappedMajors = append(
			mappedMajors, domains.Major{
				Id:   major.ID,
				Name: major.Name,
			},
		)
	}

	return &responses.GetAllMajorsByBatchId{
		Majors: mappedMajors,
	}, nil
}

func (s *Major) Update(majorId uint, req requests.Update) (*domains.Major, error) {
	major := domains.Major{
		Name:    req.Name,
		BatchId: req.BatchId,
	}

	return s.majorRepo.Update(majorId, major)
}

func (s *Major) Delete(majorId uint) error {
	return s.majorRepo.Delete(majorId)
}
