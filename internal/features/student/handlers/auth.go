package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/student/dto/requests"
	_ "api/internal/features/student/dto/responses"
	"api/internal/features/student/services"
	"api/internal/shared/dto/responses"
	_ "api/internal/shared/dto/responses"

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
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id			RefreshTokenStudent
// @tags 		student
// @param 		body body requests.RefreshTokenStudent true "body"
// @success 	200 {object} responses.RefreshTokenStudent
// @router 		/api/v1/auth/students/refresh-token [post]
func (h *StudentAuth) RefreshTokenStudent(c *gin.Context) {
	var req requests.RefreshTokenStudent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.RefreshTokenStudent(req); err != nil {
		c.AbortWithStatusJSON(err.Code, responses.Error{
			Message: err.Message,
		})
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @tags 		student
// @param 		student_token_id path int true "student token id"
// @success 	200 {object} responses.EjectStudentToken
// @router 		/api/v1/auth/students/accounts/{student_token_id}/eject [post]
func (h *StudentAuth) Eject(c *gin.Context) {
	studentTokenId, err := strconv.Atoi(c.Param("student_token_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Eject(uint(studentTokenId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
