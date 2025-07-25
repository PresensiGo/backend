package services

import (
	"api/internal/dto"
	"api/internal/dto/requests"
	"api/internal/dto/responses"
	"api/internal/repository"
	"gorm.io/gorm"
	"time"
)

type Attendance struct {
	db                *gorm.DB
	attendance        *repository.Attendance
	attendanceStudent *repository.AttendanceStudent
}

func NewAttendance(
	db *gorm.DB,
	attendance *repository.Attendance,
	attendanceStudent *repository.AttendanceStudent,
) *Attendance {
	return &Attendance{db, attendance, attendanceStudent}
}

func (s *Attendance) Create(req requests.CreateAttendance) error {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		parsedDate, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return err
		}

		attendance := dto.Attendance{
			ClassroomID: req.ClassroomID,
			Date:        parsedDate,
		}
		if err := s.attendance.Create(tx, &attendance); err != nil {
			return err
		}

		mappedAttendanceStudents := make([]dto.AttendanceStudent, len(req.AttendanceStudents))
		for i, attendanceStudent := range req.AttendanceStudents {
			mappedAttendanceStudents[i] = dto.AttendanceStudent{
				AttendanceID: attendance.ID,
				StudentID:    attendanceStudent.StudentID,
				Status:       attendanceStudent.Status,
				Note:         attendanceStudent.Note,
			}
		}
		if err := s.attendanceStudent.CreateBatch(tx, &mappedAttendanceStudents); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Attendance) GetAll(classId uint) (*responses.GetAllAttendances, error) {
	attendances, err := s.attendance.GetAll(classId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllAttendances{
		Attendances: *attendances,
	}, nil
}
