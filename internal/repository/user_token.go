package repository

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type UserToken struct {
	db *gorm.DB
}

func NewUserToken(db *gorm.DB) *UserToken {
	return &UserToken{db}
}

func (r *UserToken) Create(tx *gorm.DB, token dto.UserToken) error {
	userToken := models.UserToken{
		UserId:       token.UserID,
		RefreshToken: token.RefreshToken,
	}
	if err := tx.Create(&userToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserToken) GetByRefreshToken(refreshToken string) (*dto.UserToken, error) {
	var userToken models.UserToken
	if err := r.db.Where("refresh_token = ?", refreshToken).
		First(&userToken).
		Error; err != nil {
		return nil, err
	}

	return &dto.UserToken{
		ID:           userToken.ID,
		RefreshToken: userToken.RefreshToken,
		UserID:       userToken.UserId,
	}, nil
}

func (r *UserToken) UpdateByRefreshToken(oldRefreshToken string, token dto.UserToken) error {
	return r.db.Model(&models.UserToken{}).
		Where("refresh_token = ?", oldRefreshToken).
		Update("refresh_token", token.RefreshToken).
		Error
}

func (r *UserToken) DeleteByRefreshToken(refreshToken string) error {
	return r.db.Where("refresh_token = ?", refreshToken).
		Unscoped().
		Delete(&models.UserToken{}).
		Error
}
