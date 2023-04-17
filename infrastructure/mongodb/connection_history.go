package mongodb

import (
	"context"
	"mongoidx/domains/models"
	"mongoidx/domains/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *connectHistoryImpl) GetHistoryByID(ctx context.Context, ID string) (*models.ConnectHistory, error) {
	var connectHistory models.ConnectHistory

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = r.db.Collection(r.collectionName).FindOne(ctx, primitive.M{"_id": objID}).Decode(&connectHistory)
	if err != nil {
		return nil, err
	}

	return &connectHistory, nil
}
