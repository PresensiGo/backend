package services

import (
	"api/internal/dto/responses"
	"api/internal/repository"
)

type Auth struct {
	auth *repository.Auth
}

func NewAuth(auth *repository.Auth) *Auth {
	return &Auth{auth}
}

func (s *Auth) Login(email string, password string) (*responses.Login, error) {
	token, err := s.auth.Login(email, password)
	if err != nil {
		return nil, err
	}

	return &responses.Login{Token: *token}, nil
}

func (s *Auth) Register(name string, email string, password string) (*responses.Register, error) {
	token, err := s.auth.Register(name, email, password)
	if err != nil {
		return nil, err
	}

	return &responses.Register{Token: *token}, nil
}

func (s *Auth) Logout(userId uint) (*responses.Logout, error) {
	if err := s.auth.Logout(userId); err != nil {
		return nil, err
	}

	return &responses.Logout{}, nil
}

func (s *Auth) RefreshToken(refreshToken string) (*responses.RefreshToken, error) {
	token, err := s.auth.Refresh(refreshToken)
	if err != nil {
		return nil, err
	}

	return &responses.RefreshToken{Token: *token}, nil
}
