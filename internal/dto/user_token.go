package dto

type UserToken struct {
	ID           uint   `json:"id" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	UserID       uint   `json:"user_id" validate:"required"`
}
