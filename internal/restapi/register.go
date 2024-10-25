package restapi

import (
	"context"
	"time"

	"fmt"
	"log"

	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/entities"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Title Register user
// @Description User registration
// @Success 200 {object} entities.User
// @Failure 400 {object} gin.H
// @Router /register [POST]
func RegisterHandler(c *gin.Context) {
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

	users := db.Collection(collectionName)
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
	userMD := entities.MongoUser{}
	res := users.FindOne(ctx, bson.D{{Key: "login", Value: user.Login}}).Decode(&userMD)
	if res == nil {
		c.JSON(400, gin.H{"error": "User already registrated"})
		return
	}
	userMD.Login = user.Login
	userMD.HashedPassword = user.HashedPassword
	userMD.Lists = []int{}
	_, insertErr := users.InsertOne(ctx, userMD)
	if insertErr != nil {
		c.JSON(400, gin.H{"error": "User registration failed"})
		return
	}
	fmt.Println("Register user:", user)
	c.JSON(200, gin.H{"message": "User registred"})
	client.Disconnect(ctx)
}
