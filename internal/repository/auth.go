package repository

import (
	"api/internal/dto"
	"api/internal/models"
	"api/pkg/authentication"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{db}
}

func (r *Auth) Login(email string, password string) (*dto.Token, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	// generate token
	accessToken, err := authentication.GenerateJWT(user.ID, user.Name, user.Email)
	if err != nil {
		return nil, err
	}
	refreshToken := uuid.New().String()

	// store token into database
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		userToken := models.UserToken{
			UserId:       user.ID,
			RefreshToken: refreshToken,
		}

		// delete previous refresh token
		if err := tx.Where("user_id = ?", user.ID).
			Unscoped().
			Delete(&models.UserToken{}).
			Error; err != nil {
			return err
		}

		// store new refresh token
		if err := r.db.Create(&userToken).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &dto.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *Auth) Register(name string, email string, password string) (*dto.Token, error) {
	var token dto.Token

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		// create user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := models.User{
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		// create user token
		accessToken, err := authentication.GenerateJWT(user.ID, name, email)
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

		token = dto.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *Auth) Logout(userId uint) error {
	if err := r.db.Where("user_id = ?", userId).
		Unscoped().
		Delete(&models.UserToken{}).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *Auth) Refresh(refreshToken string) (*dto.Token, error) {
	var userToken models.UserToken
	if err := r.db.Preload("User").
		Where("refresh_token = ?", refreshToken).
		First(&userToken).
		Error; err != nil {
		return nil, err
	}

	// generate new token
	newAccessToken, err := authentication.GenerateJWT(userToken.User.ID, userToken.User.Name, userToken.User.Email)
	if err != nil {
		return nil, err
	}
	newRefreshToken := uuid.New().String()

	// update new token into database
	userToken.RefreshToken = newRefreshToken
	if err := r.db.Save(&userToken).Error; err != nil {
		return nil, err
	}

	return &dto.Token{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
