package storage

import (
	"crawl/models"
	"crawl/pkg"
	"gorm.io/gorm"
	"math"
)

func (s *mysqlStorage) FindArticle(condition map[string]interface{}) (*models.Article, error) {
	var article models.Article

	err := s.db.Select(
		"id", "title", "slug", "link", "content", "created_at", "image", "is_update_content", "website_id", "viewed",
	).Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "title", "slug")
	}).Preload("Website", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "title", "slug", "image")
	}).Where(condition).First(&article).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (s *mysqlStorage) FindArticleCron(condition map[string]interface{}) (*models.Article, error) {
	var article models.Article

	err := s.db.Preload("Tags").Preload("Website").Where(condition).First(&article).Error
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

func (s *mysqlStorage) UpdateViewed(condition map[string]interface{}, viewed int) {
	var articles models.Article
	s.db.Model(&articles).Where(condition).Update("viewed", viewed)
}

func (s *mysqlStorage) CreateArticle(article *models.Article) {
	s.db.Create(&article)
}

func (s *mysqlStorage) GetAllArticles(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var articles []models.Article
	var totalRows int64
	if len(pagination.Condition) == 0 {
		s.db.Model(&articles).Count(&totalRows)
	} else {
		s.db.Model(&articles).Where(pagination.Condition).Count(&totalRows)
	}
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.SetListPages()

	if len(pagination.Condition) == 0 {
		s.db.Select(
			"id", "title", "slug", "Link", "created_at", "image", "is_update_content", "website_id",
		).Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug")
		}).Preload("Website", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "image")
		}).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles)
	} else {
		s.db.Select(
			"id", "title", "slug", "Link", "created_at", "image", "is_update_content", "website_id",
		).Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "hot")
		}).Preload("Website", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "image")
		}).Where(pagination.Condition).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles)
	}

	pagination.Rows = articles
	return pagination, nil
}

func (s *mysqlStorage) GetAllArticlesCron(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var articles []models.Article
	var totalRows int64
	if len(pagination.Condition) == 0 {
		s.db.Model(&articles).Count(&totalRows)
	} else {
		s.db.Model(&articles).Where(pagination.Condition).Count(&totalRows)
	}
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.SetListPages()

	if len(pagination.Condition) == 0 {
		s.db.Preload("Tags").Preload("Website").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles)
	} else {
		s.db.Preload("Tags").Preload("Website").Where(pagination.Condition).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles)
	}

	pagination.Rows = articles
	return pagination, nil
}

func (s *mysqlStorage) GetAllArticlesByIds(ids []int, pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var articles []models.Article
	var totalRows int64
	s.db.Preload("Tags").Preload("Website").Find(&articles, ids).Count(&totalRows)
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.SetListPages()

	s.db.Preload("Tags").Preload("Website").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles, ids)
	pagination.Rows = articles
	return pagination, nil
}

func (s *mysqlStorage) FindArticleOther(tagId []int, pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var articles []models.Article
	var totalRows int64
	if len(tagId) > 0 {
		s.db.Select(
			"id", "title", "slug", "Link", "created_at", "image", "is_update_content", "website_id",
		).Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "hot")
		}).Preload("Website", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "image")
		}).Not(pagination.Condition).Joins("JOIN article_tag on article_tag.article_id=articles.id").Where("tag_id IN ?", tagId).Find(&articles).Count(&totalRows)
	} else {
		s.db.Select(
			"id", "title", "slug", "Link", "created_at", "image", "is_update_content", "website_id",
		).Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "hot")
		}).Preload("Website", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "image")
		}).Not(pagination.Condition).Find(&articles).Count(&totalRows)
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.SetListPages()

	if len(tagId) > 0 {
		s.db.Select(
			"id", "title", "slug", "Link", "created_at", "image", "is_update_content", "website_id",
		).Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "hot")
		}).Preload("Website", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "image")
		}).Not(pagination.Condition).Joins("JOIN article_tag on article_tag.article_id=articles.id").Where("tag_id IN ?", tagId).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles)
	} else {
		s.db.Select(
			"id", "title", "slug", "Link", "created_at", "image", "is_update_content", "website_id",
		).Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "hot")
		}).Preload("Website", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "slug", "image")
		}).Not(pagination.Condition).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&articles)
	}

	pagination.Rows = articles
	return pagination, nil
}
