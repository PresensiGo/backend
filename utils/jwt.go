package utils

import (
	"api/pkg/authentication"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id uint64, name string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authentication.JWTClaim{
		ID:    id,
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "API Presensi Sekolah",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	tokenString, err := token.SignedString([]byte("password-sementara"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(token string) (*authentication.JWTClaim, error) {
	accessToken, err := jwt.ParseWithClaims(
		token,
		&authentication.JWTClaim{},
		func(t *jwt.Token) (any, error) {
			return []byte("password-sementara"), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := accessToken.Claims.(*authentication.JWTClaim)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}

func GetAuthenticatedUser(ctx context.Context) authentication.AuthenticatedUser {
	authenticatedUser, exists := ctx.Value("token").(authentication.AuthenticatedUser)
	if exists {
		return authenticatedUser
	}

	return authentication.AuthenticatedUser{}
}
