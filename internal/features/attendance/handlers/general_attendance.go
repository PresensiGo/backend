package handlers

import (
	"net/http"

	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type GeneralAttendance struct {
	service *services.GeneralAttendance
}

func NewGeneralAttendance(service *services.GeneralAttendance) *GeneralAttendance {
	return &GeneralAttendance{service: service}
}

// @id 			createGeneralAttendance
// @tags 		attendance
// @param 		body body requests.CreateGeneralAttendance true "body"
// @success 	200 {object} responses.CreateGeneralAttendance
// @router 		/api/v1/general_attendances [post]
func (h *GeneralAttendance) Create(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var req requests.CreateGeneralAttendance
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(authUser.SchoolId, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			getAllGeneralAttendances
// @tags 		attendance
// @success 	200 {object} responses.GetAllGeneralAttendances
// @router 		/api/v1/general_attendances [get]
func (h *GeneralAttendance) GetAll(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	result, err := h.service.GetAll(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
