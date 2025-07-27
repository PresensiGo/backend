package authentication

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user JWTClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		SchoolId:   user.SchoolId,
		SchoolName: user.SchoolName,
		SchoolCode: user.SchoolCode,
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

func VerifyJWT(token string) (*JWTClaim, error) {
	accessToken, err := jwt.ParseWithClaims(
		token,
		&JWTClaim{},
		func(t *jwt.Token) (any, error) {
			return []byte("password-sementara"), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := accessToken.Claims.(*JWTClaim)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}

func GetAuthenticatedUser(ctx context.Context) JWTClaim {
	authenticatedUser, exists := ctx.Value("token").(JWTClaim)
	if exists {
		return authenticatedUser
	}

	return JWTClaim{}
}
