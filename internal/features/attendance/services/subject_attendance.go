package services

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
	batch "api/internal/features/batch/repositories"
	classroom "api/internal/features/classroom/repositories"
	major "api/internal/features/major/repositories"
	subjectDomain "api/internal/features/subject/domains"
	subjectRepo "api/internal/features/subject/repositories"
	shared "api/internal/shared/domains"
	"api/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectAttendance struct {
	db                          *gorm.DB
	batchRepo                   *batch.Batch
	majorRepo                   *major.Major
	classroomRepo               *classroom.Classroom
	subjectAttendanceRepo       *repositories.SubjectAttendance
	subjectAttendanceRecordRepo *repositories.SubjectAttendanceRecord
	subjectRepo                 *subjectRepo.Subject
}

func NewSubjectAttendance(
	db *gorm.DB,
	batchRepo *batch.Batch,
	majorRepo *major.Major,
	classroomRepo *classroom.Classroom,
	subjectAttendanceRepo *repositories.SubjectAttendance,
	subjectAttendanceRecordRepo *repositories.SubjectAttendanceRecord,
	subjectRepo *subjectRepo.Subject,
) *SubjectAttendance {
	return &SubjectAttendance{
		db:                          db,
		batchRepo:                   batchRepo,
		majorRepo:                   majorRepo,
		classroomRepo:               classroomRepo,
		subjectAttendanceRepo:       subjectAttendanceRepo,
		subjectAttendanceRecordRepo: subjectAttendanceRecordRepo,
		subjectRepo:                 subjectRepo,
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

func (s *SubjectAttendance) CreateRecordStudent(
	studentId uint, req requests.CreateSubjectAttendanceRecordStudent,
) (
	*responses.CreateSubjectAttendanceRecordStudent, error,
) {
	subjectAttendance, err := s.subjectAttendanceRepo.GetByCode(req.Code)
	if err != nil {
		return nil, err
	}

	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			// hapus semua record yang sudah ada
			if err := s.subjectAttendanceRecordRepo.DeleteByAttendanceIdStudentIdInTx(
				tx, subjectAttendance.Id, studentId,
			); err != nil {
				return err
			}

			// buat record baru untuk student
			subjectAttendanceRecord := domains.SubjectAttendanceRecord{
				SubjectAttendanceId: subjectAttendance.Id,
				StudentId:           studentId,
			}
			if _, err := s.subjectAttendanceRecordRepo.CreateInTx(
				tx, subjectAttendanceRecord,
			); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return &responses.CreateSubjectAttendanceRecordStudent{
		Message: "ok",
	}, nil
}

func (s *SubjectAttendance) GetAll(classroomId uint) (*responses.GetAllSubjectAttendances, error) {
	subjectAttendances, err := s.subjectAttendanceRepo.GetAllByClassroomId(classroomId)
	if err != nil {
		return nil, err
	}

	subjectIds := make([]uint, len(*subjectAttendances))
	for i, v := range *subjectAttendances {
		subjectIds[i] = v.SubjectId
	}

	subject, err := s.subjectRepo.GetMany(subjectIds)
	if err != nil {
		return nil, err
	}

	mapSubject := make(map[uint]*subjectDomain.Subject)
	for _, v := range *subject {
		mapSubject[v.Id] = &v
	}

	result := make([]shared.SubjectAttendanceSubject, len(*subjectAttendances))
	for i, v := range *subjectAttendances {
		result[i] = shared.SubjectAttendanceSubject{
			SubjectAttendance: v,
			Subject:           *mapSubject[v.SubjectId],
		}
	}

	return &responses.GetAllSubjectAttendances{
		Items: result,
	}, nil
}

func (s *SubjectAttendance) Get(subjectAttendanceId uint) (*responses.GetSubjectAttendance, error) {
	result, err := s.subjectAttendanceRepo.Get(subjectAttendanceId)
	if err != nil {
		return nil, err
	}

	return &responses.GetSubjectAttendance{
		SubjectAttendance: *result,
	}, nil
}
