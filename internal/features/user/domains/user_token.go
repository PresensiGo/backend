package domains

import (
	"time"

	"api/internal/features/user/models"
)

type UserToken struct {
	Id           uint      `json:"id" validate:"required"`
	RefreshToken string    `json:"refresh_token" validate:"required"`
	UserId       uint      `json:"user_id" validate:"required"`
	TTL          time.Time `json:"ttl" validate:"required"`
}

func FromUserTokenModel(m *models.UserToken) *UserToken {
	return &UserToken{
		Id:           m.ID,
		RefreshToken: m.RefreshToken,
		UserId:       m.UserId,
		TTL:          m.TTL,
	}
}

func (u *UserToken) ToModel() *models.UserToken {
	return &models.UserToken{
		RefreshToken: u.RefreshToken,
		UserId:       u.UserId,
		TTL:          u.TTL,
	}
}
