package router

import (
	"gopportunities/handler"

	"github.com/gin-gonic/gin"

	docs "gopportunities/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.Init()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handler.FindAllOpeningsHandler)
		v1.GET("/openings/:id", handler.FindByIdOpeningsHandler)
		v1.POST("/openings", handler.CreateOpeningsHandler)
		v1.PATCH("/openings/:id", handler.UpdateDataOpeningsHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningsHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
