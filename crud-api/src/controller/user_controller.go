package controller

import (
	"crud-api/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewUserController(userService service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: userService,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
