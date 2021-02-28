package start

import (
	"context"
	"fmt"
	"time"

	mg "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/mongo"
	cl "github.com/wade-sam/fyp-backup-server/usecase/client"
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
	client, err := mg.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s", "mongodb", "database:27017")).SetAuth(creds))
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	ClientsRepo := Mongo.NewClientMongo(client, "maindb", "clients_collection", 10)
	clientService := cl.NewService(ClientsRepo)

}
