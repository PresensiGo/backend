package main

import (
	"fmt"
	"os"

	"api/pkg/cron"
	"api/pkg/http"
	"github.com/joho/godotenv"
)

// @title		PresensiGo API
// @version	1.0
func main() {
	fmt.Println("appenv", os.Getenv("APP_ENV"))
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file")
	}

	cron.New()
	http.NewServer()
}
