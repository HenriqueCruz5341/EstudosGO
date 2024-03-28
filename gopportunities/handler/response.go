package handler

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"status":  code,
		"message": msg,
	})
}

func sendNoContent(ctx *gin.Context) {
	ctx.Status(204)
}

func sendSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"status": code,
		"data":   data,
	})
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
