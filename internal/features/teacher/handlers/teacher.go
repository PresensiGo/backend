package handlers

import (
	"net/http"
	"path/filepath"

	"api/internal/features/teacher/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Teacher struct {
	service *services.Teacher
}

func NewTeacher(service *services.Teacher) *Teacher {
	return &Teacher{
		service: service,
	}
}

func (h *Teacher) Import(c *gin.Context) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".xlsx" && ext != ".xls" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	src, err := file.Open()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer src.Close()

	result, err := h.service.Import(user.SchoolId, src)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
