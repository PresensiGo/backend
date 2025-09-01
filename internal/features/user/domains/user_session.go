package domains

import (
	"time"

	"api/internal/features/user/models"
)

type UserSession struct {
	Id        uint
	UserId    uint
	Token     string
	ExpiresAt time.Time
}

func (u *UserSession) ToModel() *models.UserSession {
	return &models.UserSession{
		UserId:    u.UserId,
		Token:     u.Token,
		ExpiresAt: u.ExpiresAt,
	}
}

func FromUserSessionModel(m *models.UserSession) *UserSession {
	return &UserSession{
		Id:        m.ID,
		UserId:    m.UserId,
		Token:     m.Token,
		ExpiresAt: m.ExpiresAt,
	}
}
