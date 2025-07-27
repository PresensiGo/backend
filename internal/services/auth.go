package services

import (
	"api/internal/dto"
	"api/internal/dto/requests"
	"api/internal/dto/responses"
	"api/internal/repositories"
	"api/pkg/authentication"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	userRepo      *repositories.User
	userTokenRepo *repositories.UserToken
	schoolRepo    *repositories.School
	db            *gorm.DB
}

func NewAuth(
	userRepo *repositories.User,
	userTokenRepo *repositories.UserToken,
	schoolRepo *repositories.School,
	db *gorm.DB,
) *Auth {
	return &Auth{
		userRepo,
		userTokenRepo,
		schoolRepo,
		db,
	}
}

func (s *Auth) Login(email string, password string) (*responses.Login, error) {
	currentUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword(
		[]byte(currentUser.Password),
		[]byte(password),
	); err != nil {
		return nil, err
	}

	school, err := s.schoolRepo.GetById(currentUser.SchoolId)
	if err != nil {
		return nil, err
	}

	// generate token
	accessToken, err := s.generateAccessToken(
		currentUser.ID, currentUser.Name, currentUser.Email, currentUser.Role,
		school.Id, school.Name, school.Code,
	)
	if err != nil {
		return nil, err
	}
	refreshToken := uuid.New().String()

	// store token into database
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// create user token
		if err := s.userTokenRepo.Create(tx, dto.UserToken{
			UserID:       currentUser.ID,
			RefreshToken: refreshToken,
			TTL:          time.Now().Add(time.Hour * 24 * 30),
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &responses.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Auth) Register(req requests.Register) (*responses.Register, error) {
	// get school by code
	school, err := s.schoolRepo.GetByCode(req.Code)
	if err != nil {
		return nil, err
	}

	var response responses.Register
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// create new user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		userID, err := s.userRepo.Create(tx, dto.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
			Role:     "teacher",
		})
		if err != nil {
			return err
		}

		// generate user token
		accessToken, err := s.generateAccessToken(
			userID, req.Name, req.Email, "teacher",
			school.Id, school.Name, school.Code,
		)
		if err != nil {
			return err
		}
		refreshToken := uuid.New().String()

		// store token into database
		if err := s.userTokenRepo.Create(tx, dto.UserToken{
			RefreshToken: refreshToken,
			UserID:       userID,
			TTL:          time.Now().Add(time.Hour * 24 * 30),
		}); err != nil {
			return err
		}

		response = responses.Register{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *Auth) Logout(refreshToken string) (*responses.Logout, error) {
	if err := s.userTokenRepo.DeleteByRefreshToken(refreshToken); err != nil {
		return nil, err
	}

	return &responses.Logout{}, nil
}

func (s *Auth) RefreshToken(oldRefreshToken string) (*responses.RefreshToken, error) {
	oldUserToken, err := s.userTokenRepo.GetByRefreshToken(oldRefreshToken)
	if err != nil {
		return nil, err
	}

	if time.Now().After(oldUserToken.TTL) {
		return nil, fmt.Errorf("refresh token expired")
	}

	currentUser, err := s.userRepo.GetByID(oldUserToken.UserID)
	if err != nil {
		return nil, err
	}

	school, err := s.schoolRepo.GetById(currentUser.SchoolId)
	if err != nil {
		return nil, err
	}

	// generate user token
	accessToken, err := s.generateAccessToken(
		currentUser.ID, currentUser.Name, currentUser.Email, currentUser.Role,
		school.Id, school.Name, school.Code,
	)
	if err != nil {
		return nil, err
	}
	refreshToken := uuid.New().String()

	// store new token into database
	if err := s.userTokenRepo.UpdateByRefreshToken(oldRefreshToken, dto.UserToken{
		UserID:       currentUser.ID,
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, err
	}

	return &responses.RefreshToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Auth) RefreshTokenTTL(refreshToken string) error {
	return s.userTokenRepo.UpdateTTLByRefreshToken(refreshToken)
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
