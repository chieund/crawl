package models

import (
	"time"
)

type Article struct {
	Id          int
	Title       string
	Slug        string
	Image       string
	Description string
	Content     string
	Link        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
