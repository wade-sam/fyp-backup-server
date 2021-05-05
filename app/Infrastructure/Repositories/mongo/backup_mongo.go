package mongo

import (
	"context"
	"time"

	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BackupMongo struct {
	db         *mongo.Client
	timeout    time.Duration
	database   string
	collection string
}

func NewBackupMongo(db *mongo.Client, database, collection string, timeout int) *BackupMongo {
	return &BackupMongo{
		db:         db,
		timeout:    time.Duration(timeout) * time.Second,
		database:   database,
		collection: collection,
	}
}

func (b *BackupMongo) Create(clientrun *entity.ClientRun, clientID, policyID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.timeout)
	defer cancel()
	collection := b.db.Database(b.database).Collection(b.collection)
	clientrun.Client = clientID
	clientrun.Policy = policyID
	mgclientrun, err := ClientRunToMClientRun(clientrun)
	if err != nil {
		return "", err
	}
	insertResult, err := collection.InsertOne(ctx, mgclientrun)
	if err != nil {
		return "", err
	}
	id := insertResult.InsertedID.(primitive.ObjectID).Hex()
	//response, _ := b.Get(id)
	//log.Println("CLIENTRUN IN MONGO", response)
	return id, nil
}

func (b *BackupMongo) ListClientRuns(client string) ([]*entity.ClientRun, error) {
	var clientruns []*entity.ClientRun
	var mgclientrun []MGClientRun
	ctx, cancel := context.WithTimeout(context.Background(), b.timeout)
	defer cancel()
	collection := b.db.Database(b.database).Collection(b.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, entity.ErrCouldNotAddItem
	}
	if err = cursor.All(ctx, &mgclientrun); err != nil {
		return nil, err
	}
	for _, i := range mgclientrun {
		clientrun, err := MClientRunToClientRun(&i)
		if err != nil {
			return nil, err
		}
		clientruns = append(clientruns, clientrun)
	}
	return clientruns, nil
}

func (b *BackupMongo) ListClientRunsAll() ([]*entity.ClientRun, error) {
	var clientruns []*entity.ClientRun
	var mgclientrun []MGClientRun
	ctx, cancel := context.WithTimeout(context.Background(), b.timeout)
	defer cancel()
	collection := b.db.Database(b.database).Collection(b.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, entity.ErrCouldNotAddItem
	}
	if err = cursor.All(ctx, &mgclientrun); err != nil {
		return nil, err
	}
	for _, i := range mgclientrun {
		clientrun, err := MClientRunToClientRun(&i)
		if err != nil {
			return nil, err
		}
		clientruns = append(clientruns, clientrun)
	}
	return clientruns, nil
}

func (b *BackupMongo) Get(id string) (*entity.ClientRun, error) {
	var mgclientrun MGClientRun
	ctx, cancel := context.WithTimeout(context.Background(), b.timeout)
	defer cancel()
	collection := b.db.Database(b.database).Collection(b.collection)
	idhex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	err = collection.FindOne(ctx, bson.M{"_id": idhex}).Decode(&mgclientrun)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	clientrun, err := MClientRunToClientRun(&mgclientrun)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	return clientrun, nil
}

func ClientRunToMClientRun(clientrun *entity.ClientRun) (*MGClientRun, error) {
	var mgclientrun MGClientRun
	if clientrun.ID == "" {
		id := primitive.NewObjectID()
		mgclientrun.ClientRunID = id
	} else {
		id, err := primitive.ObjectIDFromHex(clientrun.ID)

		if err != nil {
			return nil, err
		}
		mgclientrun.ClientRunID = id
	}
	cid, _ := primitive.ObjectIDFromHex(clientrun.Client)
	pid, _ := primitive.ObjectIDFromHex(clientrun.Policy)
	mgclientrun.Status = clientrun.Status
	mgclientrun.Client = cid
	mgclientrun.Policy = pid
	mgclientrun.BackupSuccess = clientrun.SuccesfullFiles
	mgclientrun.BackupFailure = clientrun.FailedFiles
	mgclientrun.TotalFiles = clientrun.TotalFiles
	return &mgclientrun, nil
}

func MClientRunToClientRun(Mgclientrun *MGClientRun) (*entity.ClientRun, error) {
	//clientruns
	id := Mgclientrun.ClientRunID.Hex()
	policyID := Mgclientrun.Policy.Hex()
	clientID := Mgclientrun.Client.Hex()
	clientrun := entity.ClientRun{
		ID:              id,
		Client:          clientID,
		Policy:          policyID,
		Status:          Mgclientrun.Status,
		TotalFiles:      Mgclientrun.TotalFiles,
		SuccesfullFiles: Mgclientrun.BackupSuccess,
		FailedFiles:     Mgclientrun.BackupFailure,
	}
	return &clientrun, nil
}
