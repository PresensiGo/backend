package repositories

import (
	"time"

	"api/internal/features/user/domains"
	"api/internal/features/user/models"
	"gorm.io/gorm"
)

type UserToken struct {
	db *gorm.DB
}

func NewUserToken(db *gorm.DB) *UserToken {
	return &UserToken{db}
}

func (r *UserToken) Create(tx *gorm.DB, token domains.UserToken) error {
	userToken := models.UserToken{
		UserId:       token.UserId,
		RefreshToken: token.RefreshToken,
		TTL:          token.TTL,
	}
	if err := tx.Create(&userToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserToken) GetByRefreshToken(refreshToken string) (*domains.UserToken, error) {
	var userToken models.UserToken
	if err := r.db.Where("refresh_token = ?", refreshToken).
		First(&userToken).
		Error; err != nil {
		return nil, err
	}

	return &domains.UserToken{
		Id:           userToken.ID,
		RefreshToken: userToken.RefreshToken,
		UserId:       userToken.UserId,
		TTL:          userToken.TTL,
	}, nil
}

func (r *UserToken) UpdateByRefreshToken(oldRefreshToken string, token domains.UserToken) error {
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
