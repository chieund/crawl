package storage

import (
	"crawl/models"
)

func (s *mysqlStorage) FindTag(condition map[string]interface{}) (*models.Tag, error) {
	var tag models.Tag

	err := s.db.Where(condition).First(&tag).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (s *mysqlStorage) UpdateTag(condition map[string]interface{}, article models.Tag) bool {
	if s.db.Where(condition).Updates(&article).RowsAffected == 0 {
		return false
	}
	return true
}

func (s *mysqlStorage) CreateTag(tag models.Tag) {
	s.db.Create(&tag)
}

func (s *mysqlStorage) GetAllTags() ([]models.Tag, error) {
	var tags []models.Tag
	s.db.Order("id DESC").Find(&tags)
	return tags, nil
}

func (s *mysqlStorage) GetAllHotTags() ([]models.Tag, error) {
	var tags []models.Tag
	s.db.Order("Title asc").Where("hot=?", 1).Find(&tags)
	return tags, nil
}
