package handlers

import (
	"net/http"

	"api/internal/features/user/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type User struct {
	service *services.User
}

func NewUser(service *services.User) *User {
	return &User{
		service: service,
	}
}

// @tags 		account
// @success 	200 {object} responses.GetAllUsers
// @router 		/api/v1/accounts [get]
func (h *User) GetAll(c *gin.Context) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	result, err := h.service.GetAll(user.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
