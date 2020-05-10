package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMongoDB() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/app"))
	if err != nil {
		return nil
	}

	db := conn.Database("app")

	return db
}
