package models

import "gorm.io/gorm"

type TODO struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
