package service

import (
	"crawl/models"
	"crawl/pkg"
	"crawl/pkg/typesense"
	"strconv"
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

func (article *ArticleService) FormatDataJson() []typesense.ArticleJson {
	var articleResponses []typesense.ArticleJson
	for _, article := range article.pagination.Rows {
		articleResponse := typesense.ArticleJson{}
		articleResponse.ID = strconv.Itoa(article.Id)
		articleResponse.Title = article.Title
		articleResponse.Link = article.Link
		articleResponse.Slug = article.Slug
		articleResponse.Image = article.Image
		articleResponse.IsUpdateContent = int32(article.IsUpdateContent)
		articleResponse.CreatedAt = article.CreatedAt.Format("2006-01-02 15:04:05")
		articleResponse.UpdatedAt = article.CreatedAt.Format("2006-01-02 15:04:05")

		var tagJsons = []string{}
		for _, tag := range article.Tags {
			tagJsons = append(tagJsons, tag.Title)
		}
		articleResponse.Tags = tagJsons
		articleResponse.Website = article.Website.Slug
		articleResponses = append(articleResponses, articleResponse)
	}

	return articleResponses
}
