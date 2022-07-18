package storage

import (
	"crawl/models"
)

func (s *mysqlStorage) CreateArticleTag(articleTag *models.ArticleTag) {
	s.db.Create(&articleTag)
}

func (s *mysqlStorage) findArticleIdByTagId(tagId int) []models.ArticleTag {
	var articleTags []models.ArticleTag
	s.db.Order("id DESC").Where("tag_id=?", tagId).Find(&articleTags)
	return articleTags
}
