package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/services"
	"api/internal/shared/dto/responses"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type GeneralAttendance struct {
	service *services.GeneralAttendance
}

func NewGeneralAttendance(service *services.GeneralAttendance) *GeneralAttendance {
	return &GeneralAttendance{service: service}
}

// @id 			CreateGeneralAttendance
// @tags 		attendance
// @param 		body body requests.CreateGeneralAttendance true "body"
// @success 	200 {object} responses.CreateGeneralAttendance
// @router 		/api/v1/general-attendances [post]
func (h *GeneralAttendance) CreateGeneralAttendance(c *gin.Context) {
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

	if response, err := h.service.CreateGeneralAttendance(c, authUser.SchoolId, req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			CreateGeneralAttendanceRecord
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @param 		body body requests.CreateGeneralAttendanceRecord true "body"
// @success 	200 {object} responses.CreateGeneralAttendanceRecord
// @router 		/api/v1/general-attendances/{general_attendance_id}/records [post]
func (h *GeneralAttendance) CreateGeneralAttendanceRecord(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateGeneralAttendanceRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.CreateGeneralAttendanceRecord(
		uint(generalAttendanceId), req,
	); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
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

// @id 			GetAllGeneralAttendances
// @tags 		attendance
// @success 	200 {object} responses.GetAllGeneralAttendances
// @router 		/api/v1/general-attendances [get]
func (h *GeneralAttendance) GetAllGeneralAttendances(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if response, err := h.service.GetAllGeneralAttendances(authUser.SchoolId); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			GetGeneralAttendance
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @success 	200 {object} responses.GetGeneralAttendance
// @router 		/api/v1/general-attendances/{general_attendance_id} [get]
func (h *GeneralAttendance) GetGeneralAttendance(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetGeneralAttendance(uint(generalAttendanceId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			GetAllGeneralAttendanceRecords
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @success 	200 {object} responses.GetAllGeneralAttendanceRecords
// @router 		/api/v1/general-attendances/{general_attendance_id}/records [get]
func (h *GeneralAttendance) GetAllGeneralAttendanceRecords(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetAllGeneralAttendanceRecords(uint(generalAttendanceId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			GetAllGeneralAttendanceRecordsByClassroomId
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @param 		classroom_id path int true "classroom id"
// @success 	200 {object} responses.GetAllGeneralAttendanceRecordsByClassroomId
// @router 		/api/v1/general-attendances/{general_attendance_id}/classrooms/{classroom_id}/records [get]
func (h *GeneralAttendance) GetAllGeneralAttendanceRecordsByClassroomId(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetAllGeneralAttendanceRecordsByClassroomId(
		uint(generalAttendanceId),
		uint(classroomId),
	); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
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
// @router 		/api/v1/general-attendances/{general_attendance_id} [delete]
func (h *GeneralAttendance) DeleteGeneralAttendance(c *gin.Context) {
	generalAttendanceId, err := strconv.Atoi(c.Param("general_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.DeleteGeneralAttendance(uint(generalAttendanceId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			DeleteGeneralAttendanceRecord
// @tags 		attendance
// @param 		general_attendance_id path int true "general attendance id"
// @param 		record_id path int true "record id"
// @success 	200 {object} responses.DeleteGeneralAttendanceRecord
// @router 		/api/v1/general-attendances/{general_attendance_id}/records/{record_id} [delete]
func (h *GeneralAttendance) DeleteGeneralAttendanceRecord(c *gin.Context) {
	recordId, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.DeleteGeneralAttendanceRecord(uint(recordId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
