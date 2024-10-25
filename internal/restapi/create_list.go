package restapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateList(c *gin.Context) {
	var list entities.List
	if err := c.Bind(&list); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}

	dbName := "lists_for_all_life_cases"
	collectionName := "users"
	db := client.Database(dbName)

	lists := db.Collection(collectionName)

	valuesToMongo := entities.ListMongo{}

	_, err = lists.InsertOne(ctx, valuesToMongo)
	if err != nil {
		log.Fatalf("Ошибка добавления листов к MongoDB: %v", err)
	}
	fmt.Println("List created:", list)
	c.JSON(200, gin.H{"message": "List created"})
}
