package models

import "gorm.io/gorm"

type Tag struct {
	Id    int `gorm: "primary_key"`
	Title string
	Slug  string
	Hot   int `gorm:"default:0"`
	gorm.Model
}

func (tag *Tag) TableName() string {
	return "tags"
}
