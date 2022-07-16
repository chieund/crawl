package storage

import (
	"crawl/models"
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

func (s *mysqlStorage) CreateArticle(article models.Article) {
	s.db.Create(&article)
}

func (s *mysqlStorage) GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	s.db.Order("id DESC").Find(&articles)
	return articles, nil
}
