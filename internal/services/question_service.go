package services

import (
	"Demo1/internal/models"
)

func CreateQuestion(question *models.Question) error {
	if err := models.DB.Create(question).Error; err != nil {
		return err
	}
	return nil
}

func UpdateQuestion(id uint, question *models.Question) error {
	if err := models.DB.Model(&models.Question{}).Where("id = ?", id).Updates(question).Error; err != nil {
		return err
	}
	return nil
}

func DeleteQuestion(id uint) error {
	if err := models.DB.Delete(&models.Question{}, id).Error; err != nil {
		return err
	}
	return nil
}
