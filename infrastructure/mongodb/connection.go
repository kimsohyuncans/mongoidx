package mongodb

import (
	"context"
	"mongoidx/domains/models"
	"mongoidx/domains/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type connectionRepoImpl struct {
	db             *mongo.Database
	collectionName string
}

func NewConnectionRepo(db *mongo.Database) repository.ConnectionRepository {
	return &connectionRepoImpl{db: db, collectionName: "connections"}
}

func (r *connectionRepoImpl) AddConnection(ctx context.Context, connection models.Connection) (connectionID string, err error) {
	res, err := r.db.Collection(r.collectionName).InsertOne(ctx, connection)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *connectionRepoImpl) ListConnection(ctx context.Context) ([]models.Connection, error) {
	var connections []models.Connection

	cursor, err := r.db.Collection(r.collectionName).Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &connections); err != nil {
		return nil, err
	}

	return connections, nil
}

func (r *connectionRepoImpl) GetConnection(ctx context.Context, ID string) (*models.Connection, error) {
	var connection models.Connection

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = r.db.Collection(r.collectionName).FindOne(ctx, primitive.M{"_id": objID}).Decode(&connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func (r *connectionRepoImpl) DeleteConnection(ctx context.Context, ID string) error {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	_, err = r.db.Collection(r.collectionName).DeleteOne(ctx, primitive.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
