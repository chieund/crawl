package models

import "gorm.io/gorm"

type Tag struct {
	Id    int `gorm: "primary_key"`
	Title string
	Slug  string
	gorm.Model
}

func (tag *Tag) TableName() string {
	return "tags"
}
