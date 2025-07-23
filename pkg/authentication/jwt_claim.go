package authentication

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
