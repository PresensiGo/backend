package main

import (
	"os"

	"api/pkg/cron"
	"api/pkg/http"
	"github.com/joho/godotenv"
)

// @title		PresensiGo API
// @version		1.0
func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	if appEnv == "development" {
		if err := godotenv.Load(); err != nil {
			panic(err.Error())
		}
	}

	cron.New()
	http.NewServer()
}
