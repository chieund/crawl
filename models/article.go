package models

import (
	"time"
)

type Article struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
