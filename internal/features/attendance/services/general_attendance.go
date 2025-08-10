package services

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
	"api/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneralAttendance struct {
	db                          *gorm.DB
	generalAttendanceRepo       *repositories.GeneralAttendance
	generalAttendanceRecordRepo *repositories.GeneralAttendanceRecord
}

func NewGeneralAttendance(
	db *gorm.DB,
	generalAttendanceRepo *repositories.GeneralAttendance,
	generalAttendanceRecordRepo *repositories.GeneralAttendanceRecord,
) *GeneralAttendance {
	return &GeneralAttendance{
		db:                          db,
		generalAttendanceRepo:       generalAttendanceRepo,
		generalAttendanceRecordRepo: generalAttendanceRecordRepo,
	}
}

func (s *GeneralAttendance) Create(
	schoolId uint, req requests.CreateGeneralAttendance,
) (*responses.CreateGeneralAttendance, error) {
	parsedDateTime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, err
	}

	generalAttendance := domains.GeneralAttendance{
		DateTime: *parsedDateTime,
		Note:     req.Note,
		SchoolId: schoolId,
		Code:     uuid.NewString(),
	}

	result, err := s.generalAttendanceRepo.Create(generalAttendance)
	if err != nil {
		return nil, err
	}

	return &responses.CreateGeneralAttendance{
		GeneralAttendance: *result,
	}, nil
}

func (s *GeneralAttendance) CreateGeneralAttendanceRecordStudent(
	studentId uint, req requests.CreateGeneralAttendanceRecordStudent,
) (*responses.CreateGeneralAttendanceRecordStudent, error) {
	generalAttendance, err := s.generalAttendanceRepo.GetByCode(req.Code)
	if err != nil {
		return nil, err
	}

	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			// hapus record lama jika sudah ada
			if err := s.generalAttendanceRecordRepo.DeleteByAttendanceIdStudentIdInTx(
				tx, generalAttendance.Id, studentId,
			); err != nil {
				return err
			}

			// create record baru
			generalAttendanceRecord := domains.GeneralAttendanceRecord{
				GeneralAttendanceId: generalAttendance.Id,
				StudentId:           studentId,
			}
			if _, err := s.generalAttendanceRecordRepo.CreateInTx(
				tx, generalAttendanceRecord,
			); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return &responses.CreateGeneralAttendanceRecordStudent{
		Message: "ok",
	}, nil
}

func (s *GeneralAttendance) GetAll(schoolId uint) (*responses.GetAllGeneralAttendances, error) {
	result, err := s.generalAttendanceRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllGeneralAttendances{
		GeneralAttendances: *result,
	}, nil
}

func (s *GeneralAttendance) Get(generalAttendanceId uint) (
	*responses.GetGeneralAttendance, error,
) {
	result, err := s.generalAttendanceRepo.Get(generalAttendanceId)
	if err != nil {
		return nil, err
	}

	return &responses.GetGeneralAttendance{
		GeneralAttendance: *result,
	}, nil
}

func (s *GeneralAttendance) Update(
	generalAttendanceId uint, req requests.UpdateGeneralAttendance,
) (*responses.UpdateGeneralAttendance, error) {
	parsedDateTime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, err
	}

	generalAttendance := domains.GeneralAttendance{
		DateTime: *parsedDateTime,
		Note:     req.Note,
	}

	result, err := s.generalAttendanceRepo.Update(generalAttendanceId, generalAttendance)
	if err != nil {
		return nil, err
	}

	return &responses.UpdateGeneralAttendance{
		GeneralAttendance: *result,
	}, nil
}

func (s *GeneralAttendance) Delete(generalAttendanceId uint) (
	*responses.DeleteGeneralAttendance, error,
) {
	if err := s.generalAttendanceRepo.Delete(generalAttendanceId); err != nil {
		return nil, err
	}

	return &responses.DeleteGeneralAttendance{
		Message: "ok",
	}, nil
}
