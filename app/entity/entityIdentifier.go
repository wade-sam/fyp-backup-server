package entity

import (
	//"github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.NewV4())
}

func StringToID(s string) (ID, error) {
	id, err := uuid.FromString(s)
	return ID(id), err
}
