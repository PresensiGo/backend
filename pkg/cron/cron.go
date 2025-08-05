package cron

import (
	"fmt"
	"time"

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
		gocron.DurationJob(10*time.Minute),
		gocron.NewTask(
			func() {
				injector.InitUserTokenCron().Start()
			},
		),
	)
	if err != nil {
		fmt.Println(err)
	}

	s.Start()
}
