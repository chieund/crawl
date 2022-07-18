package business

import (
	"crawl/models"
)

type ArticleTagStorageInterface interface {
	CreateArticleTag(*models.ArticleTag)
	FindArticleIdByTagId(tagId int) []int
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

func (articleTagBusiness *ArticleTagBusiness) FindArticleIdByTagId(tagId int) []int {
	return articleTagBusiness.articleTagStore.FindArticleIdByTagId(tagId)
}
