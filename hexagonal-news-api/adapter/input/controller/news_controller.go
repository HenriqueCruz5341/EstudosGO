package controller

import (
	"hexagonal-news-api/adapter/input/mapper"
	"hexagonal-news-api/adapter/input/model/request"
	"hexagonal-news-api/application/domain"
	"hexagonal-news-api/application/port/input"
	"hexagonal-news-api/configuration/logger"
	"hexagonal-news-api/configuration/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(newsUseCase input.NewsUseCase) *newsController {
	return &newsController{newsUseCase}
}

func (nc *newsController) GetNews(c *gin.Context) {
	logger.Info("Init GetNews controller api")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From:    newsRequest.From.Format("2006-01-02"),
	}

	newsResponseDomain, err := nc.newsUseCase.GetNewsService(newsDomain)
	if err != nil {
		logger.Error("Error trying to get news from news service", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, mapper.MapDomainToResponse(newsResponseDomain))
}
