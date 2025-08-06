package handlers

import (
	"net/http"

	"api/internal/features/subject/dto/requests"
	"api/internal/features/subject/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Subject struct {
	service *services.Subject
}

func NewSubject(service *services.Subject) *Subject {
	return &Subject{service: service}
}

// @id 			createSubject
// @tags 		subject
// @param 		body body requests.CreateSubject true "body"
// @success 	200 {object} responses.CreateSubject
// @router 		/api/v1/subjects [post]
func (h *Subject) Create(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var req requests.CreateSubject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(authUser.SchoolId, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			getAllSubjects
// @tags 		subject
// @success 	200 {object} responses.GetAllSubjects
// @router 		/api/v1/subjects [get]
func (h *Subject) GetAll(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	result, err := h.service.GetAll(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
