package models

import "gorm.io/gorm"

type Article struct {
	Id              int `gorm:"primary_key"`
	Title           string
	Slug            string
	Image           string
	Description     string
	Content         string
	Link            string
	Viewed          int `gorm:"default:0"`
	WebsiteId       int
	WebsiteSlug     string
	Website         Website `gorm:"foreignkey:website_id"`
	IsUpdateContent int     `gorm:"default:0"`
	gorm.Model
	Tags []Tag `gorm:"many2many:article_tag"`
}

func (article *Article) TableName() string {
	return "articles"
}

type ArticleResponse struct {
	Title           string
	Slug            string
	Image           string
	Link            string
	UpdateAt        string
	IsUpdateContent int
	Tags            []Tag
	Website         Website
}
