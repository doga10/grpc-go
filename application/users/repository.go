package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	domain "grpc-go/domain/users"
	"time"
)

type UserRepository struct {
	Db *mongo.Database
}

type IUserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	Fetch(email string) (*domain.User, error)
}

func (repository *UserRepository) Create(user *domain.User) (*domain.User, error)  {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := repository.Db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Fetch(email string) (*domain.User, error) {
	var user *domain.User
	filter := bson.M{"email": email}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := repository.Db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}