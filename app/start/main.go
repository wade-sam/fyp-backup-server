package start

import (
	"context"
	"fmt"
	"time"

	"github.com/wade-sam/fyp-backup-server/Infrastructure/Mongo"
	"github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/Mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func start() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	creds := options.Credential{
		Username: "root",
		Password: "fypproject",
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s", "mongodb", "database:27017")).SetAuth(creds))
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	persistentRepo := Mongo.NewClientMongo()
}
