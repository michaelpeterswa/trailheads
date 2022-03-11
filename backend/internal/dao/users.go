package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersDAO struct {
	mongoClient *mongo.Client
}

type User struct {
	Username string `json:"username" bson:"username"`
	APIKey   string `json:"api-key" bson:"api-key"`
}

func NewUsersDAO(m *mongo.Client) *UsersDAO {
	return &UsersDAO{mongoClient: m}
}

func (ud UsersDAO) GetUser(ctx context.Context, username string) (*User, error) {
	usersColl := ud.mongoClient.Database("main").Collection("users")

	res := usersColl.FindOne(ctx, bson.M{"username": username})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var resUser *User
	err := res.Decode(&resUser)
	if err != nil {
		return nil, err
	}

	return resUser, nil
}
