package handlers

import (
	"net/http"

	"api/internal/features/school/dto/requests"
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

// @tags 		school
// @param 		body body requests.UpdateSchool true "body"
// @success 	200 {object} responses.UpdateSchool
// @router 		/api/v1/schools/profile [put]
func (h *School) UpdateSchool(c *gin.Context) {
	var req requests.UpdateSchool
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.UpdateSchool(c, req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
