package dao

import (
	"context"

	"github.com/michaelpeterswa/trailheads/backend/internal/trailheads"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrailheadsDAO struct {
	mongoClient *mongo.Client
}

func NewTrailheadsDAO(m *mongo.Client) *TrailheadsDAO {
	return &TrailheadsDAO{mongoClient: m}
}

func (td TrailheadsDAO) CreateTrailhead(ctx context.Context, trailhead *trailheads.Trailhead) error {
	trailheadsColl := td.mongoClient.Database("main").Collection("trailheads")
	_, err := trailheadsColl.InsertOne(ctx, trailhead)
	if err != nil {
		return err
	}

	return nil
}

func (td TrailheadsDAO) GetTrailhead(ctx context.Context, name string) (*trailheads.Trailhead, error) {
	trailheadsColl := td.mongoClient.Database("main").Collection("trailheads")

	res := trailheadsColl.FindOne(ctx, bson.M{"name": name})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var resTrailhead *trailheads.Trailhead
	err := res.Decode(&resTrailhead)
	if err != nil {
		return nil, err
	}

	return resTrailhead, nil
}

func (td TrailheadsDAO) GetTrailheads(ctx context.Context) ([]trailheads.Trailhead, error) {
	trailheadsColl := td.mongoClient.Database("main").Collection("trailheads")

	trailheadsCursor, err := trailheadsColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer trailheadsCursor.Close(ctx)

	var trailheadList []trailheads.Trailhead
	for trailheadsCursor.Next(ctx) {
		var trailhead trailheads.Trailhead
		err := trailheadsCursor.Decode(&trailhead)
		if err != nil {
			// silent continue
			continue
		}
		trailheadList = append(trailheadList, trailhead)
	}

	return trailheadList, nil
}

func (td TrailheadsDAO) UpdateTrailhead(ctx context.Context, trailhead *trailheads.Trailhead) error {
	trailheadsColl := td.mongoClient.Database("main").Collection("trailheads")

	res := trailheadsColl.FindOneAndReplace(ctx, bson.M{"name": trailhead.Name}, trailhead)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (td TrailheadsDAO) DeleteTrailhead(ctx context.Context, name string) error {
	trailheadsColl := td.mongoClient.Database("main").Collection("trailheads")

	res := trailheadsColl.FindOneAndDelete(ctx, bson.M{"name": name})
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
