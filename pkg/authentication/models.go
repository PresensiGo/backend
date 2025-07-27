package authentication

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	SchoolId   uint   `json:"school_id"`
	SchoolName string `json:"school_name"`
	SchoolCode string `json:"school_code"`

	jwt.RegisteredClaims
}
