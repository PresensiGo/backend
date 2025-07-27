package cron

import (
	"api/internal/repository"
	"fmt"
)

type UserToken struct {
	userTokenRepo *repository.UserToken
}

func NewUserToken(userTokenRepo *repository.UserToken) *UserToken {
	return &UserToken{userTokenRepo}
}

func (c *UserToken) Start() {
	if err := c.userTokenRepo.DeleteExpiredTokens(); err != nil {
		fmt.Println(err.Error())
	}
}
