package authentication

import (
	"net/http"

	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
)

func ValidateAdmin(c *gin.Context) *failure.App {
	auth := GetAuthenticatedUser(c)
	if auth.Role != "admin" {
		return failure.NewApp(
			http.StatusForbidden,
			"Anda tidak memiliki akses untuk melakukan tindakan ini!",
			nil,
		)
	}

	return nil
}
