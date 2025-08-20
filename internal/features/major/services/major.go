package services

import (
	batchRepo "api/internal/features/batch/repositories"
	classroomRepo "api/internal/features/classroom/repositories"
	"api/internal/features/major/domains"
	"api/internal/features/major/dto"
	"api/internal/features/major/dto/requests"
	"api/internal/features/major/dto/responses"
	"api/internal/features/major/repositories"
	"api/pkg/http/failure"
	"gorm.io/gorm"
)

type Major struct {
	db            *gorm.DB
	batchRepo     *batchRepo.Batch
	majorRepo     *repositories.Major
	classroomRepo *classroomRepo.Classroom
}

func NewMajor(
	db *gorm.DB, batchRepo *batchRepo.Batch, majorRepo *repositories.Major,
	classroomRepo *classroomRepo.Classroom,
) *Major {
	return &Major{
		db:            db,
		batchRepo:     batchRepo,
		majorRepo:     majorRepo,
		classroomRepo: classroomRepo,
	}
}

func (s *Major) Create(batchId uint, req requests.CreateMajor) (*domains.Major, error) {
	major := domains.Major{
		Name:    req.Name,
		BatchId: batchId,
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

func (s *Major) GetAllMajorsByBatchId(batchId uint) (
	*responses.GetAllMajorsByBatchId, *failure.App,
) {
	if majors, err := s.majorRepo.GetAllByBatchId(batchId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		result := make([]dto.GetAllMajorsByBatchIdItem, len(*majors))

		for i, major := range *majors {
			if count, err := s.classroomRepo.GetCountByMajorId(major.Id); err != nil {
				return nil, failure.NewInternal(err)
			} else {
				result[i] = dto.GetAllMajorsByBatchIdItem{
					Major:          major,
					ClassroomCount: count,
				}
			}
		}

		return &responses.GetAllMajorsByBatchId{
			Items: result,
		}, nil
	}
}

func (s *Major) GetMajor(batchId uint) (
	*responses.GetMajor, *failure.App,
) {
	if major, err := s.majorRepo.Get(batchId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetMajor{
			Major: *major,
		}, nil
	}
}

func (s *Major) Update(majorId uint, req requests.UpdateMajor) (*domains.Major, error) {
	major := domains.Major{
		Name: req.Name,
	}

	return s.majorRepo.Update(majorId, major)
}

func (s *Major) Delete(majorId uint) error {
	return s.majorRepo.Delete(majorId)
}
