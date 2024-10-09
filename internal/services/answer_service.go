package services

import (
	"Demo1/internal/models"
)

func AddAnswer(answer *models.Answer) error {
	if err := models.DB.Create(answer).Error; err != nil {
		return err
	}
	return nil
}
