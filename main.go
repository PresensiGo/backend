package main

import (
	"api/pkg/http"
	"fmt"
	"github.com/joho/godotenv"
)

// @title		PresensiGo API
// @version	1.0
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file")
	}

	http.NewServer()
}
