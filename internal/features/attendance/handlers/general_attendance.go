package handlers

import (
	"net/http"
	"strconv"

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

// @id 			createGeneralAttendanceRecordStudent
// @tags 		attendance
// @param 		body body requests.CreateGeneralAttendanceRecordStudent true "body"
// @success 	200 {object} responses.CreateGeneralAttendanceRecordStudent
// @router 		/api/v1/general-attendances/records/student [post]
func (h *GeneralAttendance) CreateRecordStudent(c *gin.Context) {
	studentClaim := authentication.GetAuthenticatedStudent(c)
	if studentClaim.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var req requests.CreateGeneralAttendanceRecordStudent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateGeneralAttendanceRecordStudent(studentClaim.Id, req)
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

// @id 			getGeneralAttendance
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @success 	200 {object} responses.GetGeneralAttendance
// @router 		/api/v1/general_attendances/{general_attendance_id} [get]
func (h *GeneralAttendance) Get(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Get(uint(generalAttendanceId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			updateGeneralAttendance
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @param 		body body requests.UpdateGeneralAttendance true "body"
// @success 	200 {object} responses.UpdateGeneralAttendance
// @router 		/api/v1/general_attendances/{general_attendance_id} [put]
func (h *GeneralAttendance) Update(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateGeneralAttendance
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Update(uint(generalAttendanceId), req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			deleteGeneralAttendance
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @success 	200 {object} responses.DeleteGeneralAttendance
// @router 		/api/v1/general_attendances/{general_attendance_id} [delete]
func (h *GeneralAttendance) Delete(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Delete(uint(generalAttendanceId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
