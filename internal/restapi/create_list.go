package restapi

import (
	"fmt"
	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/entities"
	"github.com/gin-gonic/gin"
)

func CreateListHandler(c *gin.Context) {
	var list entities.List
	if err := c.BindJSON(&list); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	fmt.Println("List created:", list)
	c.JSON(200, gin.H{"message": "List created"})
}
