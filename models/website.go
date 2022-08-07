package models

import "gorm.io/gorm"

type Website struct {
	Id          int `gorm:"primary_key"`
	Title       string
	Slug        string
	Image       string
	Description string
	Content     string
	Link        string
	gorm.Model
}

func (website *Website) TableName() string {
	return "websites"
}
