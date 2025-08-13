package services

import (
	"api/internal/features/batch/domains"
	"api/internal/features/batch/dto/requests"
	"api/internal/features/batch/dto/responses"
	"api/internal/features/batch/repositories"
	classroomRepo "api/internal/features/classroom/repositories"
	majorRepo "api/internal/features/major/repositories"
	"api/pkg/http/failure"
	"gorm.io/gorm"
)

type Batch struct {
	db            *gorm.DB
	batchRepo     *repositories.Batch
	majorRepo     *majorRepo.Major
	classroomRepo *classroomRepo.Classroom
}

func NewBatch(
	db *gorm.DB, batchRepo *repositories.Batch, majorRepo *majorRepo.Major,
	classroomRepo *classroomRepo.Classroom,
) *Batch {
	return &Batch{db, batchRepo, majorRepo, classroomRepo}
}

func (s *Batch) Create(schoolId uint, req requests.Create) (*domains.Batch, error) {
	batch := domains.Batch{
		Name:     req.Name,
		SchoolId: schoolId,
	}

	result, err := s.batchRepo.Create(batch)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Batch) GetAllBatches(schoolId uint) (*responses.GetAllBatches, *failure.App) {
	if batches, err := s.batchRepo.GetAllBySchoolId(schoolId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetAllBatches{
			Batches: *batches,
		}, nil
	}
}

func (s *Batch) Get(batchId uint) (*responses.GetBatch, error) {
	batch, err := s.batchRepo.Get(batchId)
	if err != nil {
		return nil, err
	}

	return &responses.GetBatch{
		Batch: *batch,
	}, nil
}

func (s *Batch) Update(batchId uint, req requests.Update) (*domains.Batch, error) {
	batch := domains.Batch{
		Id:   batchId,
		Name: req.Name,
	}

	return s.batchRepo.Update(batch)
}

func (s *Batch) Delete(batchId uint) error {
	return s.batchRepo.Delete(batchId)
}
