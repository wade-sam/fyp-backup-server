package entity

import "errors"

var ErrNotFound = errors.New("Not found")

var ErrInvalidEntity = errors.New("Invalid entity")

var ErrClientAlreadyAdded = errors.New("The client has already been added to the client")
var ErrPolicyAlreadyAdded = errors.New("The policy has already been added to the client")

var ErrInvalidBackupPlan = errors.New("Invalid backup plan")

var ErrNoNewClient = errors.New("There are no new clients")

var ErrClientCannotBeDeleted = errors.New("This client could not be deleted, please check whether it is associated with any policies")