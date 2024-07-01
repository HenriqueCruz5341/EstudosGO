package controller

import (
	"crud-api/src/configuration/logger"
	"crud-api/src/configuration/validation"
	"crud-api/src/controller/model/request"
	"crud-api/src/model"
	"crud-api/src/view"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetId()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
