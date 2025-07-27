package cron

import (
	"api/internal/repositories"
	"fmt"
)

type UserToken struct {
	userTokenRepo *repositories.UserToken
}

func NewUserToken(userTokenRepo *repositories.UserToken) *UserToken {
	return &UserToken{userTokenRepo}
}

func (c *UserToken) Start() {
	if err := c.userTokenRepo.DeleteExpiredTokens(); err != nil {
		fmt.Println(err.Error())
	}
}
