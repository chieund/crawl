package models

import "gorm.io/gorm"

type Article struct {
	Id          int `gorm:"primary_key"`
	Title       string
	Slug        string
	Image       string
	Description string
	Content     string
	Link        string
	gorm.Model
	Tags []Tag `gorm:"many2many:article_tag"`
}

func (article *Article) TableName() string {
	return "articles"
}
