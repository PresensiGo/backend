package services

import (
	"api/internal/dto/responses"
	"api/internal/models"
	"api/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{
		db,
	}
}

func (s *Auth) Login(email string, password string) (*responses.Login, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateJWT(uint64(user.ID), user.Name, user.Email)
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

	return &responses.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Auth) Register(name string, email string, password string) (*responses.Register, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var response responses.Register
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
		accessToken, err := utils.GenerateJWT(uint64(user.ID), name, email)
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

func (s *Auth) Logout(userId uint64) (*responses.Logout, error) {
	if err := s.db.Where("user_id = ?", userId).
		Delete(&models.UserToken{}).
		Error; err != nil {
		return nil, err
	}

	return &responses.Logout{}, nil
}

func (s *Auth) RefreshToken(accessToken string) (*responses.RefreshToken, error) {
	var userToken models.UserToken
	if err := s.db.Preload("User").
		Where("access_token = ?", accessToken).
		First(&userToken).
		Error; err != nil {
		return nil, err
	}

	// generate new accessToken
	accessToken, err := utils.GenerateJWT(uint64(userToken.User.ID), userToken.User.Name, userToken.User.Email)
	if err != nil {
		return nil, err
	}
	newRefreshToken := uuid.New().String()

	// update new accessToken into database
	userToken.RefreshToken = newRefreshToken
	if err := s.db.Save(&userToken).Error; err != nil {
		return nil, err
	}

	return &responses.RefreshToken{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
