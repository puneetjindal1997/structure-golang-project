package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDb() (mongoSession *mongo.Client, ctx context.Context, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("DB_HOST")))
	if err != nil {
		return client, ctx, err
	}
	ctx, _ = context.WithTimeout(context.Background(), 60*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return client, ctx, err
	}
	return client, ctx, err
}

func InsertData(data interface{}) error {
	collection := Client.Database(os.Getenv("Database")).Collection(os.Getenv("Collection"))
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)
	return nil
}
