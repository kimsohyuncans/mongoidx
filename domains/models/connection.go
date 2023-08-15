package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	ID          string        `json:"id" bson:"_id,omitempty"`
	URI         string        `json:"uri" bson:"uri"`
	ConnectedAt time.Time     `json:"connected_at" bson:"connected_at"`
	driver      *mongo.Client `json:"-" bson:"-"`
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

func (c *Connection) ListDatabases(ctx context.Context) ([]string, error) {
	return c.driver.ListDatabaseNames(ctx, bson.M{})
}

func (c *Connection) StartProfiling(ctx context.Context, db string) error {
	return c.driver.Database(db).RunCommand(ctx, bson.D{{Key: "profile", Value: 2}}).Err()
}

func (c *Connection) StopProfiling(ctx context.Context, db string) error {
	return c.driver.Database(db).RunCommand(ctx, bson.D{{Key: "profile", Value: 0}}).Err()
}

type Metrics struct {
	ID                string `json:"_id" bson:"_id"`
	TotalOps          int64  `json:"total_ops" bson:"total_ops"`
	TotalDocsExamined int64  `json:"total_docs_examined" bson:"total_docs_examined"`
	TotalDocsReturned int64  `json:"total_docs_returned" bson:"total_docs_returned"`
}

type ProfileData struct {
	Global  Metrics   `json:"global" bson:"global"`     // global metrics
	PerColl []Metrics `json:"per_coll" bson:"per_coll"` // summary per collection
}

func (c *Connection) GetProfileData(ctx context.Context, db string) (*ProfileData, error) {
	res, err := c.driver.Database(db).Collection("system.profile").Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "op", Value: bson.D{{Key: "$ne", Value: "command"}}}}}},
		bson.D{
			{Key: "$facet",
				Value: bson.D{
					{Key: "global",
						Value: bson.A{
							bson.D{
								{Key: "$group",
									Value: bson.D{
										{Key: "_id", Value: "global"},
										{Key: "total_ops", Value: bson.D{{Key: "$sum", Value: 1}}},
										{Key: "total_docs_examined", Value: bson.D{{Key: "$sum", Value: "$docsExamined"}}},
										{Key: "total_docs_returned", Value: bson.D{{Key: "$sum", Value: "$nreturned"}}},
									},
								},
							},
						},
					},
					{Key: "per_coll",
						Value: bson.A{
							bson.D{
								{Key: "$group",
									Value: bson.D{
										{Key: "_id", Value: "$ns"},
										{Key: "total_ops", Value: bson.D{{Key: "$sum", Value: 1}}},
										{Key: "total_docs_examined", Value: bson.D{{Key: "$sum", Value: "$docsExamined"}}},
										{Key: "total_docs_returned", Value: bson.D{{Key: "$sum", Value: "$nreturned"}}},
									},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$addFields",
				Value: bson.D{
					{Key: "global",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$global",
									0,
								},
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	var result []ProfileData
	if err := res.All(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func (c *Connection) Profiling(ctx context.Context, db string) (bool, error) {
	var result struct {
		Ok      int `bson:"ok"`
		Profile int `bson:"was"`
	}
	err := c.driver.Database(db).RunCommand(ctx, bson.D{{Key: "profile", Value: 2}}).Decode(&result)
	if err != nil {
		return false, err
	}

	return result.Profile != 0, nil
}

func (c *Connection) Ping(ctx context.Context) error {
	return c.driver.Ping(ctx, nil)
}
