package service

import (
	"crawl/models"
	"crawl/pkg"
	"crawl/pkg/typesense"
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
		articleResponse.ID = article.Id
		articleResponse.Title = article.Title
		articleResponse.Link = article.Link
		articleResponse.Slug = article.Slug
		articleResponse.Image = article.Image

		var tagJsons []typesense.TagJson
		for _, tag := range article.Tags {
			var tagJson = typesense.TagJson{
				Title: tag.Title,
				Slug:  tag.Slug,
			}
			tagJsons = append(tagJsons, tagJson)
		}
		//articleResponse.Tags = tagJsons
		//articleResponse.Website = typesense.WebsiteJson{
		//	Title: article.Website.Title,
		//	Slug: article.Website.Slug,
		//}

		articleResponses = append(articleResponses, articleResponse)
	}

	return articleResponses
}
