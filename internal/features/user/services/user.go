package services

import (
	"io"
	"net/http"

	"api/internal/features/user/domains"
	"api/internal/features/user/dto/responses"
	"api/internal/features/user/repositories"
	"api/pkg/http/failure"
	"github.com/xuri/excelize/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	userRepo *repositories.User
}

func NewUser(userRepo *repositories.User) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (s *User) ImportAccounts(schoolId uint, reader io.Reader) (
	*responses.ImportAccounts, *failure.App,
) {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, failure.NewInternal(err)
	}
	defer file.Close()

	sheets := file.GetSheetList()
	if len(sheets) == 0 {
		return nil, failure.NewApp(http.StatusBadRequest, "Sheet tidak ditemukan!", err)
	}

	firstSheetName := sheets[0]
	rows, err := file.GetRows(firstSheetName)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	if len(rows) < 2 {
		return nil, failure.NewApp(http.StatusBadRequest, "Jumlah baris tidak mencukupi!", err)
	}

	var users []domains.User
	for _, v := range rows[1:] {
		userName := v[0]
		userEmail := v[1]
		userPassword := v[2]

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(userPassword), bcrypt.DefaultCost,
		)
		if err != nil {
			return nil, failure.NewInternal(err)
		}

		users = append(
			users, domains.User{
				Name:     userName,
				Email:    userEmail,
				Password: string(hashedPassword),
				Role:     domains.TeacherUserRole,
				SchoolId: schoolId,
			},
		)
	}

	if _, err = s.userRepo.CreateBatch(users); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.ImportAccounts{
			Message: "ok",
		}, nil
	}
}

func (s *User) GetAll(schoolId uint) (*responses.GetAllUsers, error) {
	result, err := s.userRepo.GetAll(schoolId)
	if err != nil {
		return nil, err
	} else {
		return &responses.GetAllUsers{
			Users: *result,
		}, nil
	}
}
