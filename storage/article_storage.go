package storage

import (
	"crawl/models"
	"crawl/pkg"
	"math"
)

func (s *mysqlStorage) FindArticle(condition map[string]interface{}) (*models.Article, error) {
	var article models.Article

	err := s.db.Where(condition).First(&article).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (s *mysqlStorage) UpdateArticle(condition map[string]interface{}, article models.Article) bool {
	if s.db.Where(condition).Updates(&article).RowsAffected == 0 {
		return false
	}
	return true
}

func (s *mysqlStorage) CreateArticle(article *models.Article) {
	s.db.Create(&article)
}

func (s *mysqlStorage) GetAllArticles(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var articles []models.Article
	var totalRows int64
	s.db.Model(&articles).Count(&totalRows)
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.SetListPages()

	s.db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles)
	pagination.Rows = articles
	return pagination, nil
}

func (s *mysqlStorage) GetAllArticlesByIds(ids []int, pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var articles []models.Article
	var totalRows int64
	s.db.Find(&articles, ids).Count(&totalRows)
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.SetListPages()

	s.db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles, ids)
	pagination.Rows = articles
	return pagination, nil
}
