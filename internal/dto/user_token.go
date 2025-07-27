package dto

import "time"

type UserToken struct {
	Id           uint      `json:"id" validate:"required"`
	RefreshToken string    `json:"refresh_token" validate:"required"`
	UserId       uint      `json:"user_id" validate:"required"`
	TTL          time.Time `json:"ttl" validate:"required"`
}
