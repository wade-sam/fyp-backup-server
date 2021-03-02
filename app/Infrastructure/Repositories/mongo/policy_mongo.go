package mongo

import (
	"time"

	"context"

	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PolicyMongo struct {
	db         *mongo.Client
	timeout    time.Duration
	database   string
	collection string
}

func NewPolicyMongo(db *mongo.Client, database, collection string, timeout int) *PolicyMongo {

	return &PolicyMongo{
		db:         db,
		timeout:    time.Duration(timeout) * time.Second,
		database:   database,
		collection: collection,
	}
}
func (p *PolicyMongo) Create(policy *entity.Policy) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	collection := p.db.Database(p.database).Collection(p.collection)
	mgpolicy, err := PolicyToMpolicy(policy)
	if err != nil {
		return "", entity.ErrCouldNotAddItem
	}
	insertResult, err := collection.InsertOne(ctx, mgpolicy)
	if err != nil {
		return "", entity.ErrCouldNotAddItem
	}
	id := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return id, nil

}

func (p *PolicyMongo) List() ([]*entity.Policy, error) {
	var policies []*entity.Policy
	var mgpolicies []MGPolicy
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	collection := p.db.Database(p.database).Collection(p.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, entity.ErrCouldNotAddItem
	}
	if err = cursor.All(ctx, &mgpolicies); err != nil {
		return nil, err
	}
	for _, i := range mgpolicies {
		policy, err := MpolicyToPolicy(&i)
		if err != nil {
			return nil, entity.ErrCouldNotAddItem
		}
		policies = append(policies, policy)
	}
	return policies, nil

}

func (p *PolicyMongo) Get(id string) (*entity.Policy, error) {
	var mgpolicy MGPolicy
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	collection := p.db.Database(p.database).Collection(p.collection)
	idhex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	err = collection.FindOne(ctx, bson.M{"_id": idhex}).Decode(&mgpolicy)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	policy, err := MpolicyToPolicy(&mgpolicy)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	return policy, nil
}

func (p *PolicyMongo) Update(policy *entity.Policy) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	collection := p.db.Database(p.database).Collection(p.collection)
	mpolicy, err := PolicyToMpolicy(policy)
	if err != nil {
		return entity.ErrCouldNotUpdateItem
	}
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": mpolicy.PolicyID},
		bson.M{"$set": mpolicy},
	)
	if err != nil {
		return entity.ErrCouldNotUpdateItem
	}
	return nil
}

func (p *PolicyMongo) Delete(sid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	collection := p.db.Database(p.database).Collection(p.collection)
	id, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return entity.ErrPolicyCantBeRemoved
	}
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return entity.ErrPolicyCantBeRemoved
	}
	if result.DeletedCount == 0 {
		return entity.ErrInvalidEntity
	}
	return nil
}

func MpolicyToPolicy(mpolicy *MGPolicy) (*entity.Policy, error) {
	//FINISHED HERE. implement the policyhextostring method
	clients, err := clientHexToString(mpolicy.Clients)
	if err != nil {
		return nil, err
	}

	policy := entity.Policy{
		PolicyID:   mpolicy.PolicyID.Hex(),
		Policyname: mpolicy.Policyname,
		Clients:    clients,
		Retention:  mpolicy.Retention,
		State:      mpolicy.State,
		Type:       mpolicy.Type,
		Fullbackup: mpolicy.Fullbackup,
		IncBackup:  mpolicy.IncBackup,
	}
	return &policy, nil

}

func clientHexToString(clients []primitive.ObjectID) ([]string, error) {
	var result []string
	if len(clients) == 0 {
		return result, nil
	}
	for i := range clients {
		chex := clients[i].Hex()
		result = append(result, chex)
	}
	return result, nil
}

func PolicyToMpolicy(policy *entity.Policy) (*MGPolicy, error) {
	clients, err := clientstringToHex(policy.Clients)
	if err != nil {
		return nil, err
	}
	var mgpolicy MGPolicy
	if policy.PolicyID == "" {
		id := primitive.NewObjectID()
		mgpolicy.PolicyID = id
	} else {
		id, _ := primitive.ObjectIDFromHex(policy.PolicyID)
		mgpolicy.PolicyID = id
	}
	mgpolicy.Policyname = policy.Policyname
	mgpolicy.Clients = clients
	mgpolicy.Retention = policy.Retention
	mgpolicy.State = policy.State
	mgpolicy.Type = policy.Type
	mgpolicy.Fullbackup = policy.Fullbackup
	mgpolicy.IncBackup = policy.IncBackup
	return &mgpolicy, nil

}
func clientstringToHex(clients []string) ([]primitive.ObjectID, error) {
	var result []primitive.ObjectID
	if len(clients) == 0 {
		return result, nil
	}

	for i := range clients {
		chex, err := primitive.ObjectIDFromHex(clients[i])
		if err != nil {
			return nil, err
		}
		result = append(result, chex)
	}
	return result, nil

}
