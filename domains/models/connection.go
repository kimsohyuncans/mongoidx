package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	ID     string        `json:"id" bson:"_id,omitempty"`
	URI    string        `json:"uri" bson:"uri"`
	driver *mongo.Client `json:"-" bson:"-"`
}

func (c *Connection) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.URI))
	if err != nil {
		return err
	}
	c.driver = client
	return client.Ping(ctx, nil)
}

func (c *Connection) Disconnect(ctx context.Context) error {
	return c.driver.Disconnect(ctx)
}
