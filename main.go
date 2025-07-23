package main

import (
	"api/pkg/http"
	"github.com/joho/godotenv"
	"log"
)

// @title		PresensiGo API
// @version	1.0
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	http.NewServer()
}
