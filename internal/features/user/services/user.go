package services

import (
	"api/internal/features/user/dto/responses"
	"api/internal/features/user/repositories"
)

type User struct {
	userRepo *repositories.User
}

func NewUser(userRepo *repositories.User) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (r *User) GetAll(schoolId uint) (*responses.GetAllUsers, error) {
	result, err := r.userRepo.GetAll(schoolId)
	if err != nil {
		return nil, err
	} else {
		return &responses.GetAllUsers{
			Users: *result,
		}, nil
	}
}
