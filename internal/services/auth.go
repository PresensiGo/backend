package services

import (
	"api/internal/dto/responses"
	"api/internal/models"
	"api/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db,
	}
}

func (s *AuthService) Login(email string, password string) (*responses.LoginResponse, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Name, user.Email)
	if err != nil {
		return nil, err
	}
	refreshToken := uuid.New().String()

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		userToken := models.UserToken{
			UserId:       user.ID,
			RefreshToken: refreshToken,
		}

		// delete previous access accessToken
		if err := tx.Where("user_id = ?", user.ID).
			Unscoped().
			Delete(&models.UserToken{}).
			Error; err != nil {
			return err
		}

		// create new access accessToken
		if err := s.db.Create(&userToken).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Register(name string, email string, password string) (*responses.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var response responses.RegisterResponse
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// create user
		user := models.User{
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		// create user accessToken
		accessToken, err := utils.GenerateJWT(user.ID, name, email)
		if err != nil {
			return err
		}
		refreshToken := uuid.New().String()

		userToken := models.UserToken{
			UserId:       user.ID,
			RefreshToken: refreshToken,
		}
		if err := tx.Create(&userToken).Error; err != nil {
			return err
		}

		response = responses.RegisterResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *AuthService) RefreshToken(accessToken string) (*responses.RefreshTokenResponse, error) {
	var userToken models.UserToken
	if err := s.db.Preload("User").
		Where("access_token = ?", accessToken).
		First(&userToken).
		Error; err != nil {
		return nil, err
	}

	// generate new accessToken
	accessToken, err := utils.GenerateJWT(userToken.User.ID, userToken.User.Name, userToken.User.Email)
	if err != nil {
		return nil, err
	}
	newRefreshToken := uuid.New().String()

	// update new accessToken into database
	userToken.RefreshToken = newRefreshToken
	if err := s.db.Save(&userToken).Error; err != nil {
		return nil, err
	}

	return &responses.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
