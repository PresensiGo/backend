package handlers

import (
	"net/http"

	"api/internal/features/student/dto/requests"
	_ "api/internal/features/student/dto/responses"
	"api/internal/features/student/services"
	"github.com/gin-gonic/gin"
)

type StudentAuth struct {
	service *services.StudentAuth
}

func NewStudentAuth(service *services.StudentAuth) *StudentAuth {
	return &StudentAuth{
		service: service,
	}
}

// @tags 		student
// param 		body body requests.StudentLogin true "body"
// @success 	200 {object} responses.StudentLogin
// @router 		/api/v1/auth/students/login [post]
func (h *StudentAuth) Login(c *gin.Context) {
	var req requests.StudentLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Login(req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @tags 		student
// param 		body body requests.StudentRefreshToken true "body"
// @success 	200 {object} responses.StudentRefreshToken
// @router 		/api/v1/auth/students/refresh-token [post]
func (h *StudentAuth) RefreshToken(c *gin.Context) {
	var req requests.StudentRefreshToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.RefreshToken(req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
