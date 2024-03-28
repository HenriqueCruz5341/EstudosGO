package handler

import (
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create a new opening
// @Description Create a new opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body OpeningInputDto true "Request body"
// @Success 201 {object} SuccessResponse{data=schemas.OpeningOutputDto} "desc"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [post]
func CreateOpeningsHandler(ctx *gin.Context) {
	request := OpeningInputDto{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	openingOutput := schemas.OpeningOutputDto{
		Id:        opening.ID,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
		CreatedAt: opening.CreatedAt,
		UpdatedAt: opening.UpdatedAt,
		DeletedAt: opening.DeletedAt.Time,
	}
	sendSuccess(ctx, http.StatusCreated, openingOutput)
}
