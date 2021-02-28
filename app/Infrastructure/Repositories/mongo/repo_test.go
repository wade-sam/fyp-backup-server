package mongo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	repo "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/mongo"
	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitialiseRepo() *repo.ClientMongo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	creds := options.Credential{
		Username: "root",
		Password: "fypproject",
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s", "mongodb", "database:27017")).SetAuth(creds))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	ClientsRepo := repo.NewClientMongo(client, "maindb", "clients_collection", 10)
	return ClientsRepo
}

func Test_CreateClient(t *testing.T) {
	repo := InitialiseRepo()
	clientId := entity.NewID()
	fmt.Println(clientId)
	//clientId.GetBSON()
	client1 := entity.Client{
		ConsumerID:    clientId,
		Clientname:    "Sam's MacBook Pro",
		Policies:      []entity.ID{entity.NewID(), entity.NewID()},
		Directorytree: []string{"/", "/home", "/home/sam"},
		Ignorepath:    []string{"/home/test"},
	}
	id, err := repo.Create(&client1)
	assert.Nil(t, err)
	assert.Equal(t, clientId, id)

}

func Test_ListClients(t *testing.T) {
	repo := InitialiseRepo()
	clients, err := repo.List()
	assert.Nil(t, err)
	fmt.Println(clients)
}
