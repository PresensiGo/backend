package main

import (
	"api/pkg/http"
	"github.com/joho/godotenv"
)

// @title		PresensiGo API
// @version	1.0
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	http.NewServer()
}
