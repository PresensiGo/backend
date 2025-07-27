package repository

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
	"time"
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
		TTL:          token.TTL,
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
		TTL:          userToken.TTL,
	}, nil
}

func (r *UserToken) UpdateByRefreshToken(oldRefreshToken string, token dto.UserToken) error {
	return r.db.Model(&models.UserToken{}).
		Where("refresh_token = ?", oldRefreshToken).
		Update("refresh_token", token.RefreshToken).
		Error
}

func (r *UserToken) UpdateTTLByRefreshToken(refreshToken string) error {
	return r.db.Model(&models.UserToken{}).
		Where("refresh_token = ?", refreshToken).
		Update("ttl", time.Now().Add(time.Hour*24*30)).
		Error
}

func (r *UserToken) DeleteByRefreshToken(refreshToken string) error {
	return r.db.Where("refresh_token = ?", refreshToken).
		Unscoped().
		Delete(&models.UserToken{}).
		Error
}

func (r *UserToken) DeleteExpiredTokens() error {
	return r.db.Where("ttl < ?", time.Now()).
		Unscoped().
		Delete(&models.UserToken{}).
		Error
}
