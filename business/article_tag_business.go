package business

import (
	"crawl/models"
)

type ArticleTagStorageInterface interface {
	CreateArticleTag(*models.ArticleTag)
	findArticleIdByTagId(tagId int) []models.ArticleTag
}

type ArticleTagBusiness struct {
	articleTagStore ArticleTagStorageInterface
}

func NewArticleTagBusiness(articleTagStore ArticleTagStorageInterface) *ArticleTagBusiness {
	return &ArticleTagBusiness{
		articleTagStore: articleTagStore,
	}
}

func (articleTagBusiness *ArticleTagBusiness) CreateArticleTag(articleTag *models.ArticleTag) {
	articleTagBusiness.articleTagStore.CreateArticleTag(articleTag)
}

func (articleTagBusiness ArticleTagBusiness) findArticleIdByTagId(tagId int) []models.ArticleTag {
	return articleTagBusiness.articleTagStore.findArticleIdByTagId(tagId)
}
