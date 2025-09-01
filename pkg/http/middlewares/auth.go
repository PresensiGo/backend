package middlewares

import (
	"net/http"
	"time"

	"api/internal/features/user/repositories"
	"api/pkg/database"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	var (
		db              = database.New()
		userRepo        = repositories.NewUser(db)
		userSessionRepo = repositories.NewUserSession(db)
	)

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
		session, err := userSessionRepo.GetByToken(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// validasi expires token
		if time.Now().After(session.ExpiresAt) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// mendapatkan user berdasarkan id dari session
		user, err := userRepo.GetByID(session.UserId)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)
		ctx.Next()

		// claims, err := authentication.VerifyJWT(token)
		// if err != nil {
		// 	fmt.Println(err)
		// 	ctx.AbortWithStatusJSON(
		// 		http.StatusUnauthorized,
		// 		gin.H{"error": "Invalid or expired token"},
		// 	)
		// 	return
		// }
		//
		// ctx.Set(
		// 	"token", authentication.JWTClaim{
		// 		ID:         claims.ID,
		// 		Name:       claims.Name,
		// 		Email:      claims.Email,
		// 		Role:       claims.Role,
		// 		SchoolId:   claims.SchoolId,
		// 		SchoolName: claims.SchoolName,
		// 		SchoolCode: claims.SchoolCode,
		// 	},
		// )
	}
}
