package services

import (
	"errors"
	"io"

	"api/internal/features/teacher/dto/responses"
	"api/internal/features/user/domains"
	"api/internal/features/user/repositories"
	"github.com/xuri/excelize/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Teacher struct {
	db       *gorm.DB
	userRepo *repositories.User
}

func NewTeacher(db *gorm.DB, userRepo *repositories.User) *Teacher {
	return &Teacher{
		db:       db,
		userRepo: userRepo,
	}
}

func (s *Teacher) Import(schoolId uint, reader io.Reader) (*responses.ImportTeacher, error) {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sheets := file.GetSheetList()
	if len(sheets) == 0 {
		return nil, errors.New("sheets is empty")
	}

	firstSheetName := sheets[0]
	rows, err := file.GetRows(firstSheetName)
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, errors.New("rows not enough")
	}

	var teachers []domains.User
	for _, v := range rows[1:] {
		teacherName := v[0]
		teacherEmail := v[1]
		teacherPassword := v[2]

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(teacherPassword), bcrypt.DefaultCost,
		)
		if err != nil {
			return nil, err
		}

		teachers = append(
			teachers, domains.User{
				Name:     teacherName,
				Email:    teacherEmail,
				Password: string(hashedPassword),
				Role:     domains.TeacherUserRole,
				SchoolId: schoolId,
			},
		)
	}

	if _, err = s.userRepo.CreateBatch(teachers); err != nil {
		return nil, err
	} else {
		return &responses.ImportTeacher{
			Message: "ok",
		}, nil
	}
}
