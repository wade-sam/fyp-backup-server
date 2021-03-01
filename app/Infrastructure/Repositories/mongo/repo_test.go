package mongo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	repo "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/mongo"
	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	//fmt.Println(client)
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
	client1 := entity.Client{
		ConsumerID:    consumerID,
		Clientname:    "Sam's MacBook Pro",
		Policies:      []string{primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex()},
		Directorytree: []string{"/", "/home", "/home/sam"},
		Ignorepath:    []string{"/home/test"},
	}
	id, err := mg.Create(&client1)
	assert.Nil(t, err)
	fmt.Println(id)
	//assert.Equal(t, clientId, id)

}

func Test_ListClients(t *testing.T) {
	repo := InitialiseRepo()
	clients, err := repo.List()
	assert.Nil(t, err)
	fmt.Println(clients[0])
}

func Test_GetClients(t *testing.T) {
	repo := InitialiseRepo()
	consumerID := "host1"
	client1 := entity.Client{
		ConsumerID:    consumerID,
		Clientname:    "Sam's MacBook Pro",
		Policies:      []string{"p1", "p2", "p3"},
		Directorytree: []string{"/", "/home", "/home/sam"},
		Ignorepath:    []string{"/home/test"},
	}
	id, err := repo.Create(&client1)
	assert.Nil(t, err)
	client, err := repo.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, client1.Clientname, client.Clientname)
	fmt.Println(client)
}

func Test_UpdateClient(t *testing.T) {
	repo := InitialiseRepo()
	client, err := repo.Get("603d295d7bec9bce0969d6c7")
	assert.Nil(t, err)
	new_client_name := "Jack's Macbook Pro"
	client.Clientname = new_client_name
	client.Ignorepath = append(client.Ignorepath, "/home/fuckoff")
	err = repo.Update(client)
	assert.Nil(t, err)
	client, err = repo.Get("603d295d7bec9bce0969d6c7")
	assert.Equal(t, "Jack's Macbook Pro", client.Clientname)
}

func Test_DeleteClient(t *testing.T) {
	repo := InitialiseRepo()
	err := repo.Delete("603d295d7bec9bce0969d6c7")
	assert.Nil(t, err)
	_, err = repo.Get("603d295d7bec9bce0969d6c7")
	assert.Equal(t, entity.ErrNotFound, err)

}
