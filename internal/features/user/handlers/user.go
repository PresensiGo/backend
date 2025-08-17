package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"api/internal/features/user/dto/requests"
	"api/internal/features/user/services"
	"api/internal/shared/dto/responses"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type User struct {
	service *services.User
}

func NewUser(service *services.User) *User {
	return &User{
		service: service,
	}
}

// @tags 		account
// @success 	200 {object} responses.GetAllUsers
// @router 		/api/v1/accounts [get]
func (h *User) GetAll(c *gin.Context) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	result, err := h.service.GetAll(user.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @accept 		multipart/form-data
// @param 		file formData file true "file"
// @success 	200 {object} responses.ImportAccounts
// @router 		/api/v1/accounts/import [post]
func (h *User) ImportAccounts(c *gin.Context) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
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

	if response, err := h.service.ImportAccounts(c, user.SchoolId, src); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @tags 		account
// @param 		account_id path int true "account id"
// @param 		body body requests.UpdateAccountPassword true "body"
// @success 	200 {object} responses.UpdateAccountPassword
// @router 		/api/v1/accounts/{account_id}/password [put]
func (h *User) UpdateAccountPassword(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateAccountPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.UpdateAccountPassword(c, uint(accountId), req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @tags 		account
// @param 		account_id path int true "account id"
// @param 		body body requests.UpdateAccountRole true "body"
// @success 	200 {object} responses.UpdateAccountRole
// @router 		/api/v1/accounts/{account_id}/role [put]
func (h *User) UpdateAccountRole(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateAccountRole
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.UpdateAccountRole(c, uint(accountId), req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @tags 		account
// @param 		account_id path int true "account id"
// @success 	200 {object} responses.DeleteAccount
// @router 		/api/v1/accounts/{account_id} [delete]
func (h *User) DeleteAccount(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.DeleteAccount(c, uint(accountId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
