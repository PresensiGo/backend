package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Reset struct {
	service *services.Reset
}

func NewReset(service *services.Reset) *Reset {
	return &Reset{service}
}

func (h *Reset) Reset(c *gin.Context) {
	response, err := h.service.Reset()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
