package Mongo

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/Mongo/objects"
	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/bson"
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
	b_client := &objects.Client{
		Clientname:    client.Clientname,
		Policies:      client.Policies,
		Directorytree: client.Directorytree,
		Ignorepath:    client.Ignorepath,
		Backups:       client.Backups,
	}
	_, err := collection.InsertOne(ctx, b_client)

	if err != nil {
		return nil, errors.Wrap(err, "repository.CreateClient")
	}
	return client, nil
}

func (c *ClientMongo) Update(client *entity.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")

	b_client := &objects.Client{
		Clientname:    client.Clientname,
		Policies:      client.Policies,
		Directorytree: client.Directorytree,
		Ignorepath:    client.Ignorepath,
		Backups:       client.Backups,
	}

	result, err := collection.UpdateMany(
		ctx,
		bson.M{"clientname": b_client.Clientname},
		bson.M{"$set": b_client},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Policy.Update")
	}
	log.Println(result)
	return nil
}

func (c *ClientMongo) Delete(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	_, err := collection.DeleteOne(ctx, bson.M{"clientname": name})
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientMongo) Get(name string) (*entity.Client, error) {
	b_client := objects.Client{}
	client := &entity.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	filter := bson.M{"clientname": name}
	err := collection.FindOne(ctx, filter).Decode(&b_client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "repository.FindClient")
		}
		return nil, errors.Wrap(err, "Repository.Client.Find")
	}
	client.Clientname = b_client.Clientname
	client.Policies = b_client.Policies
	client.Directorytree = b_client.Directorytree
	client.Ignorepath = b_client.Ignorepath
	client.Backups = b_client.Backups
	return client, nil
}

func (c *ClientMongo) List() ([]*entity.Client, error) {
	clients := []*entity.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	cursor, err := collection.Find(ctx)
}
