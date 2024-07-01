package input

import (
	"hexagonal-news-api/application/domain"
	"hexagonal-news-api/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
