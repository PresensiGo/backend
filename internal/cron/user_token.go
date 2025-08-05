package cron

import (
	"fmt"

	"api/internal/features/user/repositories"
)

type UserTokenCron struct {
	userTokenRepo *repositories.UserToken
}

func NewUserTokenCron(userTokenRepo *repositories.UserToken) *UserTokenCron {
	return &UserTokenCron{userTokenRepo}
}

func (c *UserTokenCron) Start() {
	if err := c.userTokenRepo.DeleteExpiredTokens(); err != nil {
		fmt.Println(err.Error())
	}
}
