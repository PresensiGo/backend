package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type ExcelHandler struct {
	service *services.ExcelService
}

func NewExcelHandler(service *services.ExcelService) *ExcelHandler {
	return &ExcelHandler{service}
}

// Import godoc
//
//	@Id			Import
//	@Tags		excel
//	@Success	200	{string}	string
//	@Router		/api/v1/excel/import [post]
func (h *ExcelHandler) Import(c *gin.Context) {
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
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
