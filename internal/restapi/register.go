package restapi

import (
	"fmt"
	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/entities"
	"github.com/gin-gonic/gin"
)

// @Title Register user
// @Description User registration
// @Success 200 {object} entities.User
// @Failure 400 {object} gin.H
// @Router /register [POST]
func RegisterHandler(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	// Проверка пользователя
	if user.Login == "" {
		c.JSON(400, gin.H{"error": "User must have login"})
		return
	}
	if user.HashedPassword == "" {
		c.JSON(400, gin.H{"error": "Not find password"})
		return
	}

	// Регистрация пользователя в базе данных
	fmt.Println("Register user:", user)
	c.JSON(200, gin.H{"message": "User registred"})
}
