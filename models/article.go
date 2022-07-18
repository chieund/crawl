package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Id          int       `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tags        []Tag     `gorm:"many2many:article_tag;"`
}
