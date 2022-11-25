package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func NewMongoDBClient(ctx context.Context) (*mongo.Client, error) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_URL := os.Getenv("DB_URL")

	credential := options.Credential{
		Username:   DB_USER,
		Password:   DB_PASS,
		AuthSource: DB_NAME,
	}

	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URL).SetAuth(credential))
	if err != nil {
		return nil, err
	}

	err = dbClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return dbClient, nil
}
