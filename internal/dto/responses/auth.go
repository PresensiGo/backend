package responses

import "api/internal/dto"

type Login struct {
	Token dto.Token `json:"token" validate:"required"`
}

type Register struct {
	Token dto.Token `json:"token" validate:"required"`
}

type Logout struct{}

type RefreshToken struct {
	Token dto.Token `json:"token" validate:"required"`
}
