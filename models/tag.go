package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	gorm.Model
	Id        int       `json:"id" gorm: "primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
