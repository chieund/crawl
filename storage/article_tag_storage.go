package storage

import (
	"crawl/models"
)

func (s *mysqlStorage) CreateArticleTag(articleTag *models.ArticleTag) {
	s.db.Create(&articleTag)
}

func (s *mysqlStorage) FindArticleIdByTagId(tagId int) []int {
	var ids []int
	s.db.Model(&models.ArticleTag{}).Select("article_id").Where("tag_id=?", tagId).Find(&ids)
	return ids
}
