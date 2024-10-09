package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}
