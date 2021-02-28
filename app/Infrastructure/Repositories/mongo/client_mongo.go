package mongo

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
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

func (c *ClientMongo) Create(client *entity.Client) (entity.ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection(c.collection)
	// b_client := &objects.Client{
	// 	ConsumerID:    client.ConsumerID,
	// 	Clientname:    client.Clientname,
	// 	Policies:      client.Policies,
	// 	Directorytree: client.Directorytree,
	// 	Ignorepath:    client.Ignorepath,
	// 	Backups:       client.Backups,
	// }
	_, err := collection.InsertOne(ctx, client)

	if err != nil {
		return client.ConsumerID, errors.Wrap(err, "repository.CreateClient")
	}
	return client.ConsumerID, nil
}

func (c *ClientMongo) Update(client *entity.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")

	// b_client := &objects.Client{
	// 	ConsumerID:    client.ConsumerID,
	// 	Clientname:    client.Clientname,
	// 	Policies:      client.Policies,
	// 	Directorytree: client.Directorytree,
	// 	Ignorepath:    client.Ignorepath,
	// 	Backups:       client.Backups,
	// }
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": client.ConsumerID},
		bson.M{"$set": client},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Policy.Update")
	}
	log.Println(result)
	return nil
}

func (c *ClientMongo) Delete(name entity.ID) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": name})
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientMongo) Get(id entity.ID) (*entity.Client, error) {
	var client *entity.Client
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	//filter := bson.M{"_id": id}
	response, err := collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "repository.FindClient")
		}
		return nil, errors.Wrap(err, "Repository.Client.Find")
	}
	if err = response.All(ctx, &client); err != nil {
		return nil, entity.ErrNotFound
	}

	return client, nil
}

func (c *ClientMongo) List() ([]*entity.Client, error) {
	clients := []*entity.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, entity.ErrCouldNotAddItem
	}
	if err = cursor.All(ctx, clients); err != nil {
		return nil, entity.ErrCouldNotAddItem
	}
	return clients, nil
}
