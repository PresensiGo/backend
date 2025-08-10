package middlewares

import (
	"net/http"

	"api/pkg/authentication"
	claims2 "api/pkg/authentication/claims"
	"github.com/gin-gonic/gin"
)

func StudentMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Authorization header missing or invalid"},
			)
			return
		}

		token := authHeader[7:]
		result, err := authentication.VerifyStudentJWT(token)
		if err != nil {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Invalid or expired token"},
			)
			return
		}

		context.Set(
			"token", claims2.Student{
				Id:       result.Id,
				Name:     result.Name,
				NIS:      result.NIS,
				SchoolId: result.SchoolId,
			},
		)

		context.Next()
	}
}
