package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Class struct {
	service *services.Class
}

func NewClass(service *services.Class) *Class {
	return &Class{service}
}

func (h *Class) GetAllClasses(c *gin.Context) {
	majorId, err := strconv.ParseUint(c.Param("major_id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllClasses(majorId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
