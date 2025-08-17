package cron

import (
	"fmt"

	"api/internal/injector"
	"github.com/go-co-op/gocron/v2"
)

func New() {
	fmt.Println("cron started")
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println(err)
	}

	_, err = s.NewJob(
		gocron.CronJob("*/5 * * * *", false),
		gocron.NewTask(
			func() {
				fmt.Println("running user token cron")
				injector.InitUserTokenCron().Start()
			},
		),
	)
	if err != nil {
		fmt.Println(err)
	}

	s.Start()
}
