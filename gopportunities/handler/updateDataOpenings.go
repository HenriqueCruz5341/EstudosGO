package handler

import (
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update an opening
// @Description Update some fields of an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body OpeningInputDto true "Request body"
// @Param opening_id path int true "Opening Id"
// @Success 200 {object} SuccessResponse{data=schemas.OpeningOutputDto} "desc"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{opening_id} [patch]
func UpdateDataOpeningsHandler(ctx *gin.Context) {
	request := OpeningUpdateDataDto{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
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
	sendSuccess(ctx, http.StatusOK, openingOutput)
}
