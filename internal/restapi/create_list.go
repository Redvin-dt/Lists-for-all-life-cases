package restapi

import (
	"fmt"
	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/entities"
	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	var list entities.List
	if err := c.Bind(&list); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	// TODO: validate list

	// Регистрация пользователя в базе данных
	fmt.Println("List created:", list)
	c.JSON(200, gin.H{"message": "List created"})
}
