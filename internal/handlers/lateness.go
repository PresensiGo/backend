package handlers

import (
	"api/internal/dto/requests"
	"api/internal/services"
	"api/pkg/authentication"
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
