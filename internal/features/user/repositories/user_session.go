package repositories

import (
	"api/internal/features/user/domains"
	"api/internal/features/user/models"
	"gorm.io/gorm"
)

type UserSession struct {
	db *gorm.DB
}

func NewUserSession(
	db *gorm.DB,
) *UserSession {
	return &UserSession{
		db: db,
	}
}

func (r *UserSession) Create(data domains.UserSession) (*domains.UserSession, error) {
	userSession := data.ToModel()
	if err := r.db.Create(&userSession).Error; err != nil {
		return nil, err
	} else {
		return domains.FromUserSessionModel(userSession), nil
	}
}

func (r *UserSession) GetByToken(token string) (*domains.UserSession, error) {
	var userSession models.UserSession
	if err := r.db.Where("token = ?", token).First(&userSession).Error; err != nil {
		return nil, err
	} else {
		return domains.FromUserSessionModel(&userSession), nil
	}
}

func (r *UserSession) UpdateByToken(token string, data domains.UserSession) (
	*domains.UserSession, error,
) {
	userSession := data.ToModel()
	if err := r.db.Where("token = ?", token).Updates(&userSession).Error; err != nil {
		return nil, err
	} else {
		return domains.FromUserSessionModel(userSession), nil
	}
}

func (r *UserSession) DeleteByToken(token string) error {
	return r.db.Where("token = ?", token).Unscoped().Delete(&models.UserSession{}).Error
}
