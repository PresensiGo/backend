package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type Excel struct {
	service *services.Excel
}

func NewExcel(service *services.Excel) *Excel {
	return &Excel{service}
}

// @ID			importData
// @Tags		excel
// @Success	200	{string}	string
// @Router		/api/v1/excel/import [post]
func (h *Excel) ImportData(c *gin.Context) {
	file, err := c.FormFile("data")
	if file == nil || err != nil {
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

	if _, err := h.service.Import(src); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
