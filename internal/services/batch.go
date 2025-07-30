package services

import (
	"api/internal/dto"
	"api/internal/dto/combined"
	"api/internal/dto/responses"
	"api/internal/models"
	"api/internal/repositories"
	"gorm.io/gorm"
)

type Batch struct {
	db            *gorm.DB
	batchRepo     *repositories.Batch
	majorRepo     *repositories.Major
	classroomRepo *repositories.Classroom
}

func NewBatch(
	db *gorm.DB, batchRepo *repositories.Batch, majorRepo *repositories.Major,
	classroomRepo *repositories.Classroom,
) *Batch {
	return &Batch{db, batchRepo, majorRepo, classroomRepo}
}

func (s *Batch) Create(name string) (*responses.CreateBatch, error) {
	batch := models.Batch{
		Name: name,
	}

	if err := s.db.Create(&batch).Error; err != nil {
		return nil, err
	}

	return &responses.CreateBatch{
		Id:   batch.ID,
		Name: batch.Name,
	}, nil
}

func (s *Batch) GetAllBySchoolId(schoolId uint) (*responses.GetAllBatches, error) {
	batches, err := s.batchRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	// map to store batch, majors, and classrooms
	mapBatches := make(map[uint]dto.Batch)
	mapMajors := make(map[uint]dto.Major)

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
	mappedBatches := make([]combined.BatchInfo, len(*batches))
	for i, v := range *batches {
		mapMajorName := mapBatchNameAndMajorName[v.Name]
		classroomCount := mapBatchIdAndClassroomCount[v.Id]

		mappedBatches[i] = combined.BatchInfo{
			Batch:           v,
			MajorsCount:     len(mapMajorName),
			ClassroomsCount: classroomCount,
		}
	}

	return &responses.GetAllBatches{
		Batches: mappedBatches,
	}, nil
}
