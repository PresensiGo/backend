package services

import (
	"api/internal/dto"
	"api/internal/dto/requests"
	"api/internal/repository"
	"gorm.io/gorm"
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
		attendance := dto.Attendance{
			ClassID: req.ClassID,
			Date:    req.Date,
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
