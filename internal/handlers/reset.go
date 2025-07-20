package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetHandler struct {
	service *services.ResetService
}

func NewResetHandler(service *services.ResetService) *ResetHandler {
	return &ResetHandler{service}
}

func (h *ResetHandler) Reset(c *gin.Context) {
	response, err := h.service.Reset()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
