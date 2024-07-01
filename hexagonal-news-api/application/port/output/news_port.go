package output

import (
	"hexagonal-news-api/application/domain"
	"hexagonal-news-api/configuration/rest_err"
)

type NewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
