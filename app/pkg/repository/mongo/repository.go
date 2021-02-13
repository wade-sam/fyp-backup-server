package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/wade-sam/fyp-backup-server/pkg/client"
	"github.com/wade-sam/fyp-backup-server/pkg/policy"
	"github.com/wade-sam/fyp-backup-server/pkg/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MongoConfig contains the the contents needed to create a connection to mongo
type MongoConfig struct {
	Prefix      string
	URI         string
	Database    string
	Credentials *options.Credential
}

//Mongo holds config and instance of client
type mongoRepository struct {
	Config  MongoConfig
	client  *mongo.Client
	timeout time.Duration
}

func newMongoClient(config *MongoConfig, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s", config.Prefix, config.URI)).SetAuth(*config.Credentials))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, err
}

func NewMongoRepo(prefix, URI, database, username, password string, mongoTimeout int) (ports.PersistentRepository, error) {
	cred := options.Credential{
		Username: username,
		Password: password,
	}
	mc := MongoConfig{
		Prefix:      prefix,
		URI:         URI,
		Database:    database,
		Credentials: &cred,
	}
	repo := &mongoRepository{
		Config:  mc,
		timeout: time.Duration(mongoTimeout) * time.Second,
	}
	client, err := newMongoClient(&mc, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongRepo")
	}
	repo.client = client
	return repo, nil
}

func (r *mongoRepository) FindClient(name string) (*client.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	client := &client.Client{}
	collection := r.client.Database(r.Config.Database).Collection("clients_collection")
	filter := bson.M{"clientname": name}
	err := collection.FindOne(ctx, filter).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "repository.FindClient")
		}
		return nil, errors.Wrap(err, "Repository.Client.Find")
	}
	return client, nil
}

func (r *mongoRepository) CreateClient(client *client.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.Config.Database).Collection("clients_collection")
	insertResult, err := collection.InsertOne(ctx, client)
	if err != nil {
		return errors.Wrap(err, "repository.Client.Create")
	}
	log.Println("Inserted %v!", insertResult.InsertedID)
	return nil
}

func (r *mongoRepository) CreatePolicy(policy *policy.Policy) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.Config.Database).Collection("policy_collection")
	insertResult, err := collection.InsertOne(ctx, policy)
	if err != nil {
		return errors.Wrap(err, "repository.Policy.Create")
	}
	log.Println("Inserted %v!", insertResult.InsertedID)

	return nil
}

func (r *mongoRepository) FindPolicy(name string) (*policy.Policy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	policy := &policy.Policy{}
	collection := r.client.Database(r.Config.Database).Collection("policy_collection")
	filter := bson.M{"policyname": name}
	err := collection.FindOne(ctx, filter).Decode(&policy)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "repository.FindPolicy")
		}
		return nil, errors.Wrap(err, "Repository.Client.Find")
	}
	return policy, nil

}
func (r *mongoRepository) UpdatePolicy(name string, policy *policy.Policy) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.Config.Database).Collection("policy_collection")
	result, err := collection.UpdateMany(
		ctx,
		bson.M{"policyname": name},
		bson.M{"$set": policy},
	)
	if err != nil {
		return errors.Wrap(err, "repoository.Policy.Update")
	}
	log.Println("Updated", result)

	return nil
}

func (r *mongoRepository) DeletePolicy(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.Config.Database).Collection("policy_collection")
	_, err := collection.DeleteOne(ctx, bson.M{"policyname": name})
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) UpdateClient(name string, client *client.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.Config.Database).Collection("clients_collection")
	result, err := collection.UpdateMany(
		ctx,
		bson.M{"clientname": name},
		bson.M{"$set": client},
	)
	if err != nil {
		return errors.Wrap(err, "repoository.Policy.Update")
	}
	log.Println("Updated", result)

	return nil
}

func (r *mongoRepository) DeleteClient(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.Config.Database).Collection("clients_collection")
	_, err := collection.DeleteOne(ctx, bson.M{"clientname": name})
	if err != nil {
		return err
	}
	return nil
}
