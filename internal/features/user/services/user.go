package services

import (
	"io"
	"net/http"

	"api/internal/features/user/domains"
	"api/internal/features/user/dto/requests"
	"api/internal/features/user/dto/responses"
	"api/internal/features/user/repositories"
	"api/pkg/authentication"
	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
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

func (s *User) UpdateAccountPassword(
	c *gin.Context, userId uint, req requests.UpdateAccountPassword,
) (
	*responses.UpdateAccountPassword, *failure.App,
) {
	auth := authentication.GetAuthenticatedUser(c)
	if auth.Role != "admin" {
		return nil, failure.NewApp(
			http.StatusForbidden, "Anda tidak memiliki akses untuk melakukan tindakan ini!", nil,
		)
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password), bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	if user, err := s.userRepo.Update(
		userId, domains.User{
			Password: string(hashedPassword),
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.UpdateAccountPassword{
			User: *user,
		}, nil
	}
}

func (s *User) UpdateAccountRole(
	c *gin.Context, userId uint, req requests.UpdateAccountRole,
) (
	*responses.UpdateAccountRole, *failure.App,
) {
	auth := authentication.GetAuthenticatedUser(c)
	if auth.Role != "admin" {
		return nil, failure.NewApp(
			http.StatusForbidden, "Anda tidak memiliki akses untuk melakukan tindakan ini!", nil,
		)
	}

	if user, err := s.userRepo.Update(
		userId, domains.User{
			Role: domains.UserRole(req.Role),
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.UpdateAccountRole{
			User: *user,
		}, nil
	}
}

func (s *User) DeleteAccount(c *gin.Context, accountId uint) (
	*responses.DeleteAccount, *failure.App,
) {
	auth := authentication.GetAuthenticatedUser(c)
	if auth.Role != "admin" {
		return nil, failure.NewApp(
			http.StatusForbidden, "Anda tidak memiliki akses untuk melakukan tindakan ini!", nil,
		)
	}

	if err := s.userRepo.Delete(accountId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.DeleteAccount{
			Message: "ok",
		}, nil
	}
}
