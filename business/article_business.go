package business

import (
	"crawl/models"
	"crawl/pkg"
)

type ArticleStorageInterface interface {
	FindArticle(map[string]interface{}) (*models.Article, error)
	FindArticleCron(map[string]interface{}) (*models.Article, error)
	UpdateArticle(map[string]interface{}, models.Article) bool
	UpdateViewed(map[string]interface{}, int)
	CreateArticle(*models.Article)
	GetAllArticles(*pkg.Pagination) (*pkg.Pagination, error)
	GetAllArticlesCron(*pkg.Pagination) (*pkg.Pagination, error)
	GetAllArticlesByIds([]int, *pkg.Pagination) (*pkg.Pagination, error)
	FindArticleOther([]int, *pkg.Pagination) (*pkg.Pagination, error)
}

type ArticleBusiness struct {
	articleStore ArticleStorageInterface
}

func NewArticleBusiness(articleStore ArticleStorageInterface) *ArticleBusiness {
	return &ArticleBusiness{
		articleStore: articleStore,
	}
}

func (articleBusiness *ArticleBusiness) FindArticle(condition map[string]interface{}) (*models.Article, error) {
	article, err := articleBusiness.articleStore.FindArticle(condition)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (articleBusiness *ArticleBusiness) FindArticleCron(condition map[string]interface{}) (*models.Article, error) {
	article, err := articleBusiness.articleStore.FindArticleCron(condition)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (articleBusiness *ArticleBusiness) UpdateViewed(condition map[string]interface{}, viewed int) {
	articleBusiness.articleStore.UpdateViewed(condition, viewed)
}

func (articleBusiness *ArticleBusiness) UpdateArticle(condition map[string]interface{}, article models.Article) bool {
	return articleBusiness.articleStore.UpdateArticle(condition, article)
}

func (articleBusiness *ArticleBusiness) CreateArticle(article *models.Article) {
	articleBusiness.articleStore.CreateArticle(article)
}

func (articleBusiness *ArticleBusiness) GetAllArticles(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	articles, err := articleBusiness.articleStore.GetAllArticles(pagination)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (articleBusiness *ArticleBusiness) GetAllArticlesCron(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	articles, err := articleBusiness.articleStore.GetAllArticlesCron(pagination)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (articleBusiness *ArticleBusiness) GetAllArticlesByIds(ids []int, pagination *pkg.Pagination) (*pkg.Pagination, error) {
	return articleBusiness.articleStore.GetAllArticlesByIds(ids, pagination)
}

func (articleBusiness *ArticleBusiness) FindArticleOther(tagId []int, pagination *pkg.Pagination) (*pkg.Pagination, error) {
	return articleBusiness.articleStore.FindArticleOther(tagId, pagination)
}
