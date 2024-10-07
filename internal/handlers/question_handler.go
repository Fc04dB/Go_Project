package handlers

import (
	"Demo1/internal/models"
	"Demo1/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateQuestion 处理提问请求
func CreateQuestion(c *gin.Context) {
	var question models.Question

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userID, _ := c.Get("userID")
	question.UserID = userID.(uint)

	if err := services.CreateQuestion(&question); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question created successfully"})
}

// UpdateQuestion 处理问题更新请求
func UpdateQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := services.UpdateQuestion(uint(id), &question); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question updated successfully"})
}

// DeleteQuestion 处理删除问题请求
func DeleteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeleteQuestion(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}
