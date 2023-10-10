package mongo_test

import (
	"context"
	"fmt"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	repo "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/mongo"
	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitialisePolicyRepo() *repo.PolicyMongo {
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
	PolicyRepo := repo.NewPolicyMongo(client, "maindb", "policy_collection", 10)
	return PolicyRepo
}
func Test_Createpolicy(t *testing.T) {
	mg := InitialisePolicyRepo()
	policy1 := entity.Policy{
		Policyname: "wednesday backup",
		Clients:    []string{primitive.NewObjectID().Hex()},
		Retention:  10,
		Type:       "full",
		Fullbackup: []string{"Monday", "Thursday", "Sunday"},
		IncBackup:  []string{},
	}
	id, err := mg.Create(&policy1)
	assert.Nil(t, err)
	fmt.Println(id)
}

func Test_ListPolicies(t *testing.T) {
	repo := InitialisePolicyRepo()
	policies, err := repo.List()
	assert.Nil(t, err)
	fmt.Println(policies[0])
}

func Test_GetPolicies(t *testing.T) {
	repo := InitialisePolicyRepo()
	policy1 := entity.Policy{
		Policyname: "wednesday backup",
		Clients:    []string{primitive.NewObjectID().Hex()},
		Retention:  10,
		Type:       "full",
		Fullbackup: []string{"Monday", "Thursday", "Sunday"},
		IncBackup:  []string{},
	}
	id, err := repo.Create(&policy1)
	assert.Nil(t, err)
	policy, err := repo.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, policy1.Policyname, policy.Policyname)
	fmt.Println(policy)
}

func Test_UpdatePolicy(t *testing.T) {
	repo := InitialisePolicyRepo()
	id := primitive.NewObjectID().Hex()
	policy1 := entity.Policy{
		PolicyID:   id,
		Policyname: "wednesday backup",
		Clients:    []string{primitive.NewObjectID().Hex()},
		Retention:  10,
		Type:       "full",
		Fullbackup: []string{"Monday", "Thursday", "Sunday"},
		IncBackup:  []string{},
	}
	id, err := repo.Create(&policy1)
	assert.Nil(t, err)
	getpolicy1, err := repo.Get(id)
	assert.Nil(t, err)
	getpolicy1.Retention = 20
	err = repo.Update(getpolicy1)
	assert.Nil(t, err)
	get2, _ := repo.Get(id)
	assert.Equal(t, 20, get2.Retention)

}

func Test_DeletePolicy(t *testing.T) {
	repo := InitialisePolicyRepo()
	err := repo.Delete("603d789af7aef63c84fd1cc2")
	assert.Nil(t, err)
	err = repo.Delete("603d442a533b92f7920cbe40")
	assert.Equal(t, entity.ErrInvalidEntity, err)
}
