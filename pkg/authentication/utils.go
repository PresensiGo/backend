package authentication

import (
	"context"
	"time"

	"api/internal/features/user/domains"
	"api/pkg/authentication/claims"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user JWTClaim) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, JWTClaim{
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
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			},
		},
	)

	tokenString, err := token.SignedString([]byte("password-sementara"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateStudentJWT(claim claims.Student) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, claims.Student{
			Id:       claim.Id,
			Name:     claim.Name,
			NIS:      claim.NIS,
			SchoolId: claim.SchoolId,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "API Presensi Sekolah",
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			},
		},
	)

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

	data, ok := accessToken.Claims.(*JWTClaim)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return data, nil
}

func VerifyStudentJWT(token string) (*claims.Student, error) {
	accessToken, err := jwt.ParseWithClaims(
		token,
		&claims.Student{},
		func(t *jwt.Token) (any, error) {
			return []byte("password-sementara"), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return nil, err
	}

	data, ok := accessToken.Claims.(*claims.Student)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return data, nil
}

// deprecated
func GetAuthenticatedUser(ctx context.Context) JWTClaim {
	authenticatedUser, exists := ctx.Value("token").(JWTClaim)
	if exists {
		return authenticatedUser
	}

	return JWTClaim{}
}

func GetAuthenticatedUser2(c *gin.Context) *domains.User {
	user, exists := c.Value("user").(domains.User)
	if exists {
		return &user
	} else {
		return nil
	}
}

func GetAuthenticatedStudent(ctx context.Context) claims.Student {
	data, exists := ctx.Value("token").(claims.Student)
	if exists {
		return data
	}

	return claims.Student{}
}
