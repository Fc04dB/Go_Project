package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Content    string `json:"content"`
	QuestionID uint   `json:"question_id"`
	UserID     uint   `json:"user_id"`
}
