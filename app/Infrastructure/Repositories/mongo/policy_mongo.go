package mongo

import (
	"log"
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

func (c *PolicyMongo) GetName(name string) (string, error) {
	mcpolicy := MGPolicy{}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	collection := c.db.Database(c.database).Collection("policy_collection")
	//filter := bson.M{"_id": id}
	idhex, err := primitive.ObjectIDFromHex(name)
	if err != nil {
		return "", err
	}
	err = collection.FindOne(ctx, bson.M{"_id": idhex}).Decode(&mcpolicy)
	if err != nil {
		return "", entity.ErrNotFound
	}
	return mcpolicy.Policyname, nil
}

func (p *PolicyMongo) AddBackupRun(id string, backupRun *entity.Backups) error {

	//mbackuprun, err := BackupRunToMGBackupRun(backupRun)
	// if err != nil {
	// 	log.Println("error", err)
	// 	return entity.ErrCouldNotUpdateItem
	// }
	policy, err := p.Get(id)
	if err != nil {
		return entity.ErrCouldNotUpdateItem
	}
	policy.BackupRun = append(policy.BackupRun, backupRun)
	log.Println("OBJ:", policy.BackupRun)
	mpolicy, err := PolicyToMpolicy(policy)
	if err != nil {
		return entity.ErrCouldNotUpdateItem
	}
	//mpolicy.PolicyRuns = append(mpolicy.PolicyRuns, *mbackuprun)
	log.Println("Current Policies policy runs", mpolicy.PolicyRuns[0])
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	collection := p.db.Database(p.database).Collection(p.collection)
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": mpolicy.PolicyID},
		bson.M{"$set": mpolicy},
	)
	if err != nil {
		return entity.ErrCouldNotUpdateItem
	}
	presult, err := p.Get(id)
	//log.Println(presult)
	log.Println("Return Result", presult)
	return nil
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
		RunTime:    mpolicy.RunTime,
		Clients:    clients,
		Retention:  mpolicy.Retention,
		State:      mpolicy.State,
		Type:       mpolicy.Type,
		Fullbackup: mpolicy.Fullbackup,
		IncBackup:  mpolicy.IncBackup,
	}

	for i := range mpolicy.PolicyRuns {
		runs, err := MGBackupRunToBackupRun(&mpolicy.PolicyRuns[i])
		if err != nil {
			return nil, err
		}
		policy.BackupRun = append(policy.BackupRun, runs)
		log.Println("BACKUP RUN", runs)
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
func MGBackupRunToBackupRun(mbackup *MGBackupRun) (*entity.Backups, error) {
	//var backuprun *entity.Backups
	fail, err := clientHexToString(mbackup.FailedClients)
	if err != nil {
		return nil, err
	}
	success, err := clientHexToString(mbackup.SuccessfullClients)
	if err != nil {
		return nil, err
	}
	backuprun := entity.Backups{
		ID:                 mbackup.BackupID,
		Status:             mbackup.Status,
		Type:               mbackup.Type,
		Date:               mbackup.Date,
		Expiry:             mbackup.Expiry,
		RunTime:            mbackup.RunTime,
		SuccessFullClients: success,
		FailedClients:      fail,
	}

	return &backuprun, nil
}
func BackupRunToMGBackupRun(backup *entity.Backups) (*MGBackupRun, error) {
	var MBackupRun MGBackupRun
	failedClients, err := clientstringToHex(backup.FailedClients)
	if err != nil {
		return nil, err
	}
	successClients, err := clientstringToHex(backup.SuccessFullClients)
	if err != nil {
		return nil, err
	}
	MBackupRun.FailedClients = failedClients
	MBackupRun.SuccessfullClients = successClients

	MBackupRun.BackupID = backup.ID
	MBackupRun.Status = backup.Status
	MBackupRun.Date = backup.Date
	MBackupRun.Status = backup.Status
	MBackupRun.Type = backup.Type
	MBackupRun.RunTime = backup.RunTime
	MBackupRun.Expiry = backup.Expiry

	return &MBackupRun, nil
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
	//var mgpolicyruns []MGBackupRun
	mgpolicy.Policyname = policy.Policyname
	mgpolicy.Clients = clients
	mgpolicy.RunTime = policy.RunTime
	mgpolicy.Retention = policy.Retention
	mgpolicy.State = policy.State
	mgpolicy.Type = policy.Type
	mgpolicy.Fullbackup = policy.Fullbackup
	mgpolicy.IncBackup = policy.IncBackup
	//mgpolicy.PolicyRuns = mgpolicyruns

	for i := range policy.BackupRun {
		runs, err := BackupRunToMGBackupRun(policy.BackupRun[i])
		if err != nil {
			return nil, err
		}
		mgpolicy.PolicyRuns = append(mgpolicy.PolicyRuns, *runs)
		log.Println("BACKUP RUN", runs)
		// sclients, _ := clientstringToHex(i.SuccessFullClients)
		// fclients, _ := clientstringToHex(i.FailedClients)
		// holder := MGBackupRun{
		// 	BackupID:           i.ID,
		// 	Status:             i.Status,
		// 	Date:               i.Date,
		// 	Type:               i.Type,
		// 	RunTime:            i.RunTime,
		// 	Expiry:             i.Expiry,
		// 	SuccessfullClients: sclients,
		// 	FailedClients:      fclients,
		// }
		// mgpolicyruns = append(mgpolicyruns, holder)
	}

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
