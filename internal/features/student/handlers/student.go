package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/student/dto/requests"
	"api/internal/features/student/services"
	"api/internal/shared/dto/responses"

	"github.com/gin-gonic/gin"
)

type Student struct {
	service *services.Student
}

func NewStudent(service *services.Student) *Student {
	return &Student{service}
}

// @id			GetAllStudentsByClassroomId
// @tags		student
// @param		batch_id path int true "batch id"
// @param		major_id path int true "major id"
// @param		classroom_id path int true "classroom id"
// @success		200	{object} responses.GetAllStudentsByClassroomId
// @router		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/students [get]
func (h *Student) GetAllStudentsByClassroomId(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetAllStudentsByClassroomId(uint(classroomId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @tags		student
// @param		batch_id path int true "batch id"
// @param		major_id path int true "major id"
// @param		classroom_id path int true "classroom id"
// @success		200	{object} responses.GetAllStudentAccountsByClassroomId
// @router		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/student-accounts [get]
func (h *Student) GetAllAccountsByClassroomId(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.GetAllAccountsByClassroomId(uint(classroomId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			getAllStudents
// @tags 		student
// @param 		keyword query string true "Keyword"
// @success 	200 {object} responses.GetAllStudents
// @router 		/api/v1/students [get]
// deprecated
// func (h *Student) GetAll(c *gin.Context) {
// 	keyword := c.Query("keyword")
// 	if len(keyword) == 0 {
// 		c.JSON(
// 			http.StatusOK, responses.GetAllStudents{
// 				Students: make([]domains.StudentMajorClassroom, 0),
// 			},
// 		)
// 		return
// 	}

// 	response, err := h.service.GetAll(keyword)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// @id 			GetProfileStudent
// @tags 		student
// @success 	200 {object} responses.GetProfileStudent
// @router 		/api/v1/students/profile [get]
func (h *Student) GetProfileStudent(c *gin.Context) {
	if response, err := h.service.GetProfileStudent(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @param		batch_id path int true "batch id"
// @param		major_id path int true "major id"
// @param		classroom_id path int true "classroom id"
// @param		student_id path int true "student id"
// @param		body body requests.UpdateStudent true "body"
// @success		200	{object} responses.UpdateStudent
// @router		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/students/{student_id} [put]
func (h *Student) UpdateStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateStudent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.UpdateStudent(uint(studentId), req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @param		batch_id path int true "batch id"
// @param		major_id path int true "major id"
// @param		classroom_id path int true "classroom id"
// @param		student_id path int true "student id"
// @success		200	{object} responses.DeleteStudent
// @router		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/students/{student_id} [delete]
func (h *Student) DeleteStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.DeleteStudent(uint(studentId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
