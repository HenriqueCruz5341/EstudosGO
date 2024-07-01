package routes

import (
	"crud-api/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/users/:userId", userController.FindUserById)
	r.GET("/users", userController.FindAllUsers)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:userId", userController.UpdateUser)
	r.DELETE("/users/:userId", userController.DeleteUser)
}
