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
	Title           string  `json:"title"`
	Snippet         string  `json:"snippet"`
	Slug            string  `json:"slug"`
	Image           string  `json:"image"`
	Link            string  `json:"link"`
	UpdateAt        string  `json:"update_at"`
	CreatedAt       string  `json:"created_at"`
	IsUpdateContent int     `json:"is_update_content"`
	Tags            []Tag   `json:"tags"`
	Website         Website `json:"website"`
}
