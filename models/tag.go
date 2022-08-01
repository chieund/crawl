package models

import (
	"gorm.io/gorm"
	"math/rand"
)

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

func (tag *Tag) ShowColor() string {
	in := []string{"badge-primary", "secondary", "success", "danger", "warning", "info", "light", "dark"}
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	return pick
}

type TagResponse struct {
	Title    string
	Slug     string
	UpdateAt string
}
