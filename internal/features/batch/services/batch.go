package services

import (
	"api/internal/features/batch/domains"
	"api/internal/features/batch/dto/requests"
	"api/internal/features/batch/dto/responses"
	repositories2 "api/internal/features/batch/repositories"
	"api/internal/features/classroom/repositories"
	domains2 "api/internal/features/major/domains"
	repositories3 "api/internal/features/major/repositories"
	domains3 "api/internal/shared/domains"
	"gorm.io/gorm"
)

type Batch struct {
	db            *gorm.DB
	batchRepo     *repositories2.Batch
	majorRepo     *repositories3.Major
	classroomRepo *repositories.Classroom
}

func NewBatch(
	db *gorm.DB, batchRepo *repositories2.Batch, majorRepo *repositories3.Major,
	classroomRepo *repositories.Classroom,
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

func (s *Batch) GetAllBySchoolId(schoolId uint) (*responses.GetAllBatches, error) {
	batches, err := s.batchRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	// map to store batch, majors, and classrooms
	mapBatches := make(map[uint]domains.Batch)
	mapMajors := make(map[uint]domains2.Major)

	mapBatchNameAndMajorName := make(map[string]map[string]bool) // batchName: { majorName: _ }

	// get majors
	batchIds := make([]uint, len(*batches))
	for i, v := range *batches {
		batchIds[i] = v.Id
		mapBatches[v.Id] = v
	}

	majors, err := s.majorRepo.GetManyByBatchIds(batchIds)
	if err != nil {
		return nil, err
	}

	majorIds := make([]uint, len(*majors))
	for i, v := range *majors {
		majorIds[i] = v.Id
		mapMajors[v.Id] = v

		batch := mapBatches[v.BatchId]
		if _, ok := mapBatchNameAndMajorName[batch.Name]; !ok {
			mapBatchNameAndMajorName[batch.Name] = make(map[string]bool)
		}

		if _, ok := mapBatchNameAndMajorName[batch.Name][v.Name]; !ok {
			mapBatchNameAndMajorName[batch.Name][v.Name] = true
		}
	}

	classrooms, err := s.classroomRepo.GetManyByMajorId(majorIds)
	if err != nil {
		return nil, err
	}

	mapBatchIdAndClassroomCount := make(map[uint]int)
	for _, v := range classrooms {
		major := mapMajors[v.MajorId]
		batch := mapBatches[major.BatchId]

		if count, ok := mapBatchIdAndClassroomCount[batch.Id]; !ok {
			mapBatchIdAndClassroomCount[batch.Id] = 1
		} else {
			mapBatchIdAndClassroomCount[batch.Id] = count + 1
		}
	}

	// mapping all data
	mappedBatches := make([]domains3.BatchInfo, len(*batches))
	for i, v := range *batches {
		mapMajorName := mapBatchNameAndMajorName[v.Name]
		classroomCount := mapBatchIdAndClassroomCount[v.Id]

		mappedBatches[i] = domains3.BatchInfo{
			Batch:           v,
			MajorsCount:     len(mapMajorName),
			ClassroomsCount: classroomCount,
		}
	}

	return &responses.GetAllBatches{
		Batches: mappedBatches,
	}, nil
}

func (s *Batch) Update(batchId uint, req requests.Update) (*domains.Batch, error) {
	batch := domains.Batch{
		Id:   batchId,
		Name: req.Name,
	}

	return s.batchRepo.Update(batch)
}
