package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"api/internal/dto/requests"
	"api/internal/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Lateness struct {
	service *services.Lateness
}

func NewLateness(service *services.Lateness) *Lateness {
	return &Lateness{service}
}

// @id 			createLateness
// @tags 		lateness
// @param 		body body requests.CreateLateness true "Payload"
// @success 	200 {string} string
// @router		/api/v1/latenesses [post]
func (h *Lateness) Create(c *gin.Context) {
	var req requests.CreateLateness
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if err := h.service.Create(authUser.SchoolId, &req); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @id 			createLatenessDetail
// @tags 		lateness
// @param 		lateness_id path int true "Payload"
// @param 		body body requests.CreateLatenessDetail true "Payload"
// @success 	200 {string} string
// @router		/api/v1/latenesses/{lateness_id} [post]
func (h *Lateness) CreateDetail(c *gin.Context) {
	latenessId, err := strconv.Atoi(c.Param("lateness_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateLatenessDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.CreateDetail(uint(latenessId), &req); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @id 			getAllLatenesses
// @tags 		lateness
// @success 	200 {object} responses.GetAllLatenesses
// @router		/api/v1/latenesses [get]
func (h *Lateness) GetAll(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	response, err := h.service.GetAllBySchoolId(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
