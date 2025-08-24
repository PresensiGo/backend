package handlers

import (
	"bytes"
	"encoding/base64"
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

	if response, err := h.service.CreateSubjectAttendance(c, uint(classroomId), req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			CreateSubjectAttendanceRecord
// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		subject_attendance_id path int true "subject attendance id"
// @param 		body body requests.CreateSubjectAttendanceRecord true "body"
// @success 	200 {object} responses.CreateSubjectAttendanceRecord
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances/{subject_attendance_id}/records [post]
func (h *SubjectAttendance) CreateSubjectAttendanceRecord(c *gin.Context) {
	subjectAttendanceId, err := strconv.Atoi(c.Param("subject_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateSubjectAttendanceRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.CreateSubjectAttendanceRecord(
		uint(subjectAttendanceId), req,
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

// @id 			CreateSubjectAttendanceRecordStudent
// @tags 		attendance
// @param 		body body requests.CreateSubjectAttendanceRecordStudent true "body"
// @success 	200 {object} responses.CreateSubjectAttendanceRecordStudent
// @router 		/api/v1/subject-attendances/records/student [post]
func (h *SubjectAttendance) CreateSubjectAttendanceRecordStudent(c *gin.Context) {
	student := authentication.GetAuthenticatedStudent(c)
	if student.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var req requests.CreateSubjectAttendanceRecordStudent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.CreateSubjectAttendanceRecordStudent(
		student.Id, req,
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

// @id 			GetAllSubjectAttendancesStudent
// @tags 		attendance
// @success 	200 {object} responses.GetAllSubjectAttendancesStudent
// @router 		/api/v1/subject-attendances/student [get]
func (h *SubjectAttendance) GetAllSubjectAttendancesStudent(c *gin.Context) {
	if response, err := h.service.GetAllSubjectAttendancesStudent(c); err != nil {
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

	if response, err := h.service.GetSubjectAttendance(uint(subjectAttendanceId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			DeleteSubjectAttendance
// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		subject_attendance_id path int true "subject attendance id"
// @success 	200 {object} responses.DeleteSubjectAttendance
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances/{subject_attendance_id} [delete]
func (h *SubjectAttendance) DeleteSubjectAttendance(c *gin.Context) {
	subjectAttendanceId, err := strconv.Atoi(c.Param("subject_attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.DeleteSubjectAttendance(uint(subjectAttendanceId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			DeleteSubjectAttendanceRecord
// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		subject_attendance_id path int true "subject attendance id"
// @param 		record_id path int true "record id"
// @success 	200 {object} responses.DeleteSubjectAttendanceRecord
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances/{subject_attendance_id}/records/{record_id} [delete]
func (h *SubjectAttendance) DeleteSubjectAttendanceRecord(c *gin.Context) {
	recordId, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.DeleteSubjectAttendanceRecord(uint(recordId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		subject_attendance_id path int true "subject attendance id"
// @param 		body body requests.ExportSubjectAttendance true "body"
// @produce     application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @success 	200 {file} file
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances/{subject_attendance_id}/export [post]
func (h *SubjectAttendance) ExportSubjectAttendance(c *gin.Context) {
	var req requests.ExportSubjectAttendance
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if f, err := h.service.ExportSubjectAttendance(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		var b bytes.Buffer
		if err := f.Write(&b); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		encodedStr := base64.StdEncoding.EncodeToString(b.Bytes())
		c.JSON(
			http.StatusOK, gin.H{
				"file":      encodedStr,
				"file_name": "data_siswa.xlsx",
			},
		)

		// fileName := "data_siswa.xlsx"
		// c.Header("Content-Disposition", "attachment; filename="+fileName)
		// c.Header(
		// 	"Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		// )
		// c.Header("Content-Transfer-Encoding", "binary")
		//
		// if err := f.Write(c.Writer); err != nil {
		// 	c.AbortWithStatus(http.StatusInternalServerError)
		// }
	}
}
