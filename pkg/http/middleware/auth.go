package middleware

import (
	"api/pkg/authentication"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Authorization header missing or invalid"},
			)
			return
		}

		token := authHeader[7:]
		claims, err := authentication.VerifyJWT(token)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Invalid or expired token"},
			)
			return
		}

		ctx.Set("token", authentication.AuthenticatedUser{
			ID:    claims.ID,
			Name:  claims.Name,
			Email: claims.Email,
		})

		ctx.Next()
	}
}
