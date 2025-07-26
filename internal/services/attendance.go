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
	studentRepo       *repository.Student
}

func NewAttendance(
	db *gorm.DB,
	attendance *repository.Attendance,
	attendanceStudent *repository.AttendanceStudent,
	studentRepo *repository.Student,
) *Attendance {
	return &Attendance{
		db,
		attendance,
		attendanceStudent,
		studentRepo,
	}
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

func (s *Attendance) GetById(attendanceId uint) (*responses.GetAttendance, error) {
	attendance, err := s.attendance.GetById(attendanceId)
	if err != nil {
		return nil, err
	}

	attendanceStudents, err := s.attendanceStudent.GetAllByAttendanceId(attendanceId)
	if err != nil {
		return nil, err
	}

	mapAttendanceStudents := make(map[uint]dto.AttendanceStudent)
	studentIds := make([]uint, len(*attendanceStudents))
	for i, item := range *attendanceStudents {
		studentIds[i] = item.StudentID
		mapAttendanceStudents[item.StudentID] = item
	}

	students, err := s.studentRepo.GetManyById(studentIds)
	if err != nil {
		return nil, err
	}

	mappedItems := make([]responses.GetAttendanceItem, len(*students))
	for i, item := range *students {
		mappedItems[i] = responses.GetAttendanceItem{
			Student:           item,
			AttendanceStudent: mapAttendanceStudents[item.ID],
		}
	}

	return &responses.GetAttendance{
		Attendance: *attendance,
		Items:      mappedItems,
	}, nil
}

func (s *Attendance) Delete(attendanceID uint) error {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// delete all attendance students
		if err := s.attendanceStudent.DeleteByAttendanceID(tx, attendanceID); err != nil {
			return err
		}

		// delete attendance
		if err := s.attendance.DeleteByID(tx, attendanceID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
