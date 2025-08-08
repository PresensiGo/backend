package services

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
	batch "api/internal/features/batch/repositories"
	classroom "api/internal/features/classroom/repositories"
	major "api/internal/features/major/repositories"
	"api/pkg/utils"
	"github.com/google/uuid"
)

type SubjectAttendance struct {
	batchRepo             *batch.Batch
	majorRepo             *major.Major
	classroomRepo         *classroom.Classroom
	subjectAttendanceRepo *repositories.SubjectAttendance
}

func NewSubjectAttendance(
	batchRepo *batch.Batch,
	majorRepo *major.Major,
	classroomRepo *classroom.Classroom,
	subjectAttendanceRepo *repositories.SubjectAttendance,
) *SubjectAttendance {
	return &SubjectAttendance{
		batchRepo:             batchRepo,
		majorRepo:             majorRepo,
		classroomRepo:         classroomRepo,
		subjectAttendanceRepo: subjectAttendanceRepo,
	}
}

func (s *SubjectAttendance) Create(classroomId uint, req requests.CreateSubjectAttendance) (
	*responses.CreateSubjectAttendance, error,
) {
	parsedDatetime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, err
	}

	subjectAttendance := domains.SubjectAttendance{
		DateTime:    *parsedDatetime,
		Code:        uuid.NewString(),
		Note:        req.Note,
		ClassroomId: classroomId,
		SubjectId:   req.SubjectId,
	}

	result, err := s.subjectAttendanceRepo.Create(subjectAttendance)
	if err != nil {
		return nil, err
	}

	return &responses.CreateSubjectAttendance{
		SubjectAttendance: *result,
	}, nil
}

func (s *SubjectAttendance) GetAll(schoolId uint) (*responses.GetAllSubjectAttendances, error) {
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

	majorIds := make([]uint, len(*majors))
	for i, v := range *majors {
		majorIds[i] = v.Id
	}

	classrooms, err := s.classroomRepo.GetManyByMajorId(majorIds)
	if err != nil {
		return nil, err
	}

	classroomIds := make([]uint, len(classrooms))
	for i, v := range classrooms {
		classroomIds[i] = v.Id
	}

	subjectAttendances, err := s.subjectAttendanceRepo.GetManyByClassroomIds(classroomIds)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllSubjectAttendances{
		SubjectAttendances: *subjectAttendances,
	}, nil
}
