package handlers

import (
	"Demo1/internal/models"
	"Demo1/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddAnswer 处理添加回答请求
func AddAnswer(c *gin.Context) {
	var answer models.Answer
	questionID := c.Param("id") // 获取 URL 中的 questionID

	// 检查 questionID 是否有效
	if _, err := strconv.Atoi(questionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userID, _ := c.Get("userID")
	answer.UserID = userID.(uint)

	// 调用 AddAnswer 并传递 questionID 和 answer
	if err := services.AddAnswer(&answer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add answer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer added successfully"})
}
