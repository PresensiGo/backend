package services

import (
	"time"

	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
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
	parsedDate, err := time.Parse("2006-01-02", req.Date)

	generalAttendance := domains.GeneralAttendance{
		Date:     parsedDate,
		Note:     req.Note,
		SchoolId: schoolId,
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
