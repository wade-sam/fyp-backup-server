package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/wade-sam/fyp-backup-server/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (c *ClientMongo) Create(client *entity.Client) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection(c.collection)
	mclient := ClientToMclient(client)
	insertResult, err := collection.InsertOne(ctx, mclient)
	id := insertResult.InsertedID.(primitive.ObjectID).Hex()
	if err != nil {
		return "", errors.Wrap(err, "repository.CreateClient")
	}
	return id, nil
}

func ClientToMclient(client *entity.Client) *MGClient {
	policies, _ := policystringToHex(client.Policies)
	var mclient MGClient
	if client.ID == "" {
		id := primitive.NewObjectID()
		mclient.ID = id
	} else {
		id, _ := primitive.ObjectIDFromHex(client.ID)
		mclient.ID = id
	}
	mclient.Clientname = client.Clientname
	mclient.ConsumerID = client.ConsumerID
	mclient.Directorytree = client.Directorytree
	mclient.Ignorepath = client.Ignorepath
	mclient.Policies = policies
	mclient.Backups = client.Backups
	return &mclient
}

func MclientToClient(mclient *MGClient) *entity.Client {
	fmt.Println("MclientToClient", mclient)
	policies, _ := policyhexToString(mclient.Policies)

	client := entity.Client{
		ID:            mclient.ID.Hex(),
		Clientname:    mclient.Clientname,
		ConsumerID:    mclient.ConsumerID,
		Directorytree: mclient.Directorytree,
		Ignorepath:    mclient.Ignorepath,
		Policies:      policies,
		Backups:       mclient.Backups,
	}
	return &client
}
func policystringToHex(policies []string) ([]primitive.ObjectID, error) {
	var result []primitive.ObjectID
	if len(policies) == 0 {
		return result, nil
	}

	for i := range policies {
		phex, err := primitive.ObjectIDFromHex(policies[i])
		if err != nil {
			return nil, err
		}
		result = append(result, phex)
	}
	return result, nil
}

func policyhexToString(policies []primitive.ObjectID) ([]string, error) {
	var result []string
	if len(policies) == 0 {
		return result, nil
	}
	for i := range policies {
		hex := policies[i].Hex()
		result = append(result, hex)
	}
	return result, nil
}

func (c *ClientMongo) Update(client *entity.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	mclient := ClientToMclient(client)
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": mclient.ID},
		bson.M{"$set": mclient},
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
	id, err := primitive.ObjectIDFromHex(name)
	if err != nil {
		return entity.ErrClientCannotBeDeleted
	}
	collection := c.db.Database(c.database).Collection("clients_collection")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return entity.ErrClientCannotBeDeleted
	}
	return nil
}

func (c *ClientMongo) Get(id string) (*entity.Client, error) {
	var mclient MGClient
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	//filter := bson.M{"_id": id}
	idhex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": idhex}).Decode(&mclient)
	if err != nil {
		return nil, entity.ErrNotFound
	}

	client := MclientToClient(&mclient)
	return client, nil
}

func (c *ClientMongo) List() ([]*entity.Client, error) {
	var clients []*entity.Client
	var mclients []MGClient

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("clients_collection")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, entity.ErrCouldNotAddItem
	}
	//var test []bson.M
	if err = cursor.All(ctx, &mclients); err != nil {
		return nil, err
	}
	for _, i := range mclients {
		client := MclientToClient(&i)
		clients = append(clients, client)

	}
	fmt.Println("testing 123", clients)
	return clients, nil
}
