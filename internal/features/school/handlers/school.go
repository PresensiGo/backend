package handlers

import (
	"net/http"

	"api/internal/features/school/services"
	"api/internal/shared/dto/responses"
	"github.com/gin-gonic/gin"
)

type School struct {
	service *services.School
}

func NewSchool(
	service *services.School,
) *School {
	return &School{
		service: service,
	}
}

// @id 			GetSchool
// @tags 		school
// @success 	200 {object} responses.GetSchool
// @router 		/api/v1/schools/profile [get]
func (h *School) GetSchool(c *gin.Context) {
	if response, err := h.service.GetSchool(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
