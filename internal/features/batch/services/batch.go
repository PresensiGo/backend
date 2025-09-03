package services

import (
	"api/internal/features/batch/domains"
	"api/internal/features/batch/dto"
	"api/internal/features/batch/dto/requests"
	"api/internal/features/batch/dto/responses"
	"api/internal/features/batch/repositories"
	classroomRepo "api/internal/features/classroom/repositories"
	majorRepo "api/internal/features/major/repositories"
	"api/pkg/authentication"
	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
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

func (s *Batch) Create(schoolId uint, req requests.CreateBatch) (*domains.Batch, error) {
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

func (s *Batch) GetAllBatches(c *gin.Context) (*responses.GetAllBatches, *failure.App) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		return nil, failure.NewForbidden()
	}

	if batches, err := s.batchRepo.GetAllBySchoolId(user.SchoolId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		result := make([]dto.GetAllBatchesItem, len(*batches))

		for i, batch := range *batches {
			if count, err := s.majorRepo.GetCountByBatchId(batch.Id); err != nil {
				return nil, failure.NewInternal(err)
			} else {
				result[i] = dto.GetAllBatchesItem{
					Batch:      batch,
					MajorCount: count,
				}
			}
		}

		return &responses.GetAllBatches{
			Items: result,
		}, nil
	}
}

func (s *Batch) GetBatch(batchId uint) (*responses.GetBatch, *failure.App) {
	batch, err := s.batchRepo.Get(batchId)
	if err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetBatch{
			Batch: *batch,
		}, nil
	}
}

func (s *Batch) Update(batchId uint, req requests.UpdateBatch) (*domains.Batch, error) {
	batch := domains.Batch{
		Id:   batchId,
		Name: req.Name,
	}

	return s.batchRepo.Update(batch)
}

func (s *Batch) Delete(batchId uint) error {
	return s.batchRepo.Delete(batchId)
}
