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

type SubjectAttendance struct {
	service *services.SubjectAttendance
}

func NewSubjectAttendance(service *services.SubjectAttendance) *SubjectAttendance {
	return &SubjectAttendance{
		service: service,
	}
}

// @id 			CreateSubjectAttendance
// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		body body requests.CreateSubjectAttendance true "body"
// @success 	200 {object} responses.CreateSubjectAttendance
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances [post]
func (h *SubjectAttendance) CreateSubjectAttendance(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateSubjectAttendance
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.CreateSubjectAttendance(uint(classroomId), req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			createSubjectAttendanceRecordStudent
// @tags 		attendance
// @param 		body body requests.CreateSubjectAttendanceRecordStudent true "body"
// @success 	200 {object} responses.CreateSubjectAttendanceRecordStudent
// @router 		/api/v1/subject-attendances/records/student [post]
func (h *SubjectAttendance) CreateRecordStudent(c *gin.Context) {
	studentClaims := authentication.GetAuthenticatedStudent(c)
	if studentClaims.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var req requests.CreateSubjectAttendanceRecordStudent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateRecordStudent(studentClaims.Id, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			GetAllSubjectAttendances
// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @success 	200 {object} responses.GetAllSubjectAttendances
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances [get]
func (h *SubjectAttendance) GetAllSubjectAttendances(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetAllSubjectAttendances(uint(classroomId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			GetAllSubjectAttendanceRecords
// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		subject_attendance_id path int true "subject attendance id"
// @success 	200 {object} responses.GetAllSubjectAttendanceRecords
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances/{subject_attendance_id}/records [get]
func (h *SubjectAttendance) GetAllSubjectAttendanceRecords(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	subjectAttendanceId, err := strconv.Atoi(c.Param("subject_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetAllSubjectAttendanceRecords(
		uint(classroomId), uint(subjectAttendanceId),
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

// @id 			GetSubjectAttendance
// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		subject_attendance_id path int true "subject attendance id"
// @success 	200 {object} responses.GetSubjectAttendance
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances/{subject_attendance_id} [get]
func (h *SubjectAttendance) GetSubjectAttendance(c *gin.Context) {
	subjectAttendanceId, err := strconv.Atoi(c.Param("subject_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.GetSubjectAttendance(uint(subjectAttendanceId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
