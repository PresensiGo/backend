package authentication

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
