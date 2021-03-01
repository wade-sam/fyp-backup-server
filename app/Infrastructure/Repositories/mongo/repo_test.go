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
	fmt.Println(client)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	ClientsRepo := repo.NewClientMongo(client, "maindb", "clients_collection", 10)
	return ClientsRepo
}

func Test_CreateClient(t *testing.T) {
	mg := InitialiseRepo()
	consumerID := "host1"
	fmt.Println(consumerID)
	//clientId.GetBSON()
	client1 := entity.Client{
		ConsumerID:    consumerID,
		Clientname:    "Sam's MacBook Pro",
		Policies:      []string{"p1", "p2", "p3"},
		Directorytree: []string{"/", "/home", "/home/sam"},
		Ignorepath:    []string{"/home/test"},
	}
	id, err := mg.Create(&client1)
	assert.Nil(t, err)
	fmt.Println(id)
	//assert.Equal(t, clientId, id)

}

// func Test_ListClients(t *testing.T) {
// 	repo := InitialiseRepo()
// 	clients, err := repo.List()
// 	assert.Nil(t, err)
// 	fmt.Println(clients)
// }
