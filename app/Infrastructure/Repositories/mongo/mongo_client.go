package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Client struct {
	ConsumerID    bson.Raw   `bson:"_id"`
	Clientname    string     `bson:"clientname"`
	Policies      []bson.Raw `bson:"policies"`
	Directorytree []string   `bson: "treepath"`
	Ignorepath    []string   `bson: "ignore"`
	Backups       []string   `bson: "backups"`
}
