package handler

import (
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Find all openings
// @Description Find all openings
// @Tags openings
// @Success 200 {object} SuccessResponse{data=[]schemas.OpeningOutputDto} "desc"
// @Router /openings [get]
func FindAllOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error fetching openings")
		return
	}
	openingsOutput := make([]schemas.OpeningOutputDto, len(openings))
	for i, opening := range openings {
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
		openingsOutput[i] = openingOutput
	}
	sendSuccess(ctx, http.StatusOK, openingsOutput)
}
