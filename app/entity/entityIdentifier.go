package entity

import (
	//"github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	//"gopkg.in/mgo.v2/bson"
)

type T struct {
	UUID string `bson: "$uuid"`
}
type Foo struct {
	ID int
	T  T
}

type MyUUID struct {
	uuid.UUID
}

type Record struct {
	ID   int
	UUID MyUUID `bson:"j"`
}

func (m *MyUUID) SetBSON(raw bson.Raw) error {
	var t struct {
		UUID string `bson: "$uuid"`
	}
	err := t.UUID.Unmarshal(&t)
}

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.NewV4())
}

func StringToID(s string) (ID, error) {
	id, err := uuid.FromString(s)
	return ID(id), err
}
