package services

import (
	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/repository"
	"api/pkg/authentication"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	userRepo      *repository.User
	userTokenRepo *repository.UserToken
	db            *gorm.DB
}

func NewAuth(
	userRepo *repository.User,
	userTokenRepo *repository.UserToken,
	db *gorm.DB,
) *Auth {
	return &Auth{
		userRepo,
		userTokenRepo,
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

	// generate token
	accessToken, err := authentication.GenerateJWT(
		currentUser.ID,
		currentUser.Name,
		currentUser.Email,
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

func (s *Auth) Register(name string, email string, password string) (*responses.Register, error) {
	var response responses.Register
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// create new user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		userID, err := s.userRepo.Create(tx, dto.User{
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		})
		if err != nil {
			return err
		}

		// generate user token
		accessToken, err := authentication.GenerateJWT(userID, name, email)
		if err != nil {
			return err
		}
		refreshToken := uuid.New().String()

		// store token into database
		if err := s.userTokenRepo.Create(tx, dto.UserToken{
			RefreshToken: refreshToken,
			UserID:       userID,
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

	// generate user token
	accessToken, err := authentication.GenerateJWT(
		currentUser.ID,
		currentUser.Name,
		currentUser.Email,
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
