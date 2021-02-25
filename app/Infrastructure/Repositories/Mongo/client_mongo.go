package Mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientMongo struct {
	db         *mongo.Client
	timeout    time.Duration
	database   string
	collection string
}

func NewClientMongo(db *mongo.Client, database, collection string, timeout int) *ClientMongo {

	return &ClientMongo{
		db:         db,
		timeout:    time.Duration(timeout) * time.Second,
		database:   database,
		collection: collection,
	}
}

func (c *ClientMongo) Create(client *entity.Client) (*entity.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection(c.collection)
	b_client := client
	insertResult, err := collection.InsertOne(ctx, b_client)

	if err != nil {
		return nil, errors.Wrap(err, "repository.CreateClient")
	}
	return client, nil
}
