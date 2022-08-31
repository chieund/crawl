package service

import (
	"crawl/models"
	"crawl/pkg"
)

type ArticleService struct {
	pagination *pkg.Pagination
}

func NewArticleService(pagination *pkg.Pagination) *ArticleService {
	return &ArticleService{
		pagination: pagination,
	}
}

func (article *ArticleService) FormatData() []models.ArticleResponse {
	var articleResponses []models.ArticleResponse
	for _, article := range article.pagination.Rows {
		articleResponse := models.ArticleResponse{}
		articleResponse.Title = article.Title
		articleResponse.Link = article.Link
		articleResponse.Slug = article.Slug
		articleResponse.CreatedAt = article.CreatedAt.Format("Jan 02")
		articleResponse.Tags = article.Tags
		articleResponse.Image = article.Image
		articleResponse.IsUpdateContent = article.IsUpdateContent
		articleResponse.Website = article.Website
		articleResponses = append(articleResponses, articleResponse)
	}

	return articleResponses
}
