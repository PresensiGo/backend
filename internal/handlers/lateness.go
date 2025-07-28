package handlers

import (
	"api/internal/dto/requests"
	"api/internal/services"
	"api/pkg/authentication"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Lateness struct {
	service *services.Lateness
}

func NewLateness(service *services.Lateness) *Lateness {
	return &Lateness{service}
}

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
