package handlers

import (
	"Demo1/internal/models"
	"Demo1/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 处理用户注册请求
func Register(c *gin.Context) {
	var user models.User
	c.JSON(http.StatusOK, gin.H{"Register": "Please POST your Username and Password!"})

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := services.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login 处理用户登录请求
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := services.AuthenticateUser(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := services.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"message": "Login successful",
		"token":   token,
	})
}

// ChangePassword 修改用户密码
func ChangePassword(c *gin.Context) {
	var userInput struct {
		NewPassword string `json:"newPassword"`
	}

	claims, err := services.ValidateToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := services.ChangePassword(models.DB, claims.UserID, userInput.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
