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

// @id			loginStudent
// @tags 		student
// @param 		body body requests.LoginStudent true "body"
// @success 	200 {object} responses.LoginStudent
// @router 		/api/v1/auth/students/login [post]
func (h *StudentAuth) Login(c *gin.Context) {
	var req requests.LoginStudent
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

// @id			refreshTokenStudent
// @tags 		student
// @param 		body body requests.RefreshTokenStudent true "body"
// @success 	200 {object} responses.RefreshTokenStudent
// @router 		/api/v1/auth/students/refresh-token [post]
func (h *StudentAuth) RefreshToken(c *gin.Context) {
	var req requests.RefreshTokenStudent
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
