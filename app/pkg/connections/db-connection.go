package connections

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoConfig contains the the contents needed to create a connection to mongo
type MongoConfig struct {
	Prefix      string
	URI         string
	Database    string
	Credentials *options.Credential
}

//Mongo holds config and instance of client
type Mongo struct {
	Config           MongoConfig
	Connection       *mongo.Client
	CancelConnection *context.CancelFunc
}

//NewMongo Holds an instance containing the config and connection to mongodb
func NewMongo(config MongoConfig) *Mongo {
	return &Mongo{
		Config: config,
	}
}

//Connect used to create a new connection to mongodb
func (r *Mongo) Connect() error {
	if r.Connection == nil {

		clientconfig := options.Client().ApplyURI(fmt.Sprintf("%s://%s", r.Config.Prefix, r.Config.URI)).SetAuth(*r.Config.Credentials)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		connection, err := mongo.Connect(ctx, clientconfig)
		if err != nil {
			return err
		}
		r.Connection = connection
		r.CancelConnection = &cancel
	}
	return nil

}

//DBstart called to start connection
func DBstart() *Mongo {

	up := options.Credential{
		Username: "root",
		Password: "fypproject",
	}

	mc := MongoConfig{
		Prefix:      "mongodb",
		URI:         "database:27017",
		Database:    "maindb",
		Credentials: &up,
	}
	mcinstance := NewMongo(mc)
	if err := mcinstance.Connect(); err != nil {
		log.Fatal("Unable to connect to MongoDB", err)
	}

	return mcinstance
}
