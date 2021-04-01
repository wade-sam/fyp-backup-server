package mongo

import (
	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MGClient struct {
	ID            primitive.ObjectID   `bson:"_id"`
	ConsumerID    string               `bson:"consumerid"`
	Clientname    string               `bson:"clientname"`
	Policies      []primitive.ObjectID `bson:"policies"`
	Directorytree []string             `bson:"treepath"`
	Ignorepath    []string             `bson:"ignore"`
	Backups       []string             `bson:"backups"`
}

type MGPolicy struct {
	PolicyID   primitive.ObjectID   `bson:"_id"`
	Policyname string               `bson:"policyname"`
	Clients    []primitive.ObjectID `bson:"clients"`
	Retention  int                  `bson:"retention"`
	State      string               `bson:"state"`
	Type       string               `bson:"type"`
	Fullbackup []string             `bson:"fullbackup"`
	IncBackup  []string             `bson:"incbackup"`
	PolicyRuns []*MGBackupRun       `bson:"policyRuns"`
}

type MGBackupRun struct {
	BackupID           string               `bson:"id"`
	Status             string               `bson:"status"`
	Type               string               `bson:"type"`
	RunTime            string               `bson:"runtime"`
	Expiry             int                  `bson:"expirydate"`
	SuccessfullClients []primitive.ObjectID `bson:"successclientruns"`
	FailedClients      []primitive.ObjectID `bson:"failedclientruns"`
}

type MGClientRun struct {
	ClientRunID   primitive.ObjectID            `bson:"_id"`
	ClientName    string                        `bson:"name"`
	Status        string                        `bson:"status"`
	TotalFiles    int                           `bson:"totalfiles"`
	BackupSuccess map[string]*entity.ClientFile `bson:"backupsuccess"`
	BackupFailure map[string]*entity.ClientFile `bson:"backupfailure"`
}
