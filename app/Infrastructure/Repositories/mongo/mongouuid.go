package mongo

import (
	uuid "github.com/satori/go.uuid"
	//"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	//	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUUID struct {
	uuid.UUID
}

// func (mu mongoUUID) MarshalBSONValue() (bsonType.Type, []byte, error) {
// 	return bsonType.Binary, bsoncore.AppendBinary(nul, 4, mu.UUID[:]), nil
// }
