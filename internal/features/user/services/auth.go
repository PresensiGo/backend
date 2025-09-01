package services

import (
	"errors"
	"net/http"
	"time"

	schoolRepo "api/internal/features/school/repositories"
	"api/internal/features/user/domains"
	"api/internal/features/user/dto/requests"
	"api/internal/features/user/dto/responses"
	"api/internal/features/user/repositories"
	"api/pkg/authentication"
	"api/pkg/http/failure"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	userRepo        *repositories.User
	userTokenRepo   *repositories.UserToken
	userSessionRepo *repositories.UserSession
	schoolRepo      *schoolRepo.School
	db              *gorm.DB
}

func NewAuth(
	userRepo *repositories.User,
	userTokenRepo *repositories.UserToken,
	userSessionRepo *repositories.UserSession,
	schoolRepo *schoolRepo.School,
	db *gorm.DB,
) *Auth {
	return &Auth{
		userRepo:        userRepo,
		userTokenRepo:   userTokenRepo,
		userSessionRepo: userSessionRepo,
		schoolRepo:      schoolRepo,
		db:              db,
	}
}

func (s *Auth) Login(email string, password string) (*responses.Login, *failure.App) {
	currentUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewApp(
				http.StatusNotFound, "Alamat email atau kata sandi tidak valid", err,
			)
		}
		return nil, failure.NewInternal(err)
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword(
		[]byte(currentUser.Password),
		[]byte(password),
	); err != nil {
		return nil, failure.NewApp(
			http.StatusNotFound, "Alamat email atau kata sandi tidak valid", err,
		)
	}

	school, err := s.schoolRepo.Get(currentUser.SchoolId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	// generate token
	accessToken, err := s.generateAccessToken(
		currentUser.Id, currentUser.Name, currentUser.Email, string(currentUser.Role),
		school.Id, school.Name, school.Code,
	)
	if err != nil {
		return nil, failure.NewInternal(err)
	}
	refreshToken := uuid.New().String()

	// store token into database
	if _, err := s.userTokenRepo.Create(
		domains.UserToken{
			UserId:       currentUser.Id,
			RefreshToken: refreshToken,
			TTL:          time.Now().Add(time.Hour * 24 * 30),
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.Login{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
}

func (s *Auth) Login2(req requests.Login2) (*responses.Login2, *failure.App) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewApp(
				http.StatusNotFound, "Alamat email atau kata sandi tidak valid", err,
			)
		}
		return nil, failure.NewInternal(err)
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, failure.NewApp(
			http.StatusNotFound, "Alamat email atau kata sandi tidak valid", err,
		)
	}

	// simpan session ke database
	if session, err := s.userSessionRepo.Create(
		domains.UserSession{
			UserId:    user.Id,
			Token:     uuid.NewString(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.Login2{
			Token: session.Token,
		}, nil
	}
}

func (s *Auth) Logout(refreshToken string) (*responses.Logout, *failure.App) {
	if err := s.userTokenRepo.DeleteByRefreshToken(refreshToken); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &responses.Logout{
				Message: "ok",
			}, nil
		}
		return nil, failure.NewInternal(err)
	} else {
		return &responses.Logout{
			Message: "ok",
		}, nil
	}
}

func (s *Auth) RefreshToken(oldRefreshToken string) (*responses.RefreshToken, *failure.App) {
	oldUserToken, err := s.userTokenRepo.GetByRefreshToken(oldRefreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewApp(
				http.StatusForbidden, "Refresh token tidak ditemukan!", nil,
			)
		}
		return nil, failure.NewInternal(err)
	}

	if time.Now().After(oldUserToken.TTL) {
		return nil, failure.NewApp(http.StatusForbidden, "Token sudah kadaluarsa", nil)
	}

	currentUser, err := s.userRepo.GetByID(oldUserToken.UserId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	school, err := s.schoolRepo.Get(currentUser.SchoolId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	// generate user token
	accessToken, err := s.generateAccessToken(
		currentUser.Id, currentUser.Name, currentUser.Email, string(currentUser.Role),
		school.Id, school.Name, school.Code,
	)
	if err != nil {
		return nil, failure.NewInternal(err)
	}
	refreshToken := uuid.New().String()

	// store new token into database
	if _, err := s.userTokenRepo.UpdateByRefreshToken(
		oldRefreshToken, domains.UserToken{
			RefreshToken: refreshToken,
			TTL:          time.Now().Add(time.Hour * 24 * 30),
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.RefreshToken{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
}

func (s *Auth) generateAccessToken(
	id uint, name string, email string, role string,
	schoolId uint, schoolName string, schoolCode string,
) (string, error) {
	return authentication.GenerateJWT(
		authentication.JWTClaim{
			ID:         id,
			Name:       name,
			Email:      email,
			Role:       role,
			SchoolId:   schoolId,
			SchoolName: schoolName,
			SchoolCode: schoolCode,
		},
	)
}
