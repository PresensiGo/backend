package services

import (
	"api/internal/dto/responses"
	models2 "api/internal/models"
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
	var user models2.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(user.ID, user.Name, user.Email)
	if err != nil {
		return nil, err
	}
	accessToken := uuid.New().String()

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		userToken := models2.UserToken{
			UserId:      user.ID,
			AccessToken: accessToken,
		}

		// delete previous access token
		if err := tx.Where("user_id = ?", user.ID).
			Unscoped().
			Delete(&models2.UserToken{}).
			Error; err != nil {
			return err
		}

		// create new access token
		if err := s.db.Create(&userToken).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		Token:       token,
		AccessToken: accessToken,
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
		var userId uint
		if err := s.db.Raw(`
			insert into users (name, email, password)
			values (?, ?, ?)
			returning id`,
			name, email, hashedPassword,
		).Find(&userId).Error; err != nil {
			return err
		}

		// create user token
		token, err := utils.GenerateJWT(userId, name, email)
		if err != nil {
			return err
		}
		accessToken := uuid.New().String()

		if err := s.db.Exec(`
			insert into user_tokens (user_id, access_token)
			values (?, ?)`,
			userId, accessToken).Error; err != nil {
			return err
		}

		response = responses.RegisterResponse{
			Token:       token,
			AccessToken: accessToken,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *AuthService) RefreshToken(accessToken string) (*responses.RefreshTokenResponse, error) {
	var userToken models2.UserToken
	if err := s.db.Preload("User").
		Where("access_token = ?", accessToken).
		First(&userToken).
		Error; err != nil {
		return nil, err
	}

	// generate new token
	token, err := utils.GenerateJWT(userToken.User.ID, userToken.User.Name, userToken.User.Email)
	if err != nil {
		return nil, err
	}
	newAccessToken := uuid.New().String()

	// update new token into database
	userToken.AccessToken = newAccessToken
	if err := s.db.Save(&userToken).Error; err != nil {
		return nil, err
	}

	return &responses.RefreshTokenResponse{
		Token:       token,
		AccessToken: newAccessToken,
	}, nil
}
