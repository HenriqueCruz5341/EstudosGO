package mapper

import (
	"hexagonal-news-api/adapter/input/model/response"
	"hexagonal-news-api/application/domain"
)

func MapDomainToResponse(newsDomain *domain.NewsDomain) *response.NewsReponse {
	newsResponse := &response.NewsReponse{
		Subject:      newsDomain.Subject,
		TotalResults: newsDomain.TotalResults,
		Articles:     mapArticleDomainToResponse(newsDomain.Articles),
	}
	return newsResponse
}

func mapArticleDomainToResponse(articlesDomain []domain.Article) []response.ArticleResponse {
	var articles []response.ArticleResponse
	for _, article := range articlesDomain {
		articleResponse := response.ArticleResponse{
			Source:      response.ArticleSourceReponse{article.Source.Id, article.Source.Name},
			Author:      article.Author,
			Title:       article.Title,
			Description: article.Description,
			URL:         article.URL,
			URLToImage:  article.URLToImage,
			PublishedAt: article.PublishedAt,
			Content:     article.Content,
		}
		articles = append(articles, articleResponse)
	}
	return articles
}
