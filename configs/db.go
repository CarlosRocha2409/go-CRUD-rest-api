package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongo()))

	if err != nil {
		log.Fatal("Error creating db client: ", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Error pinging db: ", err)
	}
	fmt.Println("Successfully connected to db")

	return client

}

var DB *mongo.Client = ConnectDb()

func GetCollection(client *mongo.Client, name string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(name)
	return collection
}
