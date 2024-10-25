package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/setup"
)

func main() {
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

	err = db.CreateCollection(ctx, collectionName)

	client.Disconnect(ctx)

	router := setup.SetupRouters()
	router.Run(":8080")
}
