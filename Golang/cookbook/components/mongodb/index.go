package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var MgoDbName string = "test"

func InitConnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var err error
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(
		options.Credential{
			Username: "user",
			Password: "123456",
		},
	)
	client, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func DisConnect() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to mongoDB closed")
}
