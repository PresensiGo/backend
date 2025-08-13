package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/classroom/dto/requests"
	"api/internal/features/classroom/services"
	"api/internal/shared/dto/responses"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Classroom struct {
	service *services.Classroom
}

func NewClassroom(service *services.Classroom) *Classroom {
	return &Classroom{service}
}

// @tags		classroom
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		major_id path int true "major id"
// @param 		body body requests.CreateClassroom true "body"
// @success		200	{object} responses.CreateClassroom
// @router		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms [post]
func (h *Classroom) Create(c *gin.Context) {
	majorId, err := strconv.Atoi(c.Param("major_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateClassroom
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(uint(majorId), req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			getAllClassrooms
// @tags 		classroom
// @success 	200 {object} responses.GetAll
// @router 		/api/v1/classrooms [get]
func (h *Classroom) GetAll(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.GetAll(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id			getAllClassroomsByMajorId
// @tags		classroom
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @success		200	{object} responses.GetAllClassroomsByMajorId
// @router		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms [get]
func (h *Classroom) GetAllClassroomsByMajorId(c *gin.Context) {
	majorId, err := strconv.Atoi(c.Param("major_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetAllClassroomsByMajorId(uint(majorId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @Id			getAllClassroomWithMajors
// @Tags		classroom
// @Param 		batch_id path int true "Batch Id"
// @Success		200	{object}	responses.GetAllClassroomWithMajors
// @Router		/api/v1/classrooms/batches/{batch_id} [get]
func (h *Classroom) GetAllWithMajors(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllWithMajor(uint(batchId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @tags		classroom
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		body body requests.UpdateClassroom true "body"
// @success		200	{object} responses.UpdateClassroom
// @router		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id} [put]
func (h *Classroom) Update(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateClassroom
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Update(uint(classroomId), req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
