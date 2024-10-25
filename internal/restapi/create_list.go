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

func GetList(c *gin.Context) {

	id := c.Query("id")

	var _ = id

	list := entities.List{
		Values: []string{},
	}

	var _ = list

	fmt.Println("Get list with id ", id)
	c.IndentedJSON(200, list)
}
