package services

import (
	"time"

	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
	"github.com/google/uuid"
)

type GeneralAttendance struct {
	generalAttendanceRepo *repositories.GeneralAttendance
}

func NewGeneralAttendance(generalAttendanceRepo *repositories.GeneralAttendance) *GeneralAttendance {
	return &GeneralAttendance{generalAttendanceRepo: generalAttendanceRepo}
}

func (s *GeneralAttendance) Create(
	schoolId uint, req requests.CreateGeneralAttendance,
) (*responses.CreateGeneralAttendance, error) {
	timezone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	parsedDateTime, err := time.ParseInLocation("2006-01-02 15:04:05", req.DateTime, timezone)

	generalAttendance := domains.GeneralAttendance{
		DateTime: parsedDateTime,
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

func (s *GeneralAttendance) GetAll(schoolId uint) (*responses.GetAllGeneralAttendances, error) {
	result, err := s.generalAttendanceRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllGeneralAttendances{
		GeneralAttendances: *result,
	}, nil
}
