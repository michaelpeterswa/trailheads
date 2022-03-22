package dao

import (
	"context"

	"github.com/michaelpeterswa/trailheads/backend/internal/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersDAO struct {
	mongoClient *mongo.Client
}

func NewUsersDAO(m *mongo.Client) *UsersDAO {
	return &UsersDAO{mongoClient: m}
}

func (ud UsersDAO) CreateUser(ctx context.Context, user *structs.User) error {
	usersColl := ud.mongoClient.Database("main").Collection("users")
	_, err := usersColl.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (ud UsersDAO) GetUser(ctx context.Context, username string) (*structs.User, error) {
	usersColl := ud.mongoClient.Database("main").Collection("users")

	res := usersColl.FindOne(ctx, bson.M{"username": username})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var resUser *structs.User
	err := res.Decode(&resUser)
	if err != nil {
		return nil, err
	}

	return resUser, nil
}

func (ud UsersDAO) UpdateUser(ctx context.Context, user *structs.User) error {
	usersColl := ud.mongoClient.Database("main").Collection("users")

	res := usersColl.FindOneAndReplace(ctx, bson.M{"username": user.Username}, user)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (ud UsersDAO) DeleteUser(ctx context.Context, username string) error {
	usersColl := ud.mongoClient.Database("main").Collection("users")

	res := usersColl.FindOneAndDelete(ctx, bson.M{"username": username})
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
