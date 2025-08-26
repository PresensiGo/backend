package handlers

import (
	"net/http"
	"path/filepath"

	"api/internal/features/data/services"
	"api/internal/shared/dto/responses"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Excel struct {
	service *services.Excel
}

func NewExcel(service *services.Excel) *Excel {
	return &Excel{service}
}

// @accept 		multipart/form-data
// @param 		file formData file true "file"
// @success 	200 {object} responses.ImportData
// @router 		/api/v1/excel/import-data [post]
func (h *Excel) ImportData(c *gin.Context) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	file, err := c.FormFile("file")
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

	if response, err := h.service.ImportData(user.SchoolId, src); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
