package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Collection *mongo.Collection

func ConnectDB() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load env file")
	}

	connectionInfo := fmt.Sprintf("mongodb+srv://%s:%s@%s.%s.mongodb.net/%s?retryWrites=true&w=majority", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("CLUSTER"), os.Getenv("KEY"), os.Getenv("DB_NAME"))
	clientOption := options.Client().ApplyURI(connectionInfo)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Database connected")

	Collection = client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COL_NAME"))
}

func DisconnectDB() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Failed to disconnect database")
	}

	log.Println("Database disconnected")
}
