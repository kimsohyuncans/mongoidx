package infrastructure

import (
	"context"
	"mongoidx/domains/models"
	"mongoidx/domains/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type connectHistoryImpl struct {
	db             *mongo.Database
	collectionName string
}

func NewConnectHistory(db *mongo.Database) repository.ConnectHistoryRepository {
	return &connectHistoryImpl{db: db, collectionName: "connect_history"}
}

func (r *connectHistoryImpl) AddHistory(ctx context.Context, connectionHistory models.ConnectHistory) error {
	_, err := r.db.Collection(r.collectionName).InsertOne(ctx, connectionHistory)
	if err != nil {
		return err
	}

	return nil
}

func (r *connectHistoryImpl) ListHistory(ctx context.Context) ([]models.ConnectHistory, error) {
	var histories []models.ConnectHistory

	res, err := r.db.Collection(r.collectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	res.All(ctx, &histories)

	return histories, nil
}
