package mongo

import (
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
